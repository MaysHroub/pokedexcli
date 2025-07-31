package main

import (
	"bufio"
	"fmt"
	"github/MaysHroub/pokedexcli/command"
	"github/MaysHroub/pokedexcli/config"
	"os"
	"strings"
)

func startRepl(cfg *config.Config) {
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
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
