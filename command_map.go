package main

import (
	"fmt"

	"github.com/hulkbusterks/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	locationsResp, err := pokeapi.ListLocations(nil)
	if err != nil {
		return nil
	}
	//nextURL := locationsResp.Next
	//previousURL := locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
