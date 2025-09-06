package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerWord := strings.ToLower(text)
	words := strings.Fields(lowerWord)
	return words
}
