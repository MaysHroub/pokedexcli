package configuration

import "github/MaysHroub/pokedexcli/internal/pokeapi"

type Config struct {
	HttpClient *pokeapi.Client
	NextUrl    *string
	PrevUrl    *string
}
