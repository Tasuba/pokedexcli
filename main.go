package main

import (
	"time"

	"github.com/Tasuba/pokedexcli/internal/pokeapi"
	"github.com/Tasuba/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokeCache,
	}

	startRepl(cfg)
}
