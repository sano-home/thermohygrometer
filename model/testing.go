package model

import (
	"context"
	"testing"
	"time"
)

func TestCreateTemperatureAndHumidityData(t *testing.T, db DBer) ([]*TemperatureAndHumidity, func()) {
	var (
		ths []*TemperatureAndHumidity
		th  *TemperatureAndHumidity
		err error
	)

	// test data 1
	th = &TemperatureAndHumidity{
		Temperature:   11.1,
		Humidity:      22.2,
		Unixtimestamp: time.Now().Unix(),
	}
	err = th.Create(context.Background(), db)
	if err != nil {
		t.Fatalf(`err: %v`, err)
	}
	ths = append(ths, th)

	// test data 2
	th = &TemperatureAndHumidity{
		Temperature:   33.3,
		Humidity:      44.4,
		Unixtimestamp: time.Now().Unix(),
	}
	err = th.Create(context.Background(), db)
	if err != nil {
		t.Fatalf(`err: %v`, err)
	}
	ths = append(ths, th)

	// test data 3
	th = &TemperatureAndHumidity{
		Temperature:   55.5,
		Humidity:      66.6,
		Unixtimestamp: time.Now().Unix(),
	}
	err = th.Create(context.Background(), db)
	if err != nil {
		t.Fatalf(`err: %v`, err)
	}
	ths = append(ths, th)

	cleanup := TestCleanUpTemperatureAndHumidityFunc(t, db)
	return ths, cleanup
}

func TestCleanUpTemperatureAndHumidityFunc(t *testing.T, db DBer) func() {
	return func() {
		_, err := db.ExecContext(context.Background(),
			`DELETE FROM temperature_and_humidity`)
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.ExecContext(context.Background(),
			`VACUUM`)
		if err != nil {
			t.Fatal(err)
		}
	}
}
