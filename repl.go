package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type urlCfg struct {
	Size   int
	Offset int
	Url    string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*urlCfg) error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	cfg := urlCfg{
		Size:   20,
		Offset: -20,
		Url:    "https://pokeapi.co/api/v2/location/",
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(&cfg)
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
	}
}
