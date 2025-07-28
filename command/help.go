package command

import "fmt"

func Help() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	commands := GetCommands()

	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}
