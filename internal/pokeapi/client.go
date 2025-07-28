package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	instance *http.Client
}

// factory method
func NewClient() *Client {
	clientHttp := http.Client{
		Timeout: 10 * time.Second, // 10 seconds
	}
	return &Client{
		instance: &clientHttp,
	}
}

func (c *Client) GetLocationAreas(url string) ([]LocationArea, error) {
	resp, err := c.instance.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch location areas: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PokeAPI returned status code %d when fetching location areas", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var locAreaResp LocationAreaResponse

	if err := decoder.Decode(&locAreaResp); err != nil {
		return nil, fmt.Errorf("Failed to decode json response: %w", err)
	}

	return locAreaResp.LocationAreas, nil
}
