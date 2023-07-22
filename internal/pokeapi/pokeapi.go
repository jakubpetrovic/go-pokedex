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

func NewLocationConfig() *locationConfig {
	conf := locationConfig{}
	fmt.Println(conf)
	return &conf
}

func FetchLocations(url string) []byte {

	res, err := http.Get(url)
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

func GetNextLocations(conf *locationConfig) {

	var body []byte
	if conf.Next == "" {
		body = FetchLocations("https://pokeapi.co/api/v2/location/")
	} else {
		body = FetchLocations(conf.Next)
	}

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(conf.Results); i++ {
		fmt.Printf("\n" + conf.Results[i].Name)
	}
	fmt.Printf("\n")

}

func GetPrevLocations(conf *locationConfig) {

	var body []byte
	if conf.Previous == "" {
		body = FetchLocations("https://pokeapi.co/api/v2/location/")
	} else {
		body = FetchLocations(conf.Previous)
	}

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(conf.Results); i++ {
		fmt.Printf("\n" + conf.Results[i].Name)
	}
	fmt.Printf("\n")

}
