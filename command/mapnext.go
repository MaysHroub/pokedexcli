package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
)

func MapNext(cfg *configuration.Config) error {
	httpClient := cfg.HttpClient

	locationAreaResp, err := httpClient.GetLocationAreaResponse(*cfg.NextUrl)

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
