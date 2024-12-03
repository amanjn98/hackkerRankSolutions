package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

//They also asked me to write workable node.js solution by starting the server and setting- getting a json record in local storage and file storage via HTTP request.

type Storage struct {
	LocalStorage map[string]interface{}
	mu           sync.RWMutex // Ensures thread-safe access to LocalStorage
	FilePath     string
}

func NewStorage(filePath string) *Storage {
	return &Storage{
		LocalStorage: make(map[string]interface{}),
		FilePath:     filePath,
	}
}

// Save to file
func (s *Storage) saveToFile() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := json.Marshal(s.LocalStorage)
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath, data, 0644)
}

// Load from file
func (s *Storage) loadFromFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist, ignore
		}
		return err
	}

	return json.Unmarshal(file, &s.LocalStorage)
}

// Set key-value in local storage and optionally save to file
func (s *Storage) Set(key string, value interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LocalStorage[key] = value
	return s.saveToFile()
}

// Get value by key from local storage
func (s *Storage) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.LocalStorage[key]
	return value, exists
}

// HTTP Handlers
func (s *Storage) setHandler(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	for key, value := range request {
		if err := s.Set(key, value); err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data saved successfully")
}

func (s *Storage) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing 'key' query parameter", http.StatusBadRequest)
		return
	}

	value, exists := s.Get(key)
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(map[string]interface{}{
		"key":   key,
		"value": value,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	// Initialize storage with a file for persistence
	filePath := "storage.json"
	storage := NewStorage(filePath)

	// Load data from file on startup
	if err := storage.loadFromFile(); err != nil {
		log.Fatalf("Failed to load storage file: %v", err)
	}

	// Set up HTTP handlers
	http.HandleFunc("/set", storage.setHandler)
	http.HandleFunc("/get", storage.getHandler)

	// Start server
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
