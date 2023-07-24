package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jpetrovic/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(config *config) {

	// cache := pokecache.NewCache(time.Duration(1000000000))

	// if &cache == nil {
	// 	fmt.Println("No Cache was created")
	// }
	// cmdMap := commandLibrary()
	// locationConf := pokeapi.NewLocationConfig()

	for {

		fmt.Printf("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case cmdMap["exit"].name:
			commandExit()
			return
		case cmdMap["help"].name:
			commandHelp(cmdMap)
		case cmdMap["map"].name:
			pokeapi.GetNextLocations(locationConf)
		case cmdMap["mapb"].name:
			pokeapi.GetPrevLocations(locationConf)
		default:
			fmt.Println("incorrect input")
		}
	}
}

func commandLibrary() map[string]cliCommands {
	m := map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 aread in the pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 areas in the pokemon World",
			callback:    commandMapb,
		},
	}
	return m
}
