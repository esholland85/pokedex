package main

import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	//fmt.Println("exit: Shut down the program\nhelp: Displays input options\nmap: Displays the next 20 locations\nmapb: Displays the previous 20 locations")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
