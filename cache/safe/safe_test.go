package safe_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ironzhang/pearls/cache/lru"
	"github.com/ironzhang/pearls/cache/safe"
	"github.com/ironzhang/pearls/cache/ttl"
)

func TestLRU(t *testing.T) {
	lru := lru.New(0, nil)
	safe := safe.New(lru)

	tests := []struct {
		key  interface{}
		miss bool
	}{
		{key: "k1", miss: false},
		{key: "k2", miss: true},
		{key: "k3", miss: false},
		{key: "k4", miss: true},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		if tt.miss {
			continue
		}
		wg.Add(1)
		go func(key interface{}) {
			defer wg.Done()
			safe.Add(key, 1)
		}(tt.key)
	}
	wg.Wait()

	for _, tt := range tests {
		_, hit := safe.Get(tt.key)
		if got, want := hit, !tt.miss; got != want {
			t.Errorf("%q: cache hit = %v; want %v", tt.key, got, want)
		}
	}
}

func TestTTL(t *testing.T) {
	wait := 100 * time.Millisecond
	lru := ttl.New(wait, 0, nil)
	safe := safe.New(lru)

	tests := []struct {
		key  interface{}
		miss bool
	}{
		{key: "k1", miss: false},
		{key: "k2", miss: true},
		{key: "k3", miss: false},
		{key: "k4", miss: true},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		if tt.miss {
			continue
		}
		wg.Add(1)
		go func(key interface{}) {
			defer wg.Done()
			safe.Add(key, 1)
		}(tt.key)
	}
	wg.Wait()

	for _, tt := range tests {
		_, hit := safe.Get(tt.key)
		if got, want := hit, !tt.miss; got != want {
			t.Errorf("%q: cache hit = %v; want %v", tt.key, got, want)
		}
	}
	time.Sleep(wait)
	for _, tt := range tests {
		_, hit := safe.Get(tt.key)
		if got, want := hit, false; got != want {
			t.Errorf("%q: wait; cache hit = %v; want %v", tt.key, got, want)
		}
	}
}

func BenchmarkSafeLRU(b *testing.B) {
	lru := lru.New(0, nil)
	safe := safe.New(lru)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			safe.Add(i, i)
			safe.Get(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkSafeTTL(b *testing.B) {
	ttl := ttl.New(0, 0, nil)
	safe := safe.New(ttl)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			safe.Add(i, i)
			safe.Get(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
