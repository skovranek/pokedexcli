package main

import (
	"strings"
)

func getInput(c *config) (string, string, error) {
	input, err := c.buffer.GetInput()
	if err != nil {
		return "", "", err
	}

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, "-", " ")
	inputs := strings.Fields(input)
	first := inputs[0]

	second := ""
	if len(inputs) > 1 {
		second = strings.Join(inputs[1:], " ")
	}

	return first, second, nil
}