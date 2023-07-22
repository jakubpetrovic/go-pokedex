package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	Data      string
	Timestamp time.Time
}

type Cache struct {
	dataMap map[string]cacheEntry
	mutex   sync.Mutex
}

func NewCache(t time.Duration) {

}

func Add(key string, value []byte) {

}

func Get(key string) ([]byte, bool) {

	return []byte, false
}

func reapLoop() {

}
