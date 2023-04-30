package main

import (
	"errors"
	"fmt"
	"strconv"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println(len(args))
		return errors.New("provide a single location number")
	}
	if cfg.nextLocationAreaURL == nil && cfg.previousLocationAreaURL == nil {
		return errors.New("you have not mapped any locations yet, do that before exploring")
	}

	tempNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	locationNumber := tempNumber - 1
	if locationNumber > len(cfg.currentLocations.Results)-1 || locationNumber < 0 {
		return errors.New("input a number from the current mapped location")
	}
	locationName := cfg.currentLocations.Results[locationNumber].Name
	resp, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", locationName)
	for i := 0; i < len(resp.PokemonEncounters); i++ {
		fmt.Println(titleCleanup(resp.PokemonEncounters[i].Pokemon.Name))
	}

	return nil
}
