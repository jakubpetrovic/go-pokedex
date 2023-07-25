package main

import (
	"fmt"
	"os"
)

func commandHelp(c *config) error {
	fmt.Println()
	for _, n := range commandLibrary() {
		fmt.Printf("%s: %s\n", n.name, n.description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *config) error {
	os.Exit(0)
	return nil
}

func commandMap(c *config) error {
	locationResp, err := c.pokeapiClient.ListLocations(c.nextLocationURL)

	if err != nil {
		return err
	}

	c.nextLocationURL = locationResp.Next
	c.previousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	locationResp, err := c.pokeapiClient.ListLocations(c.previousLocationURL)

	if err != nil {
		return err
	}

	c.nextLocationURL = locationResp.Next
	c.previousLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
