package main

import "github.com/Triyaambak/Pokedex/internal/pokeclient"

func commandPokedex(pokecfg *pokeclient.Client) error {
	pokecfg.Pokedex.Print()
	return nil
}
