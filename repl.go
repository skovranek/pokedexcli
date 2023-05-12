package main

import (
	"fmt"
)

func REPL(c *config) {
	line := "----------------------"
	fmt.Printf("\n%s\nWelcome to PokedexCLI!\n%s\n\n", line, line)

	commands := getCommands()
	
	for ; c.ontinue == true; {
		fmt.Print("pokedex > ")

		input := ""
		input, c.arg = getInput(c)
		//fmt.Println()
		command, ok := commands[input]

		if !ok {
			msg := "(Enter 'help' to list available commands)"
			fmt.Println(invalidMsg+msg)
		} else {
			err := command.callback(c)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}