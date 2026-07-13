package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	slicedText := strings.Fields(text)
	return slicedText
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cInput := cleanInput(input)
		if len(cInput) == 0 {
			continue
		}
		firstWord := cInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
