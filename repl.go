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
	Callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations areas in the Pokemon world",
			Callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "commandExit the Pokedex",
			Callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		// Prompt
		fmt.Print("Pokedex >")

		scanner.Scan()
		commandName := cleanInput(scanner.Text())[0]
		cmd, ok := commands[commandName]
		if !ok {
			fmt.Printf("Unknown command %q\n", commandName)
			commandHelp(nil)
			continue
		}

		if err := cmd.Callback(cfg); err != nil {
			fmt.Printf("command %q failed: %s\n", commandName, err.Error())
		}
	}
	return
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
