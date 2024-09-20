package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokeclient "github.com/Triyaambak/Pokedex/internal/pokeclient"
)

func commandMap(pokecfg *pokeclient.Client) error {
	pokecfg.Cfg.Offset += 20
	err := getLocationData(pokecfg)

	if err != nil {
		fmt.Println("Something went wrong while moving in forward direction")
		return err
	}

	return nil
}

func commandMapback(pokecfg *pokeclient.Client) error {
	if pokecfg.Cfg.Offset == 0 {
		fmt.Println("Cannot move in the backward direction , you are in the starting point")
	}
	pokecfg.Cfg.Offset -= 20

	err := getLocationData(pokecfg)
	if err != nil {
		fmt.Println("Something went wrong while moving in backward direction")
		return err
	}

	return nil
}

func getLocationData(pokecfg *pokeclient.Client) error {
	url := fmt.Sprintf("%s?offset=%d&size=%d", pokecfg.Cfg.Url, pokecfg.Cfg.Offset, pokecfg.Cfg.Size)

	data, ok := pokecfg.Client.Get(url)

	if ok {
		err := printLocationData(data)
		if err == nil {
			return nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	w, err := client.Do(req)
	if err != nil {
		return err
	}
	defer w.Body.Close()

	data, err = io.ReadAll(w.Body)
	if err != nil {
		return err
	}

	pokecfg.Client.Add(url, data)

	err = printLocationData(data)
	return err
}

func printLocationData(data []byte) error {
	type Location struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
	}

	locations := Location{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location)
	}

	return nil
}
