package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/config"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
)

func MapPrev(cfg *config.Config, param string) error {
	client := cfg.HttpClient

	if cfg.PrevUrl == nil {
		url := pokeapi.BaseUrl + "/location-area"
		cfg.PrevUrl = &url
	}

	locationAreaResp, err := pokeapi.FetchData[pokeapi.LocationAreaResponse](client, cfg.PrevUrl)
	if err != nil {
		return err
	}

	cfg.PrevUrl = locationAreaResp.Previous
	cfg.NextUrl = locationAreaResp.Next

	for _, loc := range locationAreaResp.LocationAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
