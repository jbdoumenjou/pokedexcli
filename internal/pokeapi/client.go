package pokeapi

import (
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
