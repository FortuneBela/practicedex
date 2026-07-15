package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please enter a Pokemon name after 'inspect'...")
	}

	if len(args) > 1 {
		return fmt.Errorf("Please enter the name as a single word...")
	}

	pokemonName := args[0]

	value, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)
	fmt.Println("Stats:")
	for _, s := range value.Stats {
		fmt.Printf("	-%v: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range value.Types {
		fmt.Printf("	- %s\n", t.Type.Name)
	}
	return nil
}
