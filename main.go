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

	cfg := configuration.Config{
		HttpClient: &httpClient,
	}

	startRepl(&cfg)
}
