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
	Client   pokecache.Cache
	Pokedex  pokecache.Pokedex
	MapUrl   urlCfg
	PokeUrl  string
	CatchUrl string
}

func NewClient(interval time.Duration) Client {
	return Client{
		Client:  pokecache.NewCache(interval),
		Pokedex: pokecache.NewPokedex(),
		MapUrl: urlCfg{
			Size:   20,
			Offset: -20,
			Url:    "https://pokeapi.co/api/v2/location/",
		},
		PokeUrl:  "https://pokeapi.co/api/v2/location-area/",
		CatchUrl: "https://pokeapi.co/api/v2/pokemon/",
	}
}
