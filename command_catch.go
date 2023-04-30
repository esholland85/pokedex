package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide the name of the pokemon you're trying to catch")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		fmt.Println("Are you sure that's the name of a pokemon?")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	xp := resp.BaseExperience
	catchNumber := rand.Intn(300)
	if catchNumber > xp {
		fmt.Printf("Succesfully captured %s!\n", args[0])
		cfg.myPokedex.Add(resp)
		return nil
	}

	fmt.Printf("%s escaped!\n", args[0])

	return nil
}
