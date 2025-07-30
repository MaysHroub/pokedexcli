package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
	"math"
	"math/rand"
	"time"
)

func Catch(cfg *configuration.Config, pokemonName string) error {
	if pokemonName == "" {
		fmt.Println("Please provide a pokemon name")
		return nil
	}

	client := cfg.HttpClient

	if _, exists := cfg.Pokedex[pokemonName]; exists {
		fmt.Printf("Pokemon %v is already caught\n", pokemonName)
		return nil
	}

	url := pokeapi.BaseUrl + "/pokemon/" + pokemonName

	pokemonInfo, err := pokeapi.FetchData[pokeapi.PokemonInfo](client, &url)

	if err != nil {
		return err
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
	fmt.Println("You can now inspect it with the inspect command")

	return nil
}

// using exponential decay
func getCatchChance(baseExperience int) float64 {
	const baseChance = 0.8
	const decayConstant = 0.8 // or lambda

	catchChance := baseChance * math.Pow(math.E, -float64(decayConstant)*float64(baseExperience)/100.0)
	return catchChance
}
