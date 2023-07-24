package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jpetrovic/go-pokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type locationConfig struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// POKEAPI CLIENT
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func NewLocationConfig() *locationConfig {
	conf := locationConfig{}
	return &conf
}

func FetchLocations(url string) []byte {

	res, err := http.Get(baseURL)

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

	if conf.Next == nil {
		body = FetchLocations(baseURL)
	} else {
		body = FetchLocations(*conf.Next)
	}

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(conf.Results); i++ {
		fmt.Printf(conf.Results[i].Name + "\n")
	}

}

func GetPrevLocations(conf *locationConfig) {

	var body []byte

	if conf.Previous == nil {
		body = FetchLocations(baseURL)
	} else {
		body = FetchLocations(*conf.Previous)
	}

	err := json.Unmarshal(body, &conf)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(conf.Results); i++ {
		fmt.Printf(conf.Results[i].Name + "\n")
	}
}
