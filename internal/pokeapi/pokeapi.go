package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locationConfig struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var conf locationConfig

func FetchLocations(url *string) []byte {

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func GetInitLocations() {
	var initUrl string = "https://pokeapi.co/api/v2/location/"
	body := FetchLocations(&initUrl)

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}
}

func GetNextLocations() {

	body := FetchLocations(&conf.Next)

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(conf.Results); i++ {
		fmt.Printf("\n" + conf.Results[i].Name)
	}
	fmt.Printf("\n")

	fmt.Println(conf.Next)
	fmt.Println(conf.Previous)
	fmt.Println(conf)

}

func GetPrevLocations() {

	body := FetchLocations(&conf.Previous)

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(conf.Results); i++ {
		fmt.Printf("\n" + conf.Results[i].Name)
	}
	fmt.Printf("\n")

	fmt.Println(conf.Next)
	fmt.Println(conf.Previous)
	fmt.Println(conf)

}
