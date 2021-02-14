package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/sano-home/thermohygrometer/model"
)

// Server is an API Server.
type Server struct {
	db model.DBer
}

// NewServer Returns Server.
func NewServer(dbPath string) (*Server, error) {
	db, err := model.NewSQLite3(dbPath)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: db,
	}, nil
}

// CurrentTemperatureAndHumidityResponse represents the response from API Server.
type CurrentTemperatureAndHumidityResponse struct {
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
}

// CurrentTemperatureAndHumidity handles a HTTP request for the latest temperature and humidity.
func (s *Server) CurrentTemperatureAndHumidity(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		log.Printf("c.db.BeginTx failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	th, err := model.GetLatestTemperatureAndHumidity(ctx, tx)
	if err != nil {
		log.Printf("model.GetLatestTemperatureAndHumidity failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := &CurrentTemperatureAndHumidityResponse{
		Temperature: th.Temperature,
		Humidity:    th.Humidity,
		Timestamp:   time.Unix(th.Unixtimestamp, 0).UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("json.NewEncoder(w).Encode(resp) failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

// TemperatureAndHumidityHistoriesResponse represents the response from API server.
type TemperatureAndHumidityHistoriesResponse struct {
	Pages Pages  `json:"pages"`
	Data  []Data `json:"data"`
}

// Pages represents page information.
type Pages struct {
	Total int `json:"total"`
}

// Data represents Temperature, Humidity and Timestamp.
type Data struct {
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
}

func retrieve(ths []*model.TemperatureAndHumidity, before time.Time, intervalMilliSecond, count int) []Data {
	var data []Data

	var (
		c int
		i int
	)
	for {
		if c >= count || i > len(ths) {
			break
		}
		d := time.Duration(-1*intervalMilliSecond*c) * time.Millisecond
		for i < len(ths) {
			tmp := ths[i].Temperature
			hum := ths[i].Humidity
			ts := ths[i].Unixtimestamp
			if time.Unix(ts, 0).UTC().Before(before.Add(d)) {
				data = append(data, Data{
					Temperature: tmp,
					Humidity:    hum,
					Timestamp:   time.Unix(ts, 0).UTC(),
				})
				c++
				i++
				break
			}
			i++
		}
		i++
	}
	return data
}

// TemperatureAndHumidityHistories handles a HTTP request for temperature and humidity history.
func (s *Server) TemperatureAndHumidityHistories(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		log.Printf("c.db.BeginTx failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	v := r.URL.Query()
	if v == nil {
		log.Printf("r.URL.Query() is nil")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// before: optional => default now
	var _before time.Time
	before := v.Get("before")
	if before == "" {
		_before, err = time.Parse(time.RFC3339, before)
		if err != nil {
			log.Printf(`time.Parse(time.RFC3339, before) failed: %v`, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_before = _before.UTC()
	} else {
		_before = time.Now().UTC()
	}

	// interval: require milli sec
	var _interval int
	interval := v.Get("interval")
	if interval == "" {
		log.Printf(`interval = ""`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_interval, err = strconv.Atoi(interval)
	if err != nil {
		log.Printf(`strconv.Atoi(interval) failed: %v`, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// count: require
	var _count int
	count := v.Get("count")
	if count == "" {
		log.Printf(`count = ""`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_count, err = strconv.Atoi(count)
	if err != nil {
		log.Printf(`strconv.Atoi(count) failed: %v`, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	diff := time.Duration(-1*_interval*_count) * time.Millisecond
	_since := _before.Add(diff)
	if !_before.After(_since) {
		log.Printf(`"since" should be after "before"`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ths, err := model.GetTemperatureAndHumidities(ctx, tx, _since, _before)
	if err != nil {
		log.Printf("model.GetLatestTemperatureAndHumidity failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := retrieve(ths, _before, _interval, _count)
	resp := TemperatureAndHumidityHistoriesResponse{
		Pages: Pages{
			Total: len(data),
		},
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("json.NewEncoder(w).Encode(resp) failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

// Run runs a http server.
func (s *Server) Run(host, port string) error {
	http.HandleFunc("/current", s.CurrentTemperatureAndHumidity)
	http.HandleFunc("/histories", s.TemperatureAndHumidityHistories)
	addr := net.JoinHostPort(host, port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}
