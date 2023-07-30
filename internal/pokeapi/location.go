package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := ResponseLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return ResponseLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocations{}, err
	}

	locationsResp := ResponseLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return ResponseLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
