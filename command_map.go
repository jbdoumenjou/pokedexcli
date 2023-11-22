package main

import (
	"fmt"
)

// https://pokeapi.co/docs/v2#location-areas
// commandMap Displays the names of the next 20 location areas in the Pokemon world.
func commandMap(cfg *config) error {
	locations, err := cfg.pokeAPIClient.ListLocationAreas(cfg.NextLocationAreaUrl)
	if err != nil {
		return fmt.Errorf("get locations: %w", err)
	}

	cfg.NextLocationAreaUrl = locations.Next
	cfg.PreviousLocationAreaUrl = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
