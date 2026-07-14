package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	curCache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}

	go curCache.reapLoop()

	return curCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mutex.Lock()

		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	interval     time.Duration
	mutex        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
