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
	lru lru.Cache
}

func New(maxEntries int, onEvicted func(key, value interface{})) *Cache {
	return new(Cache).Init(maxEntries, onEvicted)
}

func (c *Cache) Init(maxEntries int, onEvicted func(key, value interface{})) *Cache {
	c.lru.Init(maxEntries, onEvicted)
	return c
}

func (c *Cache) Len() int {
	return c.lru.Len()
}

func (c *Cache) Add(key, value interface{}, timeout time.Duration) {
	e := entry{value: value, deadline: time.Now().Add(timeout)}
	c.lru.Add(key, e)
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	e, ok := c.getEntry(key)
	if !ok {
		return nil, false
	}
	if time.Now().After(e.deadline) {
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
