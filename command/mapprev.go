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

	locationAreas, err := httpClient.GetLocationAreas(*url)

	if err != nil {
		return err
	}

	for _, loc := range locationAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
