package main

import (
	"time"

	"github.com/hulkbusterks/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 10*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
