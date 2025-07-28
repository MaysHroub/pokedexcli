package main

import (
	"bufio"
	"fmt"
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

		fmt.Println("Your command was: " + words[0])
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
