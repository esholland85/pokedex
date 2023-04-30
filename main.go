package main

import "github.com/esholland85/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	myPokedex               pokeapi.Pokedex
	currentLocations        pokeapi.LocationAreasResp
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		myPokedex:     pokeapi.NewPokedex(),
	}
	startRepl(&cfg)
}
