package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

// Client is the HTTP client used to make requests to the PokeAPI.
type Client struct {
	httpClient http.Client
}

// NewClient returns a new Client.
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
