package main

import (
	"time"

	"github.com/jbdoumenjou/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPIClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
}

func main() {
	config := &config{
		pokeAPIClient: pokeapi.NewClient(2 * time.Minute),
	}
	startRepl(config)
}
