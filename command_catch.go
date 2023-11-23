package main

import (
	"fmt"
	"math/rand"
)

// https://pokeapi.co/docs/v2#location-areas
// commandCatch try to catch a Pokemon and add it to the Pokedex.
func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon name>")
	}

	pokemon, err := cfg.pokeAPIClient.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("get pokemon: %w", err)
	}

	const threshold = 50
	randomNb := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("%s has %d base experience points\n", pokemon.Name, pokemon.BaseExperience)
	fmt.Printf("You need to get less than %d to catch it\n", threshold)
	fmt.Printf("You got %d\n", randomNb)
	if randomNb > threshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.caughtPokemons[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}
