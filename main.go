package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/F0RG-2142/pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*api.Config) error
}

var c = api.Config{
	CurrentPage: 0,
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    exitCommand,
	},
	"help": {
		name:        "help",
		description: "Help with Pokedex",
		callback:    helpCommand,
	},
	"map": {
		name:        "map",
		description: "map location areas",
		callback:    api.MapCommand,
	},
	"mapb": {
		name:        "mapback",
		description: "get previous 20 locations",
		callback:    api.MapBackCommand,
	},
}

func cleanInput(text string) []string {
	cleanStr := []string{}
	for _, v := range strings.Split(text, " ") {
		if v != "" { // Skip empty strings
			cleanStr = append(cleanStr, strings.ToLower(v))
		}
	}
	return cleanStr
}

func exitCommand(c *api.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(c *api.Config) error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		command := text[0]
		commands[command].callback(&c)
	}
}
