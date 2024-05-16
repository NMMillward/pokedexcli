package main

import (
	"fmt"
	"github.com/nmmillward/pokedexcli/internal/pokecache"
)

func commandPokedex(cfg *Config, cache *pokecache.Cache) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("No pokemon have been caught")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range cfg.pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}

	return nil
}
