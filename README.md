## Overview
This project is a Command-Line Interface (CLI) for interacting with the PokeAPI. It allows users to explore different Pokemon areas, catch Pokemon, inspect their details, and manage their Pokedex. The CLI is implemented in Go and supports various commands for a seamless experience.

## Prerequisites
Go (version 1.24+)

Internet connection (for fetching data from PokeAPI)

## Installation
Clone the repository:

bash
```
git clone https://github.com/hulkbusterks/pokedexcli.git
cd pokedexcli
```
## Build the project:
bash
```
go build -o pokedexcli
```
## Run the CLI:
bash
```
./pokedexcli
```
## Usage
Upon running the CLI, you will enter the Pokedex REPL (Read-Eval-Print Loop). You can use the following commands:

## Commands
help: Displays a help message with a list of available commands.

exit: Exits the Pokedex CLI.

map: Gets the next page of locations.

mapb: Gets the previous page of locations.

explore <area-name>: Finds the Pokemons in the specified area.

catch <pokemon-name>: Throws a pokeball at the specified Pokemon.

inspect <pokemon-name>: Gets more information about the specified Pokemon.

pokedex: Displays all the Pokemons you have caught.
