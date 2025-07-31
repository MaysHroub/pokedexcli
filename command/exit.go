package command

import (
	"fmt"
	"github/MaysHroub/pokedexcli/config"
	"os"
)

func Exit(c *config.Config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
