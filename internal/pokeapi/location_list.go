package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) GetLocationAreas(url string) (locationResponse, error) {
	var body []byte
	body, ok := c.cache.Get(url)
	if !ok {
		response, err := c.httpClient.Get(url)
		if err != nil {
			return locationResponse{}, fmt.Errorf("Failed to GET areas...")
		}

		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return locationResponse{}, fmt.Errorf("ReadAll failed...")
		}
		c.cache.Add(url, body)
	}

	locRes := locationResponse{}
	err := json.Unmarshal(body, &locRes)
	if err != nil {
		return locationResponse{}, fmt.Errorf("failed to Unmarshal...")
	}
	return locRes, nil
}

type locationResponse struct {
	Count   int        `json:"count"`
	NextURL *string    `json:"next"`
	PrevURL *string    `json:"previous"`
	Results []pokeArea `json:"results"`
}

type pokeArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
