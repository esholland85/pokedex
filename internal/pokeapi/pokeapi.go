package pokeapi

import (
	"net/http"
	"time"

	"github.com/esholland85/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

/*
type Response struct {
	Areas []Area `json:"results"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type Area struct {
	Name string `json:"name"`
}
*/

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		cache: pokecache.NewCache(),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
