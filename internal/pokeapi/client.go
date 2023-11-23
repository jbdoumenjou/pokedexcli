package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jbdoumenjou/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

// Client is the HTTP client used to make requests to the PokeAPI.
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient returns a new Client.
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

// get makes a GET request to the PokeAPI with cache.
func (c *Client) get(url string) ([]byte, error) {
	// First check the cache
	if cached, ok := c.cache.Get(url); ok {
		fmt.Println("Using cache")
		return cached, nil
	}

	// no cache, let's make the request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request get on %q: %w", url, err)
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("resp failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	// Populate the cache
	c.cache.Add(url, body)

	return body, nil
}
