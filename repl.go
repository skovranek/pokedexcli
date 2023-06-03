package main

import (
	"fmt"
)

func REPL(c *config) {
	line := "----------------------"
	fmt.Printf("\n%s\nWelcome to PokedexCLI!\n%s\n\n", line, line)
	
	for ; c.ontinue == true; {
		fmt.Print("pokedex > ")

		input, arg, err := getInput(c)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.arg = arg

		command, ok := c.ommands[input]

		if !ok {
			msg := "(Enter 'help' to list available commands)"
			fmt.Println(invalidMsg+msg)
		} else {
			err = command.callback(c)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}