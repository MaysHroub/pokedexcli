package command

import "github/MaysHroub/pokedexcli/configuration"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *configuration.UrlConfig) error
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
			Name:        "map-next",
			Description: "Displays next 20 location areas",
			Callback:    MapNext,
		},
		"mapp": CliCommand{
			Name:        "map-previous",
			Description: "Displays previous 20 location areas",
			Callback:    MapPrev,
		},
	}
}
