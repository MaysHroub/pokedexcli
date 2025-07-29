package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	instance http.Client
}

// factory method
func NewClient(timeout time.Duration) Client {
	clientHttp := http.Client{
		Timeout: timeout,
	}
	return Client{
		instance: clientHttp,
	}
}

func (c *Client) GetLocationAreaResponse(pageUrl *string) (LocationAreaResponse, error) {
	url := BaseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	resp, err := c.instance.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Failed to fetch location area response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("PokeAPI returned status code %d when fetching location area response", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var locAreaResp LocationAreaResponse

	if err := decoder.Decode(&locAreaResp); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Failed to decode json response: %w", err)
	}

	return locAreaResp, nil
}
