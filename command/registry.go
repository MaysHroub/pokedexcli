package command

import "github/MaysHroub/pokedexcli/configuration"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *configuration.Config, param string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": CliCommand{
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": CliCommand{
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
		"mapn": CliCommand{
			Name:        "mapn",
			Description: "Displays next 20 location areas",
			Callback:    MapNext,
		},
		"mapp": CliCommand{
			Name:        "mapp",
			Description: "Displays previous 20 location areas",
			Callback:    MapPrev,
		},
		"explore": CliCommand{
			Name:        "explore <location-name | location-id>",
			Description: "Displays all pokemon names in given location area",
			Callback:    Explore,
		},
		"catch": CliCommand{
			Name:        "catch <pokemon-name>",
			Description: "Catches a pokemon given its name; the higher the base experience of the pokemon, the harder to catch",
			Callback:    Catch,
		},
	}
}
