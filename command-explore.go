package main

import (
	"fmt"
	"strings"
	"errors"
	"github.com/skovranek/pokedexcli/internal/pokeapi"
)

func commandExplore(c *config) error {
	if c.arg == "" {
		err := errors.New("(Must include an area to explore. Enter 'explore area')")
		return err
	}

	area := c.arg
	line := strings.Repeat("-", len(area) + 13)
	fmt.Printf("%s\nExploring %s...\n%s\n", line, strings.Title(area), line)

	exploredArea, err := pokeapi.ExploreArea(area)
	if err != nil {
		msg := "(Must include the correct name of an area to explore. Enter 'explore area')"
		err := errors.New(invalidMsg+msg)
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, p := range exploredArea.PokemonEncounters {
		name := p.Pokemon.Name
		c.encounters = fmt.Sprintf("%s%s,", c.encounters, name)
		fmt.Printf(" > %s\n", name)
	}
	return nil
}