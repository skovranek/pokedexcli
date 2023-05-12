package pokeapi

import (
	"github.com/skovranek/pokedexcli/internal/pokecache"
)

const (
	pokeapiURL = "https://pokeapi.co/api/v2/"
	LimitParam = "10"
	LocationAreasURL = pokeapiURL+"location-area/?offset=0&limit="+LimitParam
)

var cache pokecache.Cache = pokecache.NewCache()