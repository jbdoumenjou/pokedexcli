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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for name, command := range getCommands() {
		fmt.Printf("- %s: %s\n", name, command.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		commandName := cleanInput(scanner.Text())[0]
		command, ok := commands[commandName]
		if !ok {
			fmt.Printf("Unknown command %q\n", commandName)
			commandHelp()
			continue
		}

		command.callback()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
