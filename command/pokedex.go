package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/config"
)

func Pokedex(cfg *config.Config, param string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("You haven't caught any pokemons yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range cfg.Pokedex {
		fmt.Println("  - " + pokemonName)
	}

	return nil
}
