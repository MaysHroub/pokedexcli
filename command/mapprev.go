package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
)

func MapPrev(cfg *configuration.Config, param string) error {
	httpClient := cfg.HttpClient

	locationAreaResp, err := httpClient.GetLocationAreaResponse(cfg.PrevUrl)

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
