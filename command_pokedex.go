package main

import (
	"fmt"
	"slices"
)

// commandInspect Displays the Pokedex.
func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.caughtPokemons) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	var names []string
	for name, _ := range cfg.caughtPokemons {
		names = append(names, name)
	}

	// Keep alphabetical order.
	slices.Sort(names)
	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}
