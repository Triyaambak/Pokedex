package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *urlCfg) error {
	cfg.Offset += 20
	err := getLocationData(cfg)

	if err != nil {
		fmt.Println("Something went wrong while moving in forward direction")
		return err
	}

	return nil
}

func commandMapback(cfg *urlCfg) error {
	if cfg.Offset == 0 {
		fmt.Println("Cannot move in the backward direction , you are in the starting point")
	}
	cfg.Offset -= 20

	err := getLocationData(cfg)
	if err != nil {
		fmt.Println("Something went wrong while moving in backward direction")
		return err
	}
	return nil
}

func getLocationData(cfg *urlCfg) error {
	url := fmt.Sprintf("%s?offset=%d&size=%d", cfg.Url, cfg.Offset, cfg.Size)
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

	data, err := io.ReadAll(w.Body)
	if err != nil {
		return err
	}

	err = printLocationData(data)
	if err != nil {
		return err
	}
	return nil
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
