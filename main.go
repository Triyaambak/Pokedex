package main

import (
	"time"

	pokeclient "github.com/Triyaambak/Pokedex/internal/pokeclient"
)

func main() {
	pokecfg := pokeclient.NewClient(time.Minute * 5)
	startRepl(pokecfg)
}
