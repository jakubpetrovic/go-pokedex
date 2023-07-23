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
	dataMap map[string]cacheEntry
	mutex   sync.Mutex
}

func NewCache(t time.Duration) *Cache {
	c := Cache{}
	reapLoop()
	return &c
}

func Add(key string, value []byte, c *Cache) {
	c.dataMap[key].val = value
	c.dataMap[createdAt] = time.Time
}

func Get(key string, c *Cache) ([]byte, bool) {
	res := c.dataMap[key].val
	return res, false
}

func reapLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
}
