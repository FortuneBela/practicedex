package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please enter a Pokemon name after 'catch'...")
	}

	if len(args) > 1 {
		return fmt.Errorf("Please enter the name as a single word...")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	response, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	catchChance := rand.IntN(response.BaseExperience)
	if catchChance > rand.IntN(response.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemonName)

	} else {
		cfg.caughtPokemon[response.Name] = response
		fmt.Printf("%s was caught!\n", pokemonName)
	}

	return nil
}
