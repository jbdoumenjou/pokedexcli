package main

import (
	"errors"
	"fmt"
)

// commandMap Displays the names of the next 20 location areas in the Pokemon world.
func commandMapb(cfg *config) error {
	if cfg.PreviousLocationAreaUrl == nil {
		return errors.New("you are on the first page")
	}
	locations, err := cfg.pokeAPIClient.ListLocationAreas(cfg.PreviousLocationAreaUrl)
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
