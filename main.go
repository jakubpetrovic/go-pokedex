package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jpetrovic/go-pokedex/internal/pokeapi"
)

type cliCommands struct {
	name        string
	description string
	//callback    func() error
}

func commandHelp(commandMap map[string]cliCommands) {
	fmt.Printf("\n\nWelcome to Pokedex!\n\n")
	for _, n := range commandMap {
		fmt.Printf("%s: %s\n", n.name, n.description)
	}
	fmt.Printf("\n")
}

func commandExit() {
	fmt.Println("See you next time!")
}

func main() {
	m := map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Display a help message",
			//callback:   	commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			//callback:    	commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 aread in the pokemon World",
			//callback:		commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 areas in the pokemon World",
			//callback:		commandMapb
		},
	}

	conf := pokeapi.NewLocationConfig()

	for {

		fmt.Printf("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case m["exit"].name:
			commandExit()
			return
		case m["help"].name:
			commandHelp(m)
		case m["map"].name:
			pokeapi.GetNextLocations(conf)
		case m["mapb"].name:
			pokeapi.GetPrevLocations(conf)
		default:
			fmt.Println("incorrect input")
		}

	}
}
