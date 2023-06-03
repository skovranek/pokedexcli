package main

import (
	"github.com/skovranek/pokedexcli/internal/pokeapi"
	"github.com/skovranek/buftermio"
)

/*
TODO:
[ ] graceful close
*/

const (
	invalidMsg = "Error: invalid input, please try again.\n"
)

func main() {
	c := &config{
		ontinue:        true,
		ommands:        getCommands(),
		buffer:         buftermio.NewBuffer(),

		arg:            "",
		encounters:     "",
		atchDifficulty: 100, // easy: 0, medium: 100, hard: 200
		pokedex:        map[string]Pokemon{},
		nextAreasURL:   func(s string) *string { return &s }(pokeapi.LocationAreasURL),
	}

	REPL(c)
}

type config struct {
	ontinue        bool
	ommands        map[string]command
	buffer         buftermio.Buffer
	arg            string

	encounters     string
	atchDifficulty int
	pokedex        map[string]Pokemon

	nextAreasURL   *string
	prevAreasURL   *string
}

type Pokemon struct {
	pokeapi.Pokemon
	caught bool `json:"caught"`
}