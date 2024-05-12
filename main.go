package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	next     int
	previous int
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
			name:        "map",
			description: "Display the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

func main() {
	commands := getCommands()
	cfg := config{
		next:     1,
		previous: 1,
	}

	for {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if command, ok := commands[input.Text()]; ok {
			if err := command.callback(&cfg); err != nil {
				break
			}
		} else {
			_ = commandHelp(&cfg)
		}
	}
}
