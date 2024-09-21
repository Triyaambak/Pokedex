package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Triyaambak/Pokedex/internal/pokeclient"
)

func commandInspect(pokecfg *pokeclient.Client) error {
	fmt.Println("Enter the name of the pokemon you want to inspect")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	pokemon := reader.Text()

	p, ok := pokecfg.Pokedex.Get(pokemon)
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
