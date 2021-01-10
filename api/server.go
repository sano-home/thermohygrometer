package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/sano-home/thermohygrometer/model"
)

type Server struct {
	db model.DBer
}

func NewServer(dbPath string) (*Server, error) {
	db, err := model.NewSQLite3(dbPath)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: db,
	}, nil
}

type CurrentTemperatureAndHumidityResponse struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Timestamp   int64   `json:"timestamp"`
}

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
		Timestamp:   th.Unixtimestamp,
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

type (
	TemperatureAndHumidityHistoriesResponse struct {
		Pages Pages  `json:"pages"`
		Data  []Data `json:"data"`
	}
	Pages struct {
		Total int `json:"total"`
	}
	Data struct {
		Temperature float32   `json:"temperature"`
		Humidity    float32   `json:"humidity"`
		Timestamp   time.Time `json:"timestamp"`
	}
)

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
	since := v.Get("since")
	before := v.Get("before")
	if since == "" || before == "" {
		log.Printf(`since = "" || before = ""`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_since, err := time.Parse(time.RFC3339, since)
	if err != nil {
		log.Printf(`time.Parse(time.RFC3339, since) failed: %v`, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_before, err := time.Parse(time.RFC3339, before)
	if err != nil {
		log.Printf(`time.Parse(time.RFC3339, before) failed: %v`, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	var data []Data
	for _, v := range ths {
		data = append(data, Data{
			Temperature: v.Temperature,
			Humidity:    v.Humidity,
			Timestamp:   time.Unix(v.Unixtimestamp, 0).UTC(),
		})
	}
	resp := TemperatureAndHumidityHistoriesResponse{
		Pages: Pages{
			Total: len(ths),
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

func (s *Server) Run(host, port string) error {
	http.HandleFunc("/current", s.CurrentTemperatureAndHumidity)
	http.HandleFunc("/histories", s.TemperatureAndHumidityHistories)
	addr := net.JoinHostPort(host, port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}
