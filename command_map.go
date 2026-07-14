package main

import (
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	targetURL := ""
	if cfg.NextURL == nil {
		targetURL = "https://pokeapi.co/api/v2/location-area"
	} else {
		targetURL = *cfg.NextURL
	}

	locRes, err := cfg.pokeapiClient.GetLocationAreas(targetURL)
	if err != nil {
		return fmt.Errorf("Failed to GetLocationAreas...")
	}

	cfg.NextURL = locRes.NextURL
	cfg.PrevURL = locRes.PrevURL

	for _, area := range locRes.Results {
		fmt.Println(area.Name)
	}
	return nil
}
