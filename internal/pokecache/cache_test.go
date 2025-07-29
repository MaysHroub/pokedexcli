package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

const (
	RED   = "\033[31m"
	GREEN = "\033[32m"
	RESET = "\033[0m"
)

func TestAddGet(t *testing.T) {
	const interval = 10 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://test.com",
			val: []byte("metadata"),
		},
		{
			key: "https://test.com/some-path",
			val: []byte("more metadata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)

			cache.Add(c.key, c.val)
			outputVal, isFound := cache.Get(c.key)

			if !isFound {
				t.Errorf(RED+"Failed to retrieve cache entry whose key is %v"+RESET, c.key)
				return
			}

			if !bytes.Equal(outputVal, c.val) {
				t.Errorf(RED+"The data of retrieved cache entry, whose key is %v, doesn't match expected data. Wanted: %v, but got %v"+RESET, c.key, c.val, outputVal)
				return
			}
		})
	}

}
