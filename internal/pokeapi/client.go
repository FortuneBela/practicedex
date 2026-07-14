package pokeapi

import (
	"net/http"
	"time"

	"github.com/FortuneBela/practicedex/internal/pokecache"
)

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}
