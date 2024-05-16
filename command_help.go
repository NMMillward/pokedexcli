package main

import (
	"fmt"
	"github.com/nmmillward/pokedexcli/internal/pokecache"
)

func commandHelp(_ *Config, _ *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", command.name, command.description))
	}
	return nil
}
