package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) GetAreaPokemon(name string) (areaResponse, error) {
	targetArea := "https://pokeapi.co/api/v2/location-area/" + name
	var body []byte
	body, ok := c.cache.Get(targetArea)
	if !ok {
		response, err := c.httpClient.Get(targetArea)
		if err != nil {
			return areaResponse{}, fmt.Errorf("Failed to GET areas...")
		}

		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return areaResponse{}, fmt.Errorf("ReadAll failed...")
		}
		c.cache.Add(targetArea, body)
	}

	areaRes := areaResponse{}
	err := json.Unmarshal(body, &areaRes)
	if err != nil {
		return areaResponse{}, fmt.Errorf("failed to Unmarshal...")
	}
	return areaRes, nil
}

type areaResponse struct {
	PokemonEncounters []pokemonList `json:"pokemon_encounters"`
}

type pokemonList struct {
	Pokemon pokemonInfo `json:"pokemon"`
}

type pokemonInfo struct {
	Name string `json:"name"`
}
