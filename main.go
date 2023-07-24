package main

import (
	"time"

	"github.com/jpetrovic/go-pokedex/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.Client(5*time.Second, time.Minute*5)
	conf := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(conf)
}
