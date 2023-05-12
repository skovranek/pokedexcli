package main

import (
	"strings"
)

func getInput(c *config) (string, string) {
	c.scanner.Scan()
	input := c.scanner.Text()
	//if err := c.scanner.Err(); err != nil {
	//	return err
	//}

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	inputs := strings.SplitN(input, " ", 2)

	first := inputs[0]
	second := ""
	if len(inputs) == 2 {
		second = strings.ReplaceAll(inputs[1], "-", " ")
	}

	return first, second
}