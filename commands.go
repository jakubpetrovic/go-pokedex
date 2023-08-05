package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func commandHelp(c *config, str string) error {
	fmt.Println()
	for _, n := range commandLibrary() {
		fmt.Printf("%s: %s\n", n.name, n.description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *config, str string) error {
	os.Exit(0)
	return nil
}

func commandMap(c *config, str string) error {
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

func commandMapb(c *config, str string) error {
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

func commandExplore(c *config, str string) error {
	locationInfoResp, err := c.pokeapiClient.ExploreLocation(str)
	if err != nil {
		return err
	}

	for _, key := range locationInfoResp.PokemonEncounters {
		fmt.Printf(" - %s\n", key.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *config, str string) error {
	pokemonResp, err := c.pokeapiClient.GetPokemon(str)

	if err != nil {
		return err
	}

	baseCatcgRate := 0.3
	scalingFactor := 0.001
	catchRate := baseCatcgRate - scalingFactor*float64(pokemonResp.BaseExperience)
	catchProbability := math.Max(0, math.Min(1, catchRate))

	fmt.Printf("Pokemon base experience: %v\n", pokemonResp.BaseExperience)
	fmt.Printf("Pokemon catch chance: %v\n", catchProbability)
	catchAttempt := rand.Float64()
	fmt.Printf("Catch attempt roll: %v\n", catchAttempt)

	if catchAttempt > catchProbability {
		fmt.Printf("\n !%s escaped! \n\n", pokemonResp.Name)
	} else {
		c.caughtPokemons[pokemonResp.Name] = pokemonResp
		fmt.Printf("\n %s was caught! \n\n", pokemonResp.Name)
	}

	return nil
}
