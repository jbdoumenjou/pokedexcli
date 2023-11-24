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
	Callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon and add it to the Pokedex",
			Callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Lists Pokemons for a location area",
			Callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays Pokemon information if already caught",
			Callback:    commandInspect,
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
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		// Prompt
		fmt.Print("Pokedex >")

		scanner.Scan()
		args := cleanInput(scanner.Text())
		commandName := args[0]
		cmd, ok := commands[commandName]
		if !ok {
			fmt.Printf("Unknown command %q\n", commandName)
			commandHelp(nil)
			continue
		}

		if err := cmd.Callback(cfg, args[1:]...); err != nil {
			fmt.Printf("command %q failed: %s\n", commandName, err.Error())
		}
	}
	return
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
