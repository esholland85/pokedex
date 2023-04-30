package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide the name of the pokemon you're trying to inspect")
	}

	resp, ok := cfg.myPokedex.Get(args[0])
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", resp.Name)
	fmt.Printf("Height: %d\n", resp.Height)
	fmt.Printf("Weight: %d\n", resp.Weight)
	fmt.Printf("Stats:\n")
	for i := 0; i < len(resp.Stats); i++ {
		statName := resp.Stats[i].Stat.Name
		statNumber := resp.Stats[i].BaseStat
		fmt.Printf("  - %s: %d\n", statName, statNumber)
	}
	fmt.Printf("Types:\n")
	for i := 0; i < len(resp.Types); i++ {
		statName := resp.Types[i].Type.Name
		fmt.Printf("  - %s\n", statName)
	}

	return nil
}
