package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
)

func Explore(cfg *configuration.Config, locationName string) error {
	client := cfg.HttpClient

	pokemonEncounters, err := client.GetPokemons(locationName)

	if err != nil {
		return err
	}

	for _, pokemonEnc := range pokemonEncounters {
		fmt.Println(pokemonEnc.Pokemon.Name)
	}

	return nil
}
