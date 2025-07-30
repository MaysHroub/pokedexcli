package pokeapi

import (
	"encoding/json"
	"fmt"
	"github/MaysHroub/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

type Client struct {
	instance http.Client
	cache    pokecache.Cache
}

// factory method
func NewClient(timeout, cacheInterval time.Duration) Client {
	clientHttp := http.Client{
		Timeout: timeout,
	}
	return Client{
		instance: clientHttp,
		cache:    *pokecache.NewCache(cacheInterval),
	}
}

func FetchData[T any](c *Client, url *string) (result T, err error) {

	if url == nil {
		return result, fmt.Errorf("failed to make request; empty url")
	}

	data, exists := c.cache.Get(*url)

	if !exists {
		resp, err := c.instance.Get(*url)
		if err != nil {
			return result, fmt.Errorf("failed to GET %v: %w", *url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return result, fmt.Errorf("bad status %v returned", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return result, fmt.Errorf("failed to read response body: %w", err)
		}
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to decode response body: %w", err)
	}

	c.cache.Add(*url, data)

	return result, nil
}
