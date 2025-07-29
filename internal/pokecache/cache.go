package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	interval     time.Duration
	mutx         sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}
	go cache.reapLoop()
	return &cache
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte // the data we are caching
}

func (c *Cache) Add(key string, data []byte) {
	c.mutx.Lock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
	c.mutx.Unlock()
}

func (c *Cache) Get(key string) (data []byte, found bool) {
	c.mutx.Lock()
	entry, found := c.cacheEntries[key]
	c.mutx.Unlock()

	if found {
		return entry.data, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		c.mutx.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mutx.Unlock()
	}
}
