package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FortuneBela/practicedex/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cInput := cleanInput(input)
		if len(cInput) == 0 {
			continue
		}
		command, ok := getCommands()[cInput[0]]
		if ok {
			err := command.callback(cfg, cInput[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	slicedText := strings.Fields(text)
	return slicedText
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokemon World.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon World.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Adding a location name after explore, will show all pokemon that can be found in that location.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Adding a Pokemon name after catch, will give you a chance at catching that Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "catch",
			description: "Inspect a pokemon you have caught.",
			callback:    commandInspect,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	NextURL       *string
	PrevURL       *string
	caughtPokemon map[string]pokeapi.Pokemon
}
