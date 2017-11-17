package lru_test

import (
	"testing"

	"github.com/ironzhang/pearls/cache/lru"
)

func TestAdd(t *testing.T) {
	var evicted int
	lru := lru.New(2, func(key, value interface{}) { evicted++ })
	lru.Add("k1", 1)
	lru.Add("k2", 2)
	lru.Add("k3", 3)

	var tests = []struct {
		name string
		key  interface{}
		hit  bool
	}{
		{"miss_k1", "k1", false},
		{"hit_k2", "k2", true},
		{"hit_k3", "k3", true},
	}
	for _, tt := range tests {
		_, ok := lru.Get(tt.key)
		if got, want := ok, tt.hit; got != want {
			t.Errorf("%s: cache hit = %v; want %v", tt.name, got, want)
		}
	}
	if got, want := evicted, 1; got != want {
		t.Errorf("evicted: %d != %d", got, want)
	}
}

func TestGet(t *testing.T) {
	var tests = []struct {
		name       string
		keyToAdd   interface{}
		keyToGet   interface{}
		expectedOk bool
	}{
		{"string_hit", "myKey", "myKey", true},
		{"string_miss", "myKey", "nonsense", false},
	}
	lru := lru.New(0, nil)
	for _, tt := range tests {
		lru.Add(tt.keyToAdd, 1234)
		val, ok := lru.Get(tt.keyToGet)
		if ok != tt.expectedOk {
			t.Fatalf("%s: cache hit = %v; want %v", tt.name, ok, !ok)
		} else if ok && val != 1234 {
			t.Fatalf("%s expected get to return 1234 but got %v", tt.name, val)
		}
	}
}

func TestReomve(t *testing.T) {
	lru := lru.New(0, nil)
	lru.Add("myKey", 1234)
	if val, ok := lru.Get("myKey"); !ok {
		t.Fatal("TestRemove returned no match")
	} else if val != 1234 {
		t.Fatalf("TestRemove falied. Expected 1234, got %v", val)
	}

	lru.Remove("myKey")
	if _, ok := lru.Get("myKey"); ok {
		t.Fatal("TestRemove returned a removed entry")
	}
}

func TestClear(t *testing.T) {
	var evicted int
	lru := lru.New(0, func(key, value interface{}) { evicted++ })
	lru.Add("k1", 1)
	lru.Add("k2", 2)
	lru.Clear()
	if got, want := evicted, 2; got != want {
		t.Errorf("evicted: %d != %d", got, want)
	}
}
