package pokecache

import (
	"sync"
)

type PokemonDet struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Pokedex struct {
	cache map[string]PokemonDet
	mux   *sync.Mutex
}

func NewPokedex() Pokedex {
	dex := Pokedex{
		cache: make(map[string]PokemonDet),
		mux:   &sync.Mutex{},
	}
	return dex
}

func (p *Pokedex) Add(name string, pokedet PokemonDet) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.cache[name] = pokedet
	return nil
}

func (p *Pokedex) Get(name string) (PokemonDet, bool) {
	data, ok := p.cache[name]
	if !ok {
		return PokemonDet{}, false
	}

	return data, true
}
