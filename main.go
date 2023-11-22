package main

import (
	"github.com/jbdoumenjou/pokedexcli/pokeapi"
)

type config struct {
	pokeAPIClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
}

func main() {
	config := &config{
		pokeAPIClient: pokeapi.NewClient(),
	}
	startRepl(config)
}
