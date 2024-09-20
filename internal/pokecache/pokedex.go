package pokecache

import "sync"

type Pokemon struct {
	name string
}

type Pokdex struct {
	cache map[string]Pokemon
	mux   *sync.Mutex
}
