package main

import (
	"fmt"
)

// commandInspect Displays Pokemon information if already caught.
func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon name>")
	}

	pokemon, ok := cfg.caughtPokemons[args[0]]
	if !ok {
		fmt.Printf("You have to catch %s pokemon before you can inspect it.\n", args[0])
		return nil
	}

	fmt.Println(pokemon.ToString())

	return nil
}
