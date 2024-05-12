package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func main() {
	commands := getCommands()
	willExit := false

	for {
		fmt.Print("Pokedex > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		switch input.Text() {
		case "exit":
			willExit = true
			break
		default:
			if command, ok := commands[input.Text()]; ok {
				if err := command.callback(); err != nil {
					break
				}
			} else {
				_ = commandHelp()
			}
		}

		if willExit {
			break
		}
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", command.name, command.description))
	}
	return nil
}

func commandExit() error {
	return nil
}
