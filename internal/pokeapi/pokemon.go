package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name                   string `json:"name"`
	Order                  int    `json:"order"`
	Weight                 int    `json:"weight"`
}

// GetPokemon return a Pokemon for the given name.
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var pokemon Pokemon
	endpoint := "/pokemon/" + name
	fullURL := baseURL + endpoint

	data, err := c.get(fullURL)
	if err != nil {
		return pokemon, fmt.Errorf("get: %w", err)
	}

	if err = json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, fmt.Errorf("unmarshal body: %w", err)
	}

	return pokemon, nil
}
