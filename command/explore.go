package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
)

func Explore(cfg *configuration.Config, locationName string) error {
	client := cfg.HttpClient

	pokemons := client.GetPokemons(locationName)

	if pokemons == nil {
		return fmt.Errorf("No pokemons returned. Invalid location name: %v", locationName)
	}

	for _, pokemon := range pokemons {
		fmt.Println(pokemon.name)
	}

	return nil
}
