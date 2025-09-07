package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var register map[string]cliCommand

func startRepl() {

	register = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		if _, ok := register[commandName]; ok {
			if err := register[commandName].callback(); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func cleanInput(text string) []string {
	lowerWord := strings.ToLower(text)
	words := strings.Fields(lowerWord)
	return words
}
