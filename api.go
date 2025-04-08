package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type locArea struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`
}

type page struct {
	Start int
	End   int
}

type config struct {
	currentPage int
}

var pages = map[int]page{
	1: {Start: 1, End: 20},
	2: {Start: 21, End: 40},
	3: {Start: 41, End: 60},
	4: {Start: 61, End: 80},
}

func mapCommand(c *config) error {
	// Move to next page
	c.currentPage++

	// Don't go beyond the defined pages
	if _, exists := pages[c.currentPage]; !exists {
		// If page doesn't exist, create it dynamically
		lastPage := c.currentPage - 1
		lastEnd := pages[lastPage].End
		pages[c.currentPage] = page{
			Start: lastEnd + 1,
			End:   lastEnd + 20,
		}
	}

	currentPage := pages[c.currentPage]
	start := currentPage.Start
	end := currentPage.End

	fmt.Printf("Displaying locations %d to %d (Page %d)\n", start, end, c.currentPage)

	var locArea locArea
	for i := start; i <= end; i++ {
		res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", strconv.Itoa(i)))
		if err != nil {
			fmt.Printf("Error fetching data: %v (Continuing...)\n", err)
			continue
		}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&locArea); err != nil {
			res.Body.Close()
			return err
		}
		fmt.Println(locArea.Name)
		res.Body.Close()
	}

	return nil
}

func mapBackCommand(c *config) error {
	// Move to previous page
	c.currentPage--

	// Don't go below page 1
	if c.currentPage < 1 {
		c.currentPage = 1
		fmt.Println("At the first page already!")
	}

	currentPage := pages[c.currentPage]
	start := currentPage.Start
	end := currentPage.End

	fmt.Printf("Displaying locations %d to %d (Page %d)\n", start, end, c.currentPage)

	var locArea locArea
	for i := start; i <= end; i++ {
		res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", strconv.Itoa(i)))
		if err != nil {
			fmt.Printf("Error fetching data: %v (Continuing...)\n", err)
			continue
		}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&locArea); err != nil {
			res.Body.Close()
			return err
		}
		fmt.Println(locArea.Name)
		res.Body.Close()
	}

	return nil
}
