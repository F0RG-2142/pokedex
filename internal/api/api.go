package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/F0RG-2142/pokedex/internal/pokecache"
)

type Pokemon struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	BaseXp int    `json:"base_experience"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocArea struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	GameIndex      int         `json:"game_index"`
	PokeEncounters []Encounter `json:"pokemon_encounters"` // Match JSON key
}

type Page struct {
	Start int
	End   int
}

type Config struct {
	CurrentPage int
	Cache       *pokecache.Cache
	Pokedex     map[string]Pokemon
}

var pages = map[int]Page{
	1: {Start: 1, End: 20},
	2: {Start: 21, End: 40},
	3: {Start: 41, End: 60},
	4: {Start: 61, End: 80},
}

func MapCommand(c *Config, _ string) error {
	// Move to next page
	c.CurrentPage++

	// Don't go beyond the defined pages
	if _, exists := pages[c.CurrentPage]; !exists {
		// If page doesn't exist, create it dynamically
		lastPage := c.CurrentPage - 1
		lastEnd := pages[lastPage].End
		pages[c.CurrentPage] = Page{
			Start: lastEnd + 1,
			End:   lastEnd + 20,
		}
	}

	CurrentPage := pages[c.CurrentPage]
	start := CurrentPage.Start
	end := CurrentPage.End

	fmt.Printf("Displaying locations %d to %d (Page %d)\n", start, end, c.CurrentPage)

	cache := pokecache.NewCache(2 * time.Minute)

	//Check if page is stored in cache
	for i := start; i <= end; i++ {
		apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", strconv.Itoa(i))

		// Try to get from cache
		cachedData, exists := cache.Get(apiUrl)

		if exists {
			// Use cached data
			fmt.Println(string(cachedData))
		} else {
			// If not stored in cache, get from API
			res, err := http.Get(apiUrl)
			if err != nil {
				fmt.Printf("Error fetching data: %v (Continuing...)\n", err)
				continue
			}

			defer res.Body.Close()

			var locArea struct {
				Name string `json:"name"`
			}

			decoder := json.NewDecoder(res.Body)
			if err := decoder.Decode(&locArea); err != nil {
				fmt.Printf("Error decoding JSON: %v (Continuing...)\n", err)
				continue
			}

			// Print and add to cache
			fmt.Println(locArea.Name)
			cache.Add(apiUrl, []byte(locArea.Name))
		}
	}
	return nil
}

func MapBackCommand(c *Config, _ string) error {
	// Move to previous page
	c.CurrentPage--

	// Don't go below page 1
	if c.CurrentPage < 1 {
		c.CurrentPage = 1
		fmt.Println("At the first page already!")
	}

	CurrentPage := pages[c.CurrentPage]
	start := CurrentPage.Start
	end := CurrentPage.End

	fmt.Printf("Displaying locations %d to %d (Page %d)\n", start, end, c.CurrentPage)

	cache := pokecache.NewCache(2 * time.Minute)

	for i := start; i <= end; i++ {
		apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", strconv.Itoa(i))

		// Try to get from cache
		cachedData, exists := cache.Get(apiUrl)

		if exists {
			// Use cached data
			fmt.Println(string(cachedData))
		} else {
			// If not stored in cache, get from API
			res, err := http.Get(apiUrl)
			if err != nil {
				fmt.Printf("Error fetching data: %v (Continuing...)\n", err)
				continue
			}

			defer res.Body.Close()

			var locArea struct {
				Name string `json:"name"`
			}

			decoder := json.NewDecoder(res.Body)
			if err := decoder.Decode(&locArea); err != nil {
				fmt.Printf("Error decoding JSON: %v (Continuing...)\n", err)
				continue
			}

			// Print and add to cache
			fmt.Println(locArea.Name)
			cache.Add(apiUrl, []byte(locArea.Name))
		}
	}

	return nil
}

func ExploreCommand(c *Config, location string) error {
	encounterCache := pokecache.NewCache(2 * time.Minute)
	apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)

	// Check if data is in cache
	cachedData, exists := encounterCache.Get(apiUrl)
	if exists {
		var locationArea LocArea
		if err := json.Unmarshal(cachedData, &locationArea); err != nil {
			return fmt.Errorf("error parsing cached data: %v", err)
		}
		for _, encounter := range locationArea.PokeEncounters {
			fmt.Printf("- %s\n", encounter.Pokemon.Name)
		}
		return nil
	}

	// Fetch from API
	res, err := http.Get(apiUrl)
	if err != nil {
		return fmt.Errorf("error exploring %s: %v", location, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}
	// fmt.Println("Raw JSON:", string(body))
	//Unmarshal JSON
	var locationArea LocArea
	if err = json.Unmarshal(body, &locationArea); err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}
	// fmt.Printf("Parsed: %+v\n", locationArea)
	//Add to cache
	encounterCache.Add(apiUrl, body)

	for _, encounter := range locationArea.PokeEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func CatchCommand(c *Config, pokemon string) error {
	//Fetch from api
	fmt.Printf("Throwing a pokeball at %s\n", pokemon)
	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon))
	if err != nil {
		return fmt.Errorf("couldnt fetch pokemon data: %v", err)
	}
	//Read and decode response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	defer res.Body.Close()

	var fetchedPokemon Pokemon
	if err := json.Unmarshal(body, &fetchedPokemon); err != nil {
		return fmt.Errorf("error unmarshaling JSON")
	}
	//Chance-to-catch logic
	catchChance := max(1.0-(float64(fetchedPokemon.BaseXp)/650.0), 0.1) //Chance floor is 10%
	time.Sleep(1 * time.Second)
	//Catch and print error or pokedex
	if rand.Float64() < catchChance {

		fmt.Printf("Gotcha! %s was caught!\n", pokemon)
		c.Pokedex[pokemon] = fetchedPokemon
		fmt.Println("Pokedex:")
		for _, v := range c.Pokedex {
			fmt.Println(v.Name)
		}
	} else {
		fmt.Printf("%s broke free!\n", pokemon)
	}

	return nil
}

func InspectCommand(c *Config, pokemon string) error {
	return nil
}
