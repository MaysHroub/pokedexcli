package main

import (
	"bufio"
	"fmt"
	"github/MaysHroub/pokedexcli/command"
	"github/MaysHroub/pokedexcli/configuration"
	"github/MaysHroub/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	locationAreaEndPoint := pokeapi.BaseUrl + "/location-area"
	config := configuration.UrlConfig{
		Next: &locationAreaEndPoint,
		Prev: nil,
	}

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

		cmd.Callback(&config)

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
