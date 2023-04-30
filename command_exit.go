package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config, args ...string) error {
	fmt.Println("Shutting down.")
	os.Exit(0)
	return nil
}
