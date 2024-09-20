package pokeclient

import (
	"time"

	pokecache "github.com/Triyaambak/Pokedex/internal/pokecache"
)

type urlCfg struct {
	Size   int
	Offset int
	Url    string
}

type Client struct {
	Client pokecache.Cache
	Cfg    urlCfg
}

func NewClient(interval time.Duration) Client {
	return Client{
		Client: pokecache.NewCache(interval),
		Cfg: urlCfg{
			Size:   20,
			Offset: -20,
			Url:    "https://pokeapi.co/api/v2/location/",
		},
	}
}
