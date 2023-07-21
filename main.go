package main

import "fmt"

func main() {
	for {
		var input string = ""
		fmt.Println("pokedex > ")
		fmt.Scanln(&input)

		if input == "help" {
			fmt.Println("NO HALP!! pres any key to continue...")
			fmt.Scanln()
		}

		if input == "exit" {
			return
		}

	}
}
