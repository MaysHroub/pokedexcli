package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/config"
)

func Inspect(cfg *config.Config, pokemonName string) error {
	pokemonInfo, exists := cfg.Pokedex[pokemonName]

	if !exists {
		fmt.Println("You have not caught that pokemon yet")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemonInfo.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typ := range pokemonInfo.Types {
		fmt.Printf("  - %v\n", typ.Type.Name)
	}

	return nil
}
