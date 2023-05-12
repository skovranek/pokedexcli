package main

import (
	"fmt"
	"github.com/skovranek/pokedexcli/internal/pokeapi"
)

type command struct {
	name		string
	order		int
	description	string
	callback	func(*config) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:		 "help",
			order:		 0,
			description: "Lists available commands.",
			callback:	 commandHelp,
		},
		"map": {
			name:		 "map",
			order:		 1,
			description: fmt.Sprintf("Displays the next %s areas.", pokeapi.LimitParam),
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			order:		 2,
			description: fmt.Sprintf("Displays back the previous %s areas.", pokeapi.LimitParam),
			callback:	 commandMapBack,
		},
		"explore": {
			name:		 "explore <area>",
			order:		 3,
			description: "Explore an area to search for pokemon.",
			callback:	 commandExplore,
		},
		"catch": {
			name: 		 "catch <pokemon>",
			order:		 4,
			description: "Catch a pokemon and add it to the Pokedex.",
			callback:	 commandCatch,
		},
		"inspect": {
			name:		 "inspect <pokemon>",
			order:		 5,
			description: "Check the Pokedex for data about a pokemon.",
			callback:	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			order:		 6,
			description: "Display a list of the pokemon in the Pokedex.",
			callback:	 commandPokedex,
		},
		"exit": {
			name:		 "exit",
			order:		 7,
			description: "Exit the Pokedex.",
			callback:	 commandExit,
		},
	}
}