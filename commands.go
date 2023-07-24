package main

import (
	"fmt"
	"os"
)

func commandHelp(commandMap map[string]cliCommands) error {
	fmt.Println()
	for _, n := range commandMap {
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
	return nil
}

func commandMapb(c *config) error {
	return nil
}
