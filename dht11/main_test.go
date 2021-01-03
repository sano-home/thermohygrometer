// +build linux

package dht11

import (
	"testing"
)

func TestNewDHT11(t *testing.T) {
	d := NewDHT11(14, true, 10)
	if d == nil {
		t.Error(`err should not be nil`)
	}
}
