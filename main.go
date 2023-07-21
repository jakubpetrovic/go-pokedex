package main

import "fmt"

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

func main() {
	m := map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Display a help message",
			//callback:    commandHelp(),
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			//callback:    commandExit(),
		},
	}

	for {
		var input string = ""
		fmt.Printf("pokedex > ")
		fmt.Scanln(&input)

		switch input {
		case "exit":
			fmt.Println("see you next time")
			return
		case "help":
			commandHelp(m)
		default:
			fmt.Println("incorrect output")
		}

	}
}
