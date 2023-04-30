package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	myEntries := cfg.myPokedex.ListEntries()
	if len(myEntries) < 1 {
		return errors.New("your pokedex is blank")
	}

	fmt.Println("Your Pokedex:")

	for i := 0; i < len(myEntries); i++ {
		fmt.Printf("  - %s\n", myEntries[i])
	}

	return nil
}
