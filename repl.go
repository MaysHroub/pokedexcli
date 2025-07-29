package main

import (
	"bufio"
	"fmt"
	"github/MaysHroub/pokedexcli/command"
	"github/MaysHroub/pokedexcli/configuration"
	"os"
	"strings"
)

func startRepl(cfg *configuration.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commands := command.GetCommands()

		commandName := words[0]

		cmd, exists := commands[commandName]

		if !exists {
			fmt.Println("Unknown command: " + commandName)
			continue
		}

		commandParam := ""

		if len(words) >= 2 {
			commandParam = words[1]
		}

		err := cmd.Callback(cfg, commandParam)

		if err != nil {
			fmt.Printf("An error occured while executing %v: %v\n", commandName, err)
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
