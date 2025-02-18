package main

import (
	"fmt"
)

var nextURL *string
var prevURL *string

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(nextURL)
	if err != nil {
		return nil
	}
	cfg.nextLocationURL = locationsResp.Next
	cfg.preLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if prevURL == nil {
		fmt.Println("Already on first page!")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(prevURL)
	if err != nil {
		return nil
	}
	cfg.nextLocationURL = locationsResp.Next
	cfg.preLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
