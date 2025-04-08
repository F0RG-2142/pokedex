package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var c = config{
	currentPage: 0,
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
		callback:    mapCommand,
	},
	"mapb": {
		name:        "mapback",
		description: "get previous 20 locations",
		callback:    mapBackCommand,
	},
}

func cleanInput(text string) []string {
	command := strings.ToLower(strings.Split(text, " ")[0])
	return []string{command}
}

func exitCommand(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(c *config) error {
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
