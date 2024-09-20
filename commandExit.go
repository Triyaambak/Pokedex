package main

import (
	"os"

	"github.com/Triyaambak/Pokedex/internal/pokeclient"
)

func commandExit(pokecfg *pokeclient.Client) error {
	os.Exit(0)
	return nil
}
