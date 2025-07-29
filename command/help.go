package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
)

func Help(c *configuration.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	commands := GetCommands()

	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}

	return nil
}
