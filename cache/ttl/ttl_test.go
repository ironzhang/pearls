package ttl_test

import (
	"testing"
	"time"

	"github.com/ironzhang/pearls/cache/ttl"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name       string
		keyToAdd   interface{}
		keyToGet   interface{}
		timeout    time.Duration
		wait       time.Duration
		expectedOK bool
	}{
		{"hit", "k1", "k1", time.Second, time.Second / 2, true},
		{"nokey_miss", "k2", "nokey", time.Second, 0, false},
		{"timeout_miss", "k3", "k3", time.Second, time.Second, false},
	}

	ttl := ttl.New(0, nil)
	for _, tt := range tests {
		ttl.Add(tt.keyToAdd, 1234, tt.timeout)
		time.Sleep(tt.wait)
		val, ok := ttl.Get(tt.keyToGet)
		if got, want := ok, tt.expectedOK; got != want {
			t.Fatalf("%s: cache hit = %v; want %v", tt.name, got, want)
		} else if ok && val != 1234 {
			t.Fatalf("%s: expected get to return 1234 but got %v", tt.name, val)
		}
	}
}
