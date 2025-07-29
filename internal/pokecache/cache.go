package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries      map[string]cacheEntry
	cacheEntryTimeout time.Duration
	mutx              sync.Mutex
}

func NewCache(cacheTimeout time.Duration) Cache {
	return Cache{
		cacheEntries:      make(map[string]cacheEntry),
		cacheEntryTimeout: cacheTimeout,
	}
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
	entry, found := c.cacheEntries[key]

	if found {
		return entry.data, true
	}

	return nil, false
}
