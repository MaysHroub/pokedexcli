package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
)

func Explore(cfg *configuration.Config, locationName string) error {
	client := cfg.HttpClient

	pokemons, err := client.GetPokemons(locationName)

	if err != nil {
		return err
	}

	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Name)
	}

	return nil
}
