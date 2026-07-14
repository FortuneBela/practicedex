package main

import "fmt"

func commandHelp(cfg *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	for name, com := range getCommands() {
		fmt.Printf("%s: %s\n", name, com.description)
	}
	return nil
}
