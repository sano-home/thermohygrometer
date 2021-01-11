package model

import (
	"context"
	"testing"
	"time"
)

// TestCreateTemperatureAndHumidityData is a helper function for testing.
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
		Unixtimestamp: time.Date(2021, time.January, 10, 1, 2, 3, 0, time.UTC).Unix(),
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
		Unixtimestamp: time.Date(2021, time.January, 10, 1, 2, 4, 0, time.UTC).Unix(),
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
		Unixtimestamp: time.Date(2021, time.January, 10, 1, 2, 5, 0, time.UTC).Unix(),
	}
	err = th.Create(context.Background(), db)
	if err != nil {
		t.Fatalf(`err: %v`, err)
	}
	ths = append(ths, th)

	cleanup := TestCleanUpTemperatureAndHumidityFunc(t, db)
	return ths, cleanup
}

// TestCleanUpTemperatureAndHumidityFunc is a helper function for testing.
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
