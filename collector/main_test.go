// +build linux

package collector

import (
	"context"
	"testing"

	"github.com/sano-home/thermohygrometer/model"
)

func TestNewCollector(t *testing.T) {
	c, err := NewCollector(14, "../model/test.db")
	if err != nil {
		t.Errorf(`err should be nil, got: %v`, err)
	}
	if c == nil {
		t.Error(`c should not be nil`)
	}
}

type mockedDHT11 struct{}

func (m *mockedDHT11) Get(ctx context.Context) (float32, float32, error) {
	return 1.1, 2.2, nil
}

func TestRun(t *testing.T) {
	db, err := model.NewSQLite3("../model/test.db")
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
	tearDown := model.TestCleanUpTemperatureAndHumidityFunc(t, db)
	t.Cleanup(tearDown)

	c := &Collector{
		dht: &mockedDHT11{},
		db:  db,
	}
	err = c.Run(context.Background())
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}

	var th model.TemperatureAndHumidity
	err = db.QueryRowContext(context.Background(),
		`select id, temperature, humidity, unixtimestamp from temperature_and_humidity limit 1`,
	).Scan(&th.ID, &th.Temperature, &th.Humidity, &th.Unixtimestamp)
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
	if th.Temperature != 1.1 {
		t.Errorf(`Temperature should be 1.1, got %f`, th.Temperature)
	}
	if th.Humidity != 2.2 {
		t.Errorf(`Humidity should be 2.2, got %f`, th.Humidity)
	}
}
