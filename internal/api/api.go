package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/F0RG-2142/pokedex/internal/pokecache"
)

type LocArea struct {
	ID        int    `json:"id"`
	Name      []byte `json:"name"`
	GameIndex int    `json:"game_index"`
}

type Page struct {
	Start int
	End   int
}

type Config struct {
	CurrentPage int
}

var pages = map[int]Page{
	1: {Start: 1, End: 20},
	2: {Start: 21, End: 40},
	3: {Start: 41, End: 60},
	4: {Start: 61, End: 80},
}

func MapCommand(c *Config) error {
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

	cache := pokecache.NewCache(5 * time.Minute)

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

func MapBackCommand(c *Config) error {
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

	cache := pokecache.NewCache(5 * time.Second)

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
