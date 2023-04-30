package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/esholland85/pokedex/internal/color"
)

func startRepl(cfg *config) {
	myScanner := bufio.NewScanner(os.Stdin)

	prompt := fmt.Sprintf(color.Yellow + "Pokedex > " + color.Reset)
	fmt.Print(prompt)

	for myScanner.Scan() {

		cleaned := cleanInput(myScanner.Text())

		if len(cleaned) == 0 {
			fmt.Print(prompt)
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Not a known command, try 'help'")
			fmt.Print(prompt)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(prompt)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays input options",
			callback:    callbackHelp,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Review all caught pokemon",
			callback:    callbackPokedex,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_number}",
			description: "See the pokemon in an area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Attempt to catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon name}",
			description: "View stats of a caught pokemon",
			callback:    callbackInspect,
		},
		"exit": {
			name:        "exit",
			description: "Shut down the program",
			callback:    callbackExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
