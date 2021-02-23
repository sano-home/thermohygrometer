package api

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sano-home/thermohygrometer/model"
)

func waitFor(host, port string) {
	addr := net.JoinHostPort(host, port)
	for {
		_, err := net.Dial("tcp", addr)
		if err == nil {
			break
		}
	}
}

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
	waitFor("127.0.0.1", "15000")

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

func TestRetrieve(t *testing.T) {
	var ths []*model.TemperatureAndHumidity
	for i := 1000; i > 0; i-- {
		// 2021-01-01 01:24:20 - 2021-01-01 00:01:05
		ths = append(ths, &model.TemperatureAndHumidity{
			ID:            int64(i),
			Temperature:   float32(i),
			Humidity:      float32(i),
			Unixtimestamp: time.Date(2021, time.January, 1, 0, 1, i*5, 0, time.UTC).Unix(),
		})
	}
	tests := []struct {
		name                string
		before              time.Time
		intervalMilliSecond int
		count               int
		ths                 []*model.TemperatureAndHumidity

		expectCount int
	}{
		{
			name:                "before: 2021-01-01 01:25:00",
			before:              time.Date(2021, time.January, 1, 1, 25, 0, 0, time.UTC),
			intervalMilliSecond: 60 * 1000, // 1 min
			count:               5,
			ths:                 ths,
			expectCount:         5,
		},
		{
			name:                "before: 2021-01-01 01:24:00",
			before:              time.Date(2021, time.January, 1, 1, 24, 0, 0, time.UTC),
			intervalMilliSecond: 60 * 1000, // 1 min
			count:               5,
			ths:                 ths,
			expectCount:         5,
		},
		{
			name:                "before: 2021-01-01 00:01:10",
			before:              time.Date(2021, time.January, 1, 0, 1, 10, 0, time.UTC),
			intervalMilliSecond: 60 * 1000, // 1 min
			count:               5,
			ths:                 ths,
			expectCount:         1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := retrieve(tt.ths, tt.before, tt.intervalMilliSecond, tt.count)
			for _, v := range data {
				t.Logf("%+v\n", v)
			}
			if tt.expectCount != len(data) {
				t.Errorf(`tt.expectCount should be equal to len(data), %d != %d`, tt.expectCount, len(data))
			}
		})
	}
}
