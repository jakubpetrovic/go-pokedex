package pokeapi

import (
	"fmt"
)

type locationSlice struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetNextMap(locationIndex int) []string {

	// res, err := http.Get("https://pokeapi.co/api/v2/location")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)

	fmt.Print(locationIndex)
	locationsSlice := []string{"location1", "location2"}

	return locationsSlice
}

func GetPrevMap() {
	fmt.Println("Getting previous maps")
	// if not prev, print error
}
