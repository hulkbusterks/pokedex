package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid pokemon name!")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if res > 40 {
		fmt.Printf("%s escaped \n", name)
		return nil
	}

	fmt.Printf("%s was caught! \n", name)
	fmt.Println("For more information about the inspect ", name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
