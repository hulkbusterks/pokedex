package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.caughtPokemon {
		fmt.Println("- ", p.Name)
	}
	return nil
}
