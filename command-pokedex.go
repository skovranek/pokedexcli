package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func commandPokedex(c *config) error {
	if c.arg != "" {
		msg := "(Command 'pokedex' does not take an input. Enter only 'pokedex')"
		err := errors.New(invalidMsg+msg)
		return err
	}

	fmt.Println("------------\nPokedex List\n------------")

	list := []Pokemon{}
	for _, pokemon := range c.pokedex {
		if pokemon.caught {
			list = append(list, pokemon)
		}
	}

	if len(list) == 0 {
		fmt.Println(" > ...\n(Catch a pokemon to add its data to the Pokedex)")
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Order < list[j].Order
	})

	for _, pokemon := range list {
		index := getIndex(pokemon.Order)
		fmt.Printf("> No.%s %s\n", index, strings.Title(pokemon.Name))
	}
	return nil
}