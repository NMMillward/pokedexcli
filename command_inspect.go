package main

import (
	"fmt"
	"github.com/nmmillward/pokedexcli/internal/pokecache"
)

func commandInspect(cfg *Config, _ *pokecache.Cache) error {
	if len(cfg.params) == 0 {
		fmt.Println("You stare into an empty Pokeball. Are you ok?")
	} else {
		name := cfg.params[0]
		if pokemon, ok := cfg.pokedex[name]; ok {
			fmt.Printf("%s", pokemon.GetInspect())
		} else {
			fmt.Println("You have not caught that pokemon")
		}
	}
	return nil
}
