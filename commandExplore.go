package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Triyaambak/Pokedex/internal/pokeclient"
)

type Pokemon struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(pokecfg *pokeclient.Client) error {
	fmt.Println("Enter the name of the location you want to explore > ")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	city := reader.Text()
	mapurl := fmt.Sprintf("%s%s", pokecfg.MapUrl.Url, city)
	pokeurl := fmt.Sprintf("%s%s-area", pokecfg.PokeUrl, city)

	data, ok := pokecfg.Client.Get(pokeurl)
	if ok {
		err := printPokemon(data)
		return err
	}

	data, err := explore(mapurl, pokeurl)
	if err != nil {
		fmt.Println("Something went wrong while exploring the city , check if the city exists or not")
	}

	printPokemon(data)
	pokecfg.Client.Add(pokeurl, data)
	return nil
}

func explore(mapurl, pokeurl string) ([]byte, error) {

	req, err := http.NewRequest("GET", mapurl, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("no location with such area")
	}

	req, err = http.NewRequest("GET", pokeurl, nil)
	if err != nil {
		return nil, err
	}

	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, err
}

func printPokemon(data []byte) error {

	pokemon := Pokemon{}

	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Pokemons encountered are....")
	for _, encounter := range pokemon.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
