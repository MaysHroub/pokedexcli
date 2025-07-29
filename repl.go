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

		err := cmd.Callback(&config)

		if err != nil {
			fmt.Printf("An error occured while executing %v: %v\n", words[0], err)
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
