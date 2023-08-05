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
	caughtPokemons      map[string]pokeapi.Pokemon
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config, string) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var commandArgument string
		if len(words) != 1 {
			commandArgument = words[1]
		} else {
			commandArgument = "noArgument"
		}

		command, exists := commandLibrary()[commandName]

		if exists {
			err := command.callback(config, commandArgument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
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
		"explore": {
			name:        "explore",
			description: "displays pokemons in given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemons that were already caught",
			callback:    commandPokedex,
		},
	}
	return m
}
