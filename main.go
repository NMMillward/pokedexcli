package main

import (
	"bufio"
	"fmt"
	"github.com/nmmillward/pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, cache *pokecache.Cache) error
}

type Config struct {
	next     int
	previous int
	params   []string
	pokedex  map[string]Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore the area of the give name or id\r\n\r\nExample: explore canalave-city-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon\r\n\r\nExample: catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon\r\n\r\nExample: inspect pikachu",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View a list of your caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func main() {
	commands := getCommands()
	cache := pokecache.NewCache(time.Minute * 5)
	cfg := Config{
		next:     1,
		previous: 1,
		pokedex:  make(map[string]Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		command := input.Text()
		if len(command) == 0 {
			_ = commandHelp(&cfg, cache)
		} else {
			fields := strings.Fields(command)

			if command, ok := commands[fields[0]]; ok {
				cfg.params = fields[1:]
				if err := command.callback(&cfg, cache); err != nil {
					break
				}
			} else {
				_ = commandHelp(&cfg, cache)
			}
		}
	}
}
