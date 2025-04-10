package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/F0RG-2142/pokedex/internal/pokecache"

	"github.com/F0RG-2142/pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*api.Config, string) error
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
	"explore": {
		name:        "explore",
		description: "explore a specified area",
		callback:    api.ExploreCommand,
	},
	"catch": {
		name:        "catch",
		description: "catch a specified pokemon",
		callback:    api.CatchCommand,
	},
	"inspect": {
		name:        "inspect",
		description: "inspect a specified pokemon",
		callback:    api.InspectCommand,
	},
	"pokedex": {
		name:        "pokedex",
		description: "print pokedex",
		callback:    api.PokedexCommand,
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

func exitCommand(c *api.Config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(c *api.Config, _ string) error {
	fmt.Println(`Welcome to the Pokedex!
	
	Usage:
	help: Displays a help message
	exit: Exit the Pokedex
	map/mapb: move forwards and backwards through the map
	explore: explore a specific area in the map
	catch: catch a pokemon in your area
	pokedex: view all the pokemon in your pokedex
	inspect: inspect a pokemon in your pokedex`)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := api.Config{
		CurrentPage: 0,
		Cache:       pokecache.NewCache(2 * time.Minute),
		Pokedex:     make(map[string]api.Pokemon), // Initialize here
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		command := text[0]
		arg := ""
		if len(text) > 1 {
			arg = text[1]
		}

		cmd, exists := commands[command]
		if !exists {
			fmt.Printf("The '%s' command does not exist\n", command)
			continue
		}
		cmd.callback(&c, arg)
	}
}
