package pokecache

import (
	"time"
)

type Cache struct {
	CacheEntries      map[string]cacheEntry
	CacheEntryTimeout time.Duration
	// lock sync.Mutex
}

func NewCache(cacheTimeout time.Duration) Cache {
	return Cache{
		CacheEntries:      make(map[string]cacheEntry),
		CacheEntryTimeout: cacheTimeout,
	}
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte // the data we are caching
}

func (c *Cache) Add(key string, data []byte) {

}

func (c *Cache) Get(key string) (data []byte, found bool) {
	return []byte{}, false
}
