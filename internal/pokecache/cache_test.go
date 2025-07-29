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
	inputKey  string
	inputData []byte
}

func TestCacheAddGet(t *testing.T) {
	cases := []testCaseCache{
		{
			inputKey:  "some-url",
			inputData: []byte{1, 2, 3, 4, 5},
		},
		{
			inputKey:  "",
			inputData: []byte{1, 2, 3, 4, 5},
		},
		{
			inputKey:  "some-url-2",
			inputData: []byte{},
		},
	}

	for _, c := range cases {
		cache := NewCache(10 * time.Second)

		cache.Add(c.inputKey, c.inputData)
		outputData, isFound := cache.Get(c.inputKey)

		if !isFound {
			t.Errorf(RED+"Failed to retrieve cache entry whose key is %v"+RESET, c.inputKey)
			continue
		}

		if !bytes.Equal(outputData, c.inputData) {
			t.Errorf(RED+"The data of retrieved cache entry, whose key is %v, doesn't match expected data. Wanted: %v, but got %v"+RESET, c.inputKey, c.inputData, outputData)
		}
	}
}
