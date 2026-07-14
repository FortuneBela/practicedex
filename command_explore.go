package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please enter a area name after 'explore'...")
	}

	if len(args) > 1 {
		return fmt.Errorf("Please enter a single word connected by dashes ' - '...")
	}

	response, err := cfg.pokeapiClient.GetAreaPokemon(args[0])
	if err != nil {
		return fmt.Errorf("Error with GetAreaPokemon...")
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Print("Found Pokemon:\n")

	for _, pokemon := range response.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}
	return nil
}
