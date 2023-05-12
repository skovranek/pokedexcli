package main

import (
	"fmt"
	"errors"
)

func commandHelp(c *config) error {
	if c.arg != "" {
		msg := "(Command 'help' does not take an input. Enter only 'help')"
		err := errors.New(invalidMsg+msg)
		return err
	}

	fmt.Println("--------------------\nPokedex Instructions\n--------------------")

	commands := getCommands()
/*
	Instructions:
	1) Map/mapb
	2) Explore
	3) Catch
	4) Inspect
	5) Pokedex
	6) Exit
*/
	list := make([]command, len(commands))
	for _, command := range commands {
		list[command.order] = command
	}

	for _, command := range list {
		if command.order == 0 {continue}
		fmt.Printf("%v) %s - %s\n", command.order, command.name, command.description)
	}

	return nil
}