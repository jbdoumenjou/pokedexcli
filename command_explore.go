package main

import (
	"fmt"
)

// https://pokeapi.co/docs/v2#location-areas
// commandMap Displays the names of the next 20 location areas in the Pokemon world.
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location name>")
	}

	location, err := cfg.pokeAPIClient.ExploreLocationArea(args[0])
	if err != nil {
		return fmt.Errorf("get location: %w", err)
	}
	fmt.Printf("Exploring %q\n", location.Name)
	fmt.Println("Found Pokemon:")

	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("  - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
