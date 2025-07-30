package main

import (
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	timeout := 10 * time.Second
	cacheInterval := 15 * time.Second

	httpClient := pokeapi.NewClient(timeout, cacheInterval)
	pokedex := make(map[string]pokeapi.PokemonInfo)

	cfg := configuration.Config{
		HttpClient: &httpClient,
		Pokedex:    pokedex,
	}

	startRepl(&cfg)
}
