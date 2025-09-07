package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cm := range register {
		fmt.Printf("%s: %s\n", cm.name, cm.description)
	}
	fmt.Println()
	return nil
}
