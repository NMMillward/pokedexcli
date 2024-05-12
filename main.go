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
	callback    func(cfg *config, cache *pokecache.Cache) error
}

type config struct {
	next     int
	previous int
	params   []string
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
	}
}

func main() {
	commands := getCommands()
	cache := pokecache.NewCache(time.Minute * 5)
	cfg := config{
		next:     1,
		previous: 1,
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
