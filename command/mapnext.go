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

	locationAreas, err := httpClient.GetLocationAreas(*url)

	if err != nil {
		return err
	}

	for _, loc := range locationAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
