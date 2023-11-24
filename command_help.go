package main

import (
	"fmt"
	"slices"
)

// HelpCommand is a command that prints the help message
func commandHelp(_ *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	availableCommands := getCommands()

	var commandNames []string
	for name, _ := range availableCommands {
		commandNames = append(commandNames, name)
	}

	// Keep alphabetical order
	slices.Sort(commandNames)
	for _, name := range commandNames {
		command := availableCommands[name]
		fmt.Printf("  - %s: %s\n", command.name, command.description)
	}

	return nil
}
