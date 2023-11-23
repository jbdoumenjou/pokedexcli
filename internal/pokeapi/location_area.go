package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreasResp contains the list of areas location and pagination metadata.
type LocationAreasResp struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// ListLocationAreas returns the list of locations with pagination metadata.
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	var locations LocationAreasResp
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// First check the cache
	if cached, ok := c.cache.Get(fullURL); ok {
		fmt.Println("Using cache")
		if err := json.Unmarshal(cached, &locations); err != nil {
			return locations, fmt.Errorf("unmarshal cached: %w", err)
		}
		return locations, nil
	}

	// no cache, let's make the request
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locations, fmt.Errorf("new request Get on %q: %w", fullURL, err)
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return locations, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return locations, fmt.Errorf("resp failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	// Populate the cache
	c.cache.Add(fullURL, body)

	if err = json.Unmarshal(body, &locations); err != nil {
		return locations, fmt.Errorf("unmarshal body: %w", err)
	}

	return locations, nil
}

type LocationArea struct {
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// ExploreLocationArea return a location for the given name.
func (c *Client) ExploreLocationArea(name string) (LocationArea, error) {
	var location LocationArea
	endpoint := "/location-area" + "/" + name
	fullURL := baseURL + endpoint

	// First check the cache
	if cached, ok := c.cache.Get(fullURL); ok {
		fmt.Println("Using cache")
		if err := json.Unmarshal(cached, &location); err != nil {
			return location, fmt.Errorf("unmarshal cached: %w", err)
		}
		return location, nil
	}

	// no cache, let's make the request
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return location, fmt.Errorf("new request Get on %q: %w", fullURL, err)
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return location, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return location, fmt.Errorf("resp failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	// Populate the cache
	c.cache.Add(fullURL, body)

	if err = json.Unmarshal(body, &location); err != nil {
		return location, fmt.Errorf("unmarshal body: %w", err)
	}

	return location, nil
}
