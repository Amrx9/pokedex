package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		fmt.Printf("Your command was: %s\n", commandName)

	}

}

func cleanInput(text string) []string {
	lowerWord := strings.ToLower(text)
	words := strings.Fields(lowerWord)
	return words
}
