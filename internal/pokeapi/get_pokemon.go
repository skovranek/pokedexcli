package pokeapi

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"errors"
	"encoding/json"
)

func GetPokemon(name string) (Pokemon, error) {
	response, err := http.Get(pokeapiURL+"pokemon/"+name)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode > 299 {
		errString := fmt.Sprintf("Error: response failed. status code: %d, body: %s", response.StatusCode, body)
		err = errors.New(errString)
		return Pokemon{}, err
	}
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

type Pokemon struct {
	Name           string  `json:"name"`
	Order 		   int     `json:"order"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []struct {
		BaseStat   int     `json:"base_stat"`
		Effort     int     `json:"effort"`
		Stat       struct {
			Name   string  `json:"name"`
			URL    string  `json:"url"`
		} 				   `json:"stat"`
	} 				       `json:"stats"`
	Types          []struct {
		Slot       int     `json:"slot"`
		Type       struct {
			Name   string  `json:"name"`
			URL    string  `json:"url"`
		}                  `json:"type"`
	}                      `json:"types"`
	Abilities      []struct {
		Ability    struct {
			Name   string  `json:"name"`
			URL    string  `json:"url"`
		}                  `json:"ability"`
		IsHidden   bool    `json:"is_hidden"`
		Slot       int     `json:"slot"`
	}                      `json:"abilities"`
	//LocationAreaEncounters string `json:"location_area_encounters"`
}