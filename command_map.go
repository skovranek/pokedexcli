package main

import (
	"fmt"
	"errors"
	"github.com/skovranek/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	if c.arg != "" {
		msg := "(Command 'map' does not take an input. Enter only 'map')"
		err := errors.New(invalidMsg+msg)
		return err
	}

	line := "-----------------"
	fmt.Printf("%s\nThe Next %s Areas\n%s\n", line, pokeapi.LimitParam, line)

	if c.nextAreasURL == nil {
		err := errors.New("Error: cannot go past the end of the map.")
		return err
	}
	return getAreas(c, false)
}