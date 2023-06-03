package main

import (
	"fmt"
	"strings"
	"errors"
)

func commandInspect(c *config) error {
	if c.arg == "" {
		msg := "(Must include a pokemon to inspect. Enter 'inspect pokemon')"
		err := errors.New(invalidMsg+msg)
		return err
	}
	name := c.arg

	if pokemon, ok := c.pokedex[name]; ok && pokemon.caught {
		pokedexIndex := getIndex(pokemon.Order)
		line := strings.Repeat("-", len(name) + 16)
		fmt.Printf("%s\nPokedex: No.%s %s\n%s\n", line, pokedexIndex, strings.Title(name), line)

		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			stat := strings.ReplaceAll(s.Stat.Name, "-", " ")
			stat = strings.Title(stat)
			fmt.Printf(" > %s: %v\n", stat, s.BaseStat)
		}

		type1 := strings.Title(pokemon.Types[0].Type.Name)
		dual := ""
		type2 := ""
		if len(pokemon.Types) > 1 {
			dual = "Dual-"
			type2 = "/" + strings.Title(pokemon.Types[1].Type.Name)
		}
		fmt.Printf("%sType: %s%s\n", dual, type1, type2)

		fmt.Println("Abilities:")
		for _, a := range pokemon.Abilities {
			ability := strings.ReplaceAll(a.Ability.Name, "-", " ")
			ability = strings.Title(ability)
			fmt.Printf(" > %s\n", ability)
		}
	} else {
		msg := fmt.Sprintf("Error: %s is not in the Pokedex.", name)
		msg = msg+"\n(Must catch a pokemon before its data can be added to the Pokedex)"
		err := errors.New(msg)
		return err
	}
	return nil
}