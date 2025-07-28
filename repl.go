package main

import (
	"bufio"
	"fmt"
	"github/MaysHroub/pokedexcli/command"
	"os"
	"strings"
)

func startRepl() {
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

		cmd, exists := commands[words[0]]

		if !exists {
			fmt.Println("Unknown command: " + words[0])
			continue
		}

		cmd.Callback()

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
