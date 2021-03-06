// +build linux

package collector

import (
	"context"
	"database/sql"
	"time"

	"github.com/sano-home/thermohygrometer/dht11"
	"github.com/sano-home/thermohygrometer/model"
)

// Collector is a Collector.
type Collector struct {
	dht dht11.DHT11er
	db  model.DBer
}

// NewCollector returns Collector.
func NewCollector(pin int, dbPath string) (*Collector, error) {
	dht := dht11.NewDHT11(pin, true, 10)

	db, err := model.NewSQLite3(dbPath)
	if err != nil {
		return nil, err
	}
	return &Collector{
		dht: dht,
		db:  db,
	}, nil
}

// Run runs a collector.
func (c *Collector) Run(ctx context.Context) error {
	// get data from dht11
	temperature, humidity, err := c.dht.Get(ctx)
	if err != nil {
		return err
	}

	// insert data into dht11
	opts := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}
	tx, err := c.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	th := &model.TemperatureAndHumidity{
		Temperature:   temperature,
		Humidity:      humidity,
		Unixtimestamp: time.Now().UTC().Unix(),
	}
	err = th.Create(ctx, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
