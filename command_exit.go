package main

import (
	"fmt"
)

func commandExit(c *config) error {
	if c.arg == "y" || c.arg == "yes" {
		c.ontinue = false
		fmt.Println("Goodbye!")
		return nil
	}

	fmt.Print("Exit? (y/n) ")

	input, _, err := getInput(c)
	if err != nil {
		return err
	}

	if input == "y" || input == "yes" {
		c.ontinue = false
		fmt.Println("\nGoodbye!")
	}
	return nil
}