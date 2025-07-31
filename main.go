package main

import (
	"github/MaysHroub/pokedexcli/config"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
	"github/MaysHroub/pokedexcli/repl"
	"time"
)

func main() {
	timeout := 10 * time.Second
	cacheInterval := 15 * time.Second

	httpClient := pokeapi.NewClient(timeout, cacheInterval)
	pokedex := make(map[string]pokeapi.PokemonInfo)

	cfg := config.Config{
		HttpClient: &httpClient,
		Pokedex:    pokedex,
	}

	repl.StartRepl(&cfg)
}
