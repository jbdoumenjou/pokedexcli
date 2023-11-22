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
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("new request Get on %q: %w", fullURL, err)
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("resp failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	var locations LocationAreasResp
	if err = json.Unmarshal(body, &locations); err != nil {
		return LocationAreasResp{}, fmt.Errorf("unmarshal body: %w", err)
	}

	return locations, nil
}
