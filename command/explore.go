package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
)

func Explore(cfg *configuration.Config, locationName string) error {
	client := cfg.HttpClient

	url := pokeapi.BaseUrl + "/location-area/" + locationName

	pokemonResp, err := pokeapi.FetchData[pokeapi.PokemonResponse](client, &url)
	if err != nil {
		return err
	}

	for _, pokemonEnc := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemonEnc.Pokemon.Name)
	}

	return nil
}
