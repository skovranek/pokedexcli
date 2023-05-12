package pokeapi

import (
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
	"errors"
	"encoding/json"
)

func ExploreArea(name string) (ExploredArea, error) {
	URL := pokeapiURL+"location-area/"+strings.ReplaceAll(name, " ", "-")

	body, ok := cache.Get(name)

	if !ok {
		response, err := http.Get(URL)
		if err != nil {
			return ExploredArea{}, err
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if response.StatusCode > 299 {
			errString := fmt.Sprintf("Error: response failed. status code: %d, body: %s", response.StatusCode, body)
			err = errors.New(errString)
			return ExploredArea{}, err
		}
		if err != nil {
			return ExploredArea{}, err
		}
	}

	area := ExploredArea{}
	err := json.Unmarshal(body, &area)
	if err != nil {
		return ExploredArea{}, err
	}
	if !ok {
		cache.Add(URL, body)
	}

	return area, nil
}

type ExploredArea struct {
	PokemonEncounters []struct {
		Pokemon       struct {
			Name      string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}