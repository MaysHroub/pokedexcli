package command

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Exit,
		},
		"help": cliCommand{
			name: "help"
			description: "Displays a help message",
			callback: Help,
		},
	}
}
