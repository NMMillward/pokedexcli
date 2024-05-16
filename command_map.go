package main

import (
	"encoding/json"
	"fmt"
	"github.com/nmmillward/pokedexcli/internal/pokecache"
	"io"
	"log"
	"net/http"
)

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandMap(cfg *Config, cache *pokecache.Cache) error {
	count := 20
	for i := 0; i < count; i++ {
		name := getLocationAreaById(cfg.next+i, cache)
		fmt.Println(name)
	}
	cfg.next += count
	cfg.previous += count
	return nil
}

func commandMapBack(cfg *Config, cache *pokecache.Cache) error {
	count := 20
	start := cfg.previous - (count * 2)
	for i := 0; i < count; i++ {
		name := getLocationAreaById(start+i, cache)
		fmt.Println(name)
	}
	cfg.next -= count
	cfg.previous -= count
	return nil
}

func commandExplore(cfg *Config, cache *pokecache.Cache) error {
	if len(cfg.params) == 0 {
		print("Please provide at least one location area")
		return nil
	}
	fmt.Println(fmt.Sprintf("Exploring %s...", cfg.params[0]))
	if area, ok := getLocationArea(cfg.params[0], cache); ok {
		fmt.Println("Found Pokemon:")
		for _, encounter := range area.PokemonEncounters {
			fmt.Println(" - ", encounter.Pokemon.Name)
		}
	}
	return nil
}

func getLocationAreaById(id int, cache *pokecache.Cache) string {
	return getLocationAreaName(fmt.Sprint(id), cache)
}

func getLocationAreaName(name string, cache *pokecache.Cache) string {
	area, _ := getLocationArea(name, cache)
	return area.Name
}

func getLocationArea(name string, cache *pokecache.Cache) (LocationArea, bool) {
	area := LocationArea{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", name)
	if val, ok := cache.Get(url); ok {
		//fmt.Println("Used cached value")
		if err := json.Unmarshal(val, &area); err != nil {
			log.Fatal(err)
		} else {
			return area, true
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	cache.Add(url, body)
	if err := json.Unmarshal(body, &area); err != nil {
		log.Fatal(err)
	}
	return area, true
}
