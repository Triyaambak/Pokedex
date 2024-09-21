package pokecache

import (
	"fmt"
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

func (p *Pokedex) Print() {
	for _, v := range p.cache {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Height: %d\n", v.Height)
		fmt.Printf("Weight: %d\n", v.Weight)
		fmt.Println("Stats:")
		for _, stat := range v.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range v.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
		fmt.Println()
	}
}
