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
		log.Printf("model.GetLatestTemperatureAndHumidity failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

type TemperatureAndHumidityHistoriesResponse struct {
	Pages struct {
		Total   int `json:"total"`
		Current int `json:"current"`
	} `json:"pages"`
	Data []struct {
		Temperature float32 `json:"temperature"`
		Humidity    float32 `json:"humidity"`
		Timestamp   int64   `json:"timestamp"`
	} `json:"data"`
}

// TODO
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

	limit := int64(100)
	offset := int64(0)
	_, err = model.GetTemperatureAndHumidities(ctx, tx, limit, offset)
	if err != nil {
		log.Printf("model.GetLatestTemperatureAndHumidity failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
