package main

import (
	"github.com/nmmillward/pokedexcli/internal/pokecache"
	"os"
)

func commandExit(_ *Config, _ *pokecache.Cache) error {
	os.Exit(0)
	return nil
}
