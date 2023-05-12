package main

import (
	"bufio"
	"os"
	"github.com/skovranek/pokedexcli/internal/pokeapi"
)

const (
	invalidMsg = "Error: invalid input, please try again.\n"
)

func main() {
	c := &config{
		ontinue: true,
		scanner: bufio.NewScanner(os.Stdin),
		arg: "",
		encounters: "",
		atchDifficulty: 100, // easy: 0, medium: 100, hard: 200
		pokedex: map[string]Pokemon{},
		nextAreasURL: func(s string) *string { return &s }(pokeapi.LocationAreasURL),
	}

	REPL(c)
}

type Pokemon struct {
	pokeapi.Pokemon
	caught bool `json:"caught"`
}

type config struct {
	ontinue bool

	scanner *bufio.Scanner
	arg string

	encounters string
	atchDifficulty int
	pokedex map[string]Pokemon

	nextAreasURL *string
	prevAreasURL *string
}