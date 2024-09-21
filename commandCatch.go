package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/Triyaambak/Pokedex/internal/pokecache"
	"github.com/Triyaambak/Pokedex/internal/pokeclient"
)

type PokemonExp struct {
	Exp int `json:"base_experience"`
}

func commandCatch(pokecfg *pokeclient.Client) error {
	fmt.Println("Enter the name of the pokemon you want to catch")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	pokemon := reader.Text()

	pokeUrl := fmt.Sprintf("%s%s", pokecfg.CatchUrl, pokemon)
	data, ok := getPokemon(pokeUrl)
	if !ok {
		fmt.Println("Something went wrong in getting the pokemon , check if the pokemon exists")
	}

	ok, err := simCatch(pokecfg, pokemon, data)
	if err != nil {
		fmt.Println("Something went wrong while catching the pokemon , try again")
	}

	if ok {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
	}
	return nil
}

func getPokemon(url string) ([]byte, bool) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, false
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, false
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, false
	}

	return data, true
}

func simCatch(pokecfg *pokeclient.Client, pokemon string, data []byte) (bool, error) {
	basexp := PokemonExp{}
	pokedet := pokecache.PokemonDet{}

	err := json.Unmarshal(data, &basexp)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &pokedet)
	if err != nil {
		return false, err
	}

	if _, ok := pokecfg.Pokedex.Get(pokemon); !ok {
		pokecfg.Pokedex.Add(pokemon, pokedet)
	}

	fmt.Printf("Throwing a ball at %s....", pokemon)

	catchRate := 1.0 / (float64(basexp.Exp)/100 + 1)

	catchAttempt := rand.Float64()

	if catchAttempt < catchRate {
		return true, nil
	} else {
		return false, nil
	}
}
