package main

import (
	"fmt"
	"strings"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.currentLocations = resp

	for i := 0; i < len(resp.Results); i++ {
		stringHolder := fmt.Sprintf("%d) %s", i+1, titleCleanup(resp.Results[i].Name))
		fmt.Println(stringHolder)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.currentLocations = resp

	for i := 0; i < len(resp.Results); i++ {
		stringHolder := fmt.Sprintf("%d) %s", i+1, titleCleanup(resp.Results[i].Name))
		fmt.Println(stringHolder)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func titleCleanup(title string) string {
	myArray := strings.Split(title, "-")
	result := ""
	for i := 0; i < len(myArray); i++ {
		result += strings.ToUpper(myArray[i][0:1]) + myArray[i][1:] + " "
	}

	return result
}
