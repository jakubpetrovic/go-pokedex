package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	dataMap   map[string]cacheEntry
	cacheLive time.Duration
	mutex     *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		dataMap: make(map[string]cacheEntry),
		mutex:   &sync.Mutex{},
	}
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.dataMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.dataMap[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.dataMap {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.dataMap, k)
		}
	}
}
