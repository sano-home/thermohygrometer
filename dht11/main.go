// +build linux

package dht11

import (
	"context"

	dht "github.com/d2r2/go-dht"
)

type DHT11er interface {
	Get(context.Context) (float32, float32, error)
}

type DHT11 struct {
	pin        int
	retry      bool
	retryCount int
}

func NewDHT11(pin int, retry bool, retryCount int) *DHT11 {
	return &DHT11{
		pin:        pin,
		retry:      retry,
		retryCount: retryCount,
	}
}

func (d *DHT11) Get(ctx context.Context) (float32, float32, error) {
	temperature, humidity, _, err :=
		dht.ReadDHTxxWithContextAndRetry(ctx, dht.DHT11, d.pin, d.retry, d.retryCount)
	if err != nil {
		return 0.0, 0.0, err
	}
	return temperature, humidity, nil
}
