package api

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sano-home/thermohygrometer/model"
)

func TestCurrent(t *testing.T) {
	db, err := model.NewSQLite3("../model/test.db")
	if err != nil {
		t.Error(`err should not be nil`)
	}
	_, tearDown := model.TestCreateTemperatureAndHumidityData(t, db)
	t.Cleanup(tearDown)

	s := &Server{
		db: db,
	}

	go func() {
		s.Run("127.0.0.1", "15000")
	}()

	c := http.Client{}
	resp, err := c.Get("http://127.0.0.1:15000/current")
	if err != nil {
		t.Error(`err should not be nil`)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(`err should not be nil`)
	}

	t.Logf("Response Header: %v", resp.Header)
	t.Logf("Response Body: %s", b)
	if resp.StatusCode != http.StatusOK {
		t.Errorf(`resp.StatusCode should be http.StatusOK, got %d`, resp.StatusCode)
	}
	if resp.Header["Content-Type"][0] != "application/json" {
		t.Errorf(`Content-Type should be "application/json", got %v`, resp.Header["Content-Type"][0])
	}
}
