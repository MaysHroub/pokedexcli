package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"math"
	"math/rand"
	"time"
)

func Catch(cfg *configuration.Config, pokemonName string) error {
	client := cfg.HttpClient

	if _, exists := cfg.Pokedex[pokemonName]; exists {
		fmt.Printf("Pokemon %v is already caught\n", pokemonName)
		return nil
	}

	pokemonInfo, err := client.GetPokemonInfo(pokemonName)

	if err != nil {
		return fmt.Errorf("An error occurred while fetching the info of pokemon %v: %w", pokemonName, err)
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	catchChance := getCatchChance(pokemonInfo.BaseExperience)
	roll := rand.Float64()

	time.Sleep(time.Second)

	if roll > catchChance {
		fmt.Printf("%v escaped!\n", pokemonName)
		return nil
	}

	cfg.Pokedex[pokemonName] = pokemonInfo

	fmt.Printf("%v was caught!\n", pokemonName)
	return nil
}

// using exponential decay
func getCatchChance(baseExperience int) float64 {
	const baseChance = 0.8
	const decayConstant = 0.9 // or lambda

	catchChance := baseChance * math.Pow(math.E, -float64(decayConstant)*float64(baseExperience)/100.0)
	return catchChance
}
