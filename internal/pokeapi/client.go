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

func (c *Client) GetLocationAreaResponse(pageUrl *string) (LocationAreaResponse, error) {
	url := BaseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// check if the url is cached
	data, found := c.cache.Get(url)

	if !found {
		resp, err := c.instance.Get(url)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("Failed to fetch location area response: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return LocationAreaResponse{}, fmt.Errorf("PokeAPI returned status code %d when fetching location area response", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("Failed to read response body: %w", err)
		}
	}

	var locAreaResp LocationAreaResponse
	if err := json.Unmarshal(data, &locAreaResp); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Failed to parse json: %w", err)
	}

	return locAreaResp, nil
}
