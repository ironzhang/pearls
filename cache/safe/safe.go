package safe

import "sync"

type Interface interface {
	Len() int
	Add(key, value interface{})
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{})
	Clear()
}

type Cache struct {
	m sync.Mutex
	i Interface
}

func New(i Interface) *Cache {
	return new(Cache).Init(i)
}

func (c *Cache) Init(i Interface) *Cache {
	c.i = i
	return c
}

func (c *Cache) Len() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.i.Len()
}

func (c *Cache) Add(key, value interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	c.i.Add(key, value)
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.i.Get(key)
}

func (c *Cache) Remove(key interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	c.i.Remove(key)
}

func (c *Cache) Clear() {
	c.m.Lock()
	defer c.m.Unlock()
	c.i.Clear()
}
