package main

import (
	"fmt"
	"strings"
	"errors"
	"math"
	"math/rand"
	"github.com/skovranek/pokedexcli/internal/pokeapi"
)

func commandCatch(c *config) error {
	if c.arg == "" {
		msg := "(Must include a pokemon to catch. Enter 'catch pokemon')"
		err := errors.New(invalidMsg+msg)
		return err
	}

	name := c.arg

	if !strings.Contains(c.encounters, name) {
		msg := fmt.Sprintf("Must find a %s after exploring an area in order to try catching it.\n", name)
		err := errors.New(msg + "(Enter 'help' for instructions)")
		return err
	}

	nameCap := strings.Title(name)
	line := strings.Repeat("-", len(name) + 26)
	fmt.Printf("%s\nThrowing a pokeball at %s...\n%s\n", line, nameCap, line)

	if _, ok := c.pokedex[name]; !ok {
		newPokemon, err := pokeapi.GetPokemon(name)
		if err != nil {
			msg := "(Must include the correct name of an pokemon to catch. Enter 'catch pokemon')"
			err := errors.New(invalidMsg+msg)
			return err
		}
		c.pokedex[name] = Pokemon{
			newPokemon, //Pokemon struct field:values
			false,      //caught bool
		}
	}
	pokemon := c.pokedex[name]

	will := float64(pokemon.BaseExperience + c.atchDifficulty)
	will = math.Log(will)
	will = math.Pow(will, 2)
	
	throw := float64(rand.Intn(50))
	fmt.Printf("Pokeball throw: %v vs. %s's willpower: %.1f\n", throw, nameCap, will)

	if throw > will {
		fmt.Printf("Gotcha! %s was caught!\n", nameCap)
		if !pokemon.caught {
			pokemon.caught = true
			c.pokedex[name] = pokemon

			fmt.Printf("%s's data was newly added to the Pokedex!\n", nameCap)
			fmt.Printf("pokedex > inspect %s\n", name)
			commandInspect(c)
		} else {
			fmt.Printf("%s's data is already in the Pokedex.\n", nameCap)
		}
	} else {
		fmt.Printf("%s escaped!\n", nameCap)
	}
	return nil
}