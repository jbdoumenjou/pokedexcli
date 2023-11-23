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
