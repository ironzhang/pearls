package ttl

import (
	"time"

	"github.com/ironzhang/pearls/cache/lru"
)

type entry struct {
	value    interface{}
	deadline time.Time
}

type Cache struct {
	lru     lru.Cache
	timeout time.Duration
}

func New(timeout time.Duration, maxEntries int, onEvicted func(key, value interface{})) *Cache {
	return new(Cache).Init(timeout, maxEntries, onEvicted)
}

func (c *Cache) Init(timeout time.Duration, maxEntries int, onEvicted func(key, value interface{})) *Cache {
	c.timeout = timeout
	c.lru.Init(maxEntries, onEvicted)
	return c
}

func (c *Cache) Len() int {
	return c.lru.Len()
}

func (c *Cache) Add(key, value interface{}) {
	c.AddWithTimeout(key, value, c.timeout)
}

func (c *Cache) AddWithTimeout(key, value interface{}, timeout time.Duration) {
	var deadline time.Time
	if timeout > 0 {
		deadline = time.Now().Add(timeout)
	}
	c.lru.Add(key, entry{value: value, deadline: deadline})
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	e, ok := c.getEntry(key)
	if !ok {
		return nil, false
	}
	if !e.deadline.IsZero() && time.Now().After(e.deadline) {
		c.lru.Remove(key)
		return nil, false
	}
	return e.value, true
}

func (c *Cache) Remove(key interface{}) {
	c.lru.Remove(key)
}

func (c *Cache) RemoveOldest() {
	c.lru.RemoveOldest()
}

func (c *Cache) getEntry(key interface{}) (entry, bool) {
	if v, ok := c.lru.Get(key); ok {
		return v.(entry), true
	}
	return entry{}, false
}

func (c *Cache) Clear() {
	c.lru.Clear()
}
