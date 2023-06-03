package main

import ("github.com/skovranek/pokedexcli/internal/pokeapi"
	"github.com/skovranek/bufferio"
)

/*
[x] fix tab problems in input
[x] use up arrows for input
[ ] graceful close?
[ ] fix link to bufferio (and bufferio name)
*/

const (
	invalidMsg = "Error: invalid input, please try again.\n"
)

func main() {
	c := &config{
		ontinue:        true,
		ommands:        getCommands(),
		buffer:         bufferio.NewBuffer(),

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
	buffer         bufferio.Buffer
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