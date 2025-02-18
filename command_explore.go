package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid location name!")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s.........\n", name)
	fmt.Println("Found Pokemons:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("--%s\n", pokemon.Pokemon.Name)
	}
	return nil
}
