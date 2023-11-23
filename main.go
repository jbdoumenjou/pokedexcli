package main

import (
	"time"

	"github.com/jbdoumenjou/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPIClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
	caughtPokemons          map[string]pokeapi.Pokemon
}

func main() {
	config := &config{
		pokeAPIClient:  pokeapi.NewClient(2 * time.Minute),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}
