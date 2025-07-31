package command

import "github/MaysHroub/pokedexcli/config"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *config.Config, param string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
		"mapn": {
			Name:        "mapn",
			Description: "Displays next 20 location areas",
			Callback:    MapNext,
		},
		"mapp": {
			Name:        "mapp",
			Description: "Displays previous 20 location areas",
			Callback:    MapPrev,
		},
		"explore": {
			Name:        "explore <location-name | location-id>",
			Description: "Displays all pokemon names in given location area",
			Callback:    Explore,
		},
		"catch": {
			Name:        "catch <pokemon-name>",
			Description: "Catches a pokemon; the higher the base experience of the pokemon, the harder to catch",
			Callback:    Catch,
		},
		"inspect": {
			Name:        "inspect <pokemon-name>",
			Description: "Displays name, height, weight, stats, and type(s) of the caught pokemon",
			Callback:    Inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays name of caught pokemons",
			Callback:    Pokedex,
		},
	}
}
