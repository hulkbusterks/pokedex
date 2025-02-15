package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	s := make([]string, 0)
	var input string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = scanner.Text()
		s = cleanInput(input)
		if len(s) > 0 {
			fmt.Printf("Your command was: %s \n", s[0])
		}
	}
}

func cleanInput(text string) []string {
	cleanslice := make([]string, 0)
	text = strings.ToLower(text)
	cleanslice = strings.Fields(text)
	return cleanslice
}
