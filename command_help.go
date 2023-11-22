package main

import (
	"fmt"
)

// HelpCommand is a command that prints the help message
func commandHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	availableCommands := getCommands()

	for name, command := range availableCommands {
		fmt.Printf("- %s: %s\n", name, command.description)
	}

	return nil
}
