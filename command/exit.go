package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/configuration"
	"os"
)

func Exit(c *configuration.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
