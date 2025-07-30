package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
)

func MapNext(cfg *configuration.Config, param string) error {
	client := cfg.HttpClient

	if cfg.NextUrl == nil {
		url := pokeapi.BaseUrl + "/location-area"
		cfg.NextUrl = &url
	}

	locationAreaResp, err := pokeapi.FetchData[pokeapi.LocationAreaResponse](client, cfg.NextUrl)
	if err != nil {
		return err
	}

	cfg.NextUrl = locationAreaResp.Next
	cfg.PrevUrl = locationAreaResp.Previous

	for _, loc := range locationAreaResp.LocationAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
