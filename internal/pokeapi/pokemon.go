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
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (p Pokemon) ToString() string {
	var str string
	str += fmt.Sprintf("Name: %s\n", p.Name)
	str += fmt.Sprintf("Height: %d\n", p.Height)
	str += fmt.Sprintf("Weight: %d\n", p.Weight)
	str += "Stats: \n"
	for _, stat := range p.Stats {
		str += fmt.Sprintf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	str += "Types: \n"
	for _, t := range p.Types {
		str += fmt.Sprintf("  - %s\n", t.Type.Name)
	}

	return str
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
