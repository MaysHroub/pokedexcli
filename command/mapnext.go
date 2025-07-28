package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
)

func MapNext(c *configuration.UrlConfig) error {
	url := c.Next

	if url == nil {
		return fmt.Errorf("There are no more location areas")
	}

	httpClient := pokeapi.NewClient()

	locationAreaResp, err := httpClient.GetLocationAreaResponse(*url)

	if err != nil {
		return err
	}

	c.Next = locationAreaResp.Next

	for _, loc := range locationAreaResp.LocationAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
