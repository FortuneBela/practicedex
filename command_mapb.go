package main

import (
	"fmt"
)

func commandMapB(cfg *config) error {
	targetURL := ""
	if cfg.PrevURL == nil {
		fmt.Print("you're on the first page")
		return nil
	}
	targetURL = *cfg.PrevURL

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
