package main

import (
	"github.com/skovranek/pokedexcli/internal/pokeapi"
	"strings"
	"fmt"
)

func getAreas(c *config, back bool) error {
	URL := *c.nextAreasURL
	if back {
		URL = *c.prevAreasURL
	}

	areas, err := pokeapi.GetLocationAreas(URL)
	if err != nil {
		return err
	}

	c.nextAreasURL = areas.Next
	c.prevAreasURL = areas.Previous

	for _, a := range areas.Results {
		area := strings.ReplaceAll(a.Name, "-", " ")
		area = strings.Title(area)
		fmt.Println(area)
	}
	return nil
}