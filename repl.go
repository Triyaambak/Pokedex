package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeclient "github.com/Triyaambak/Pokedex/internal/pokeclient"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeclient.Client) error
}

func startRepl(pokecfg pokeclient.Client) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(&pokecfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}

func cleanInput(inp string) []string {
	output := strings.ToLower(inp)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Move forward on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Move backwards on the map",
			callback:    commandMapback,
		},
		"explore": {
			name:        "explore",
			description: "Explore all the pokemon in the given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch the pokemon with the given name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the given pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your pokedex",
			callback:    commandPokedex,
		},
	}
}
