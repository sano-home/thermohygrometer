package model

import (
	"context"
	"testing"
	"time"
)

const testDBPath = "test.db"

func TestNewSQLite3(t *testing.T) {
	db, err := NewSQLite3(testDBPath)
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
	if db == nil {
		t.Error(`db should not be nil`)
	}
}

func TestCreate(t *testing.T) {
	db, err := NewSQLite3(testDBPath)
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
	tearDown := TestCleanUpTemperatureAndHumidityFunc(t, db)
	t.Cleanup(tearDown)

	th := &TemperatureAndHumidity{
		Temperature:   21.1,
		Humidity:      40.3,
		Unixtimestamp: time.Now().Unix(),
	}
	err = th.Create(context.Background(), db)
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
}

func TestGetLatestTemperatureAndHumidity(t *testing.T) {
	db, err := NewSQLite3(testDBPath)
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
	ths, tearDown := TestCreateTemperatureAndHumidityData(t, db)
	t.Cleanup(tearDown)
	latest := ths[len(ths)-1]

	t.Run("latest data", func(t *testing.T) {
		th, err := GetLatestTemperatureAndHumidity(context.Background(), db)
		if err != nil {
			t.Errorf(`err should be nil, got %v`, err)
		}
		if th.Temperature != latest.Temperature {
			t.Errorf(`Temperature should be %f, got %f`, latest.Temperature, th.Temperature)
		}
		if th.Humidity != latest.Humidity {
			t.Errorf(`Humidity should be %f, got %f`, latest.Humidity, th.Humidity)
		}
	})
}
func TestGetTemperatureAndHumidities(t *testing.T) {
	db, err := NewSQLite3(testDBPath)
	if err != nil {
		t.Errorf(`err should be nil, got %v`, err)
	}
	ths, tearDown := TestCreateTemperatureAndHumidityData(t, db)
	t.Cleanup(tearDown)
	size := int64(len(ths))

	tests := []struct {
		name   string
		since  time.Time
		before time.Time
		expect []*TemperatureAndHumidity
	}{
		{
			name:   "all data",
			since:  time.Date(2021, time.January, 10, 1, 2, 3, 0, time.Local),
			before: time.Date(2021, time.January, 10, 1, 2, 5, 0, time.Local),
			expect: []*TemperatureAndHumidity{
				&TemperatureAndHumidity{
					ID:            ths[size-1].ID,
					Temperature:   ths[size-1].Temperature,
					Humidity:      ths[size-1].Humidity,
					Unixtimestamp: ths[size-1].Unixtimestamp,
				},
				&TemperatureAndHumidity{
					ID:            ths[size-2].ID,
					Temperature:   ths[size-2].Temperature,
					Humidity:      ths[size-2].Humidity,
					Unixtimestamp: ths[size-2].Unixtimestamp,
				},
				&TemperatureAndHumidity{
					ID:            ths[size-3].ID,
					Temperature:   ths[size-3].Temperature,
					Humidity:      ths[size-3].Humidity,
					Unixtimestamp: ths[size-3].Unixtimestamp,
				},
			},
		},
		{
			name:   "latest data",
			since:  time.Date(2021, time.January, 10, 1, 2, 5, 0, time.Local),
			before: time.Date(2021, time.January, 10, 1, 2, 5, 0, time.Local),
			expect: []*TemperatureAndHumidity{
				&TemperatureAndHumidity{
					ID:            ths[size-1].ID,
					Temperature:   ths[size-1].Temperature,
					Humidity:      ths[size-1].Humidity,
					Unixtimestamp: ths[size-1].Unixtimestamp,
				},
			},
		},
		{
			name:   "oldest data",
			since:  time.Date(2021, time.January, 10, 1, 2, 3, 0, time.Local),
			before: time.Date(2021, time.January, 10, 1, 2, 3, 0, time.Local),
			expect: []*TemperatureAndHumidity{
				&TemperatureAndHumidity{
					ID:            ths[size-3].ID,
					Temperature:   ths[size-3].Temperature,
					Humidity:      ths[size-3].Humidity,
					Unixtimestamp: ths[size-3].Unixtimestamp,
				},
			},
		},
		{
			name:   "out of range",
			since:  time.Date(2021, time.January, 10, 1, 2, 1, 0, time.Local),
			before: time.Date(2021, time.January, 10, 1, 2, 2, 0, time.Local),
			expect: []*TemperatureAndHumidity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ths, err := GetTemperatureAndHumidities(ctx, db, tt.since, tt.before)
			if err != nil {
				t.Errorf(`err should be nil, got %v`, err)
			}
			if len(ths) != len(tt.expect) {
				t.Errorf(`len(ths) should be %d, got %d`, len(tt.expect), len(ths))
			}
			for i := len(ths); i < 0; i-- {
				if ths[i].ID != tt.expect[i].ID {
					t.Errorf(`ths[%d].ID (actual: %d) != expect[%d].ID (expect: %d)`,
						i, ths[i].ID, i, tt.expect[i].ID)
				}
				if ths[i].Temperature != tt.expect[i].Temperature {
					t.Errorf(`ths[%d].Temperature (actual: %f) != expect[%d].Temperature (expect: %f)`,
						i, ths[i].Temperature, i, tt.expect[i].Temperature)
				}
				if ths[i].Humidity != tt.expect[i].Humidity {
					t.Errorf(`ths[%d].Humidity (actual: %f) != expect[%d].Humidity (expect: %f)`,
						i, ths[i].Humidity, i, tt.expect[i].Humidity)
				}
				if float32(ths[i].Unixtimestamp) != float32(tt.expect[i].Unixtimestamp) {
					t.Errorf(`ths[%d].Unixtimestamp (actual: %d) != expect[%d].Unixtimestamp (expect: %d)`,
						i, ths[i].Unixtimestamp, i, tt.expect[i].Unixtimestamp)
				}
			}
		})
	}
}
