package pokeapi

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"errors"
	"encoding/json"
)

func GetLocationAreas(URL string) (LocationAreas, error) {
	body, ok := cache.Get(URL)

	if !ok {
		response, err := http.Get(URL)
		if err != nil {
			return LocationAreas{}, err
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if response.StatusCode > 299 {
			msg := fmt.Sprintf("Error: response failed. status-code: %d, body: %s", response.StatusCode, body)
			err = errors.New(msg)
			return LocationAreas{}, err
		}
		if err != nil {
			return LocationAreas{}, err
		}
	}

	areas := LocationAreas{}
	err := json.Unmarshal(body, &areas)
	if err != nil {
		return LocationAreas{}, err
	}

	if !ok {
		cache.Add(URL, body)
	}

	return areas, nil
}

type LocationAreas struct {
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []struct {
		Name string   `json:"name"`
	}                 `json:"results"`
}