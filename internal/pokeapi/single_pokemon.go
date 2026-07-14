package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) GetPokemon(name string) (Pokemon, error) {
	targetPokemon := "https://pokeapi.co/api/v2/pokemon/" + name
	var body []byte
	body, ok := c.cache.Get(targetPokemon)
	if !ok {
		response, err := c.httpClient.Get(targetPokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Failed to get Pokemon...")
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			return Pokemon{}, fmt.Errorf("failed to find this Pokemon")
		}

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("ReadAll failed...")
		}
		c.cache.Add(targetPokemon, body)
	}

	pokemonRes := Pokemon{}
	err := json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to Unmarshal...")
	}
	return pokemonRes, nil
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}
