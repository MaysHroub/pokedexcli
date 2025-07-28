package command

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
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
	}
}
