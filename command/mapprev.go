package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
)

func MapPrev(c *configuration.UrlConfig) error {
	url := c.Prev

	if url == nil {
		return fmt.Errorf("No more previous locations")
	}

	httpClient := pokeapi.NewClient()

	locationAreaResp, err := httpClient.GetLocationAreaResponse(*url)

	if err != nil {
		return err
	}

	c.Prev = locationAreaResp.Previous

	for _, loc := range locationAreaResp.LocationAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
