package pokecache

import (
	"bytes"
	"testing"
	"time"
)

const (
	RED   = "\033[31m"
	GREEN = "\033[32m"
	RESET = "\033[0m"
)

type testCaseCache struct {
	inputKey           string
	inputData          []byte
	expectedAddedEntry cacheEntry
}

func TestAddCacheEntry(t *testing.T) {
	cases := []testCaseCache{
		testCaseCache{
			inputKey:  "some-url",
			inputData: []byte{1, 2, 3, 4, 5},
			expectedAddedEntry: cacheEntry{
				createdAt: time.Now(),
				data:      []byte{1, 2, 3, 4, 5},
			},
		},
	}

	cache := NewCache(10 * time.Second)

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputData)

		if len(cache.CacheEntries) == 0 {
			t.Error(RED + "Failed to add a cache entry. The cache is empty!" + RESET)
			continue
		}

		addedCacheEntry := cache.CacheEntries[c.inputKey]

		if !bytes.Equal(c.expectedAddedEntry.data, addedCacheEntry.data) {
			t.Errorf(RED+"Failed to add a cache entry. Expected data: %v, but got: %v"+RESET, c.expectedAddedEntry, addedCacheEntry.data)
			continue
		}
	}
}
