package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Struct to unmarshal JSON response
type ApiResponse struct {
	Data string `json:"data"`
}

// Fetch data from a given URL and send the response back through a channel
func fetchData(url string, wg *sync.WaitGroup, ch chan<- ApiResponse) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", url, err)
		return
	}

	var data ApiResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error unmarshalling data from %s: %v\n", url, err)
		return
	}

	ch <- data
}

func main() {
	// List of URLs to fetch data from
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
	}

	// Channel to receive API responses
	ch := make(chan ApiResponse, len(urls))
	var wg sync.WaitGroup

	// Start goroutines for each API request
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg, ch)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(ch)

	// Process received data
	for data := range ch {
		fmt.Printf("Received data: %s\n", data.Data)
	}
}
