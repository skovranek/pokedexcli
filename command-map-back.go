package main

import (
	"fmt"
	"errors"
	"github.com/skovranek/pokedexcli/internal/pokeapi"
)

func commandMapBack(c *config) error {
	if c.arg != "" {
		msg := "(Command 'mapb' does not take an input. Enter only 'mapb')"
		err := errors.New(invalidMsg+msg)
		return err
	}

	line := "---------------------"
	fmt.Printf("%s\nThe Previous %s Areas\n%s\n", line, pokeapi.LimitParam, line)

	if c.prevAreasURL == nil {
		err := errors.New("Error: cannot go back before the start of the map.")
		return err
	}

	return getAreas(c, true)
}