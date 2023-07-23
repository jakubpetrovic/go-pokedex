package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jpetrovic/go-pokedex/internal/pokeapi"
	"github.com/jpetrovic/go-pokedex/internal/pokecache"
)

type cliCommands struct {
	name        string
	description string
	//callback    func() error
}

func commandHelp(commandMap map[string]cliCommands) {
	fmt.Println()
	for _, n := range commandMap {
		fmt.Printf("%s: %s\n", n.name, n.description)
	}
	fmt.Println()
}

func commandExit() {
	fmt.Println("See you next time!")
}

func commandLibrary() map[string]cliCommands {
	m := map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Display a help message",
			//callback:    commandHelp(m),
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
	return m
}

func main() {

	cmdMap := commandLibrary()
	locationConf := pokeapi.NewLocationConfig()
	cache := pokecache.NewCache(time.Duration(1000000000))

	if cache == nil {
		fmt.Println("No Cache was create")
	}

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
