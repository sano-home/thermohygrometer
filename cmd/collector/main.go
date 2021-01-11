// +build linux

package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sano-home/thermohygrometer/collector"
)

var (
	pin      = flag.Int("gpio", 0, "The GPIO pin number where data from")
	dbPath   = flag.String("db", "", "The path to the SQLite3 db file")
	interval = flag.Duration("interval", 10*time.Second, "The time-interval (ex: -interval 10s, -interval 1m)")
)

func run(c *collector.Collector) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := c.Run(ctx); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	c, err := collector.NewCollector(*pin, *dbPath)
	if err != nil {
		log.Printf("api.NewServer failed: %v", err)
		os.Exit(1)
	}
	for {
		err := run(c)
		if err != nil {
			log.Printf("run failed: %v", err)
		}
		time.Sleep(*interval)
	}
}
