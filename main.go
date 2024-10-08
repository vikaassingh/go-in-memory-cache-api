package main

import "go-in-memory-cache-api/config"

func main() {
	config.IntializeRouter()
}

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Cache struct to hold the in-memory cache
type Cache struct {
	data map[string]string
	mu   sync.RWMutex // Mutex for safe concurrent access
}

// NewCache creates a new Cache instance
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()         // Acquire read lock
	defer c.mu.RUnlock() // Release read lock
	value, exists := c.data[key]
	return value, exists
}

// Set adds a value to the cache
func (c *Cache) Set(key, value string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	c.data[key] = value
}

// Handler for getting a value from the cache
func (c *Cache) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if value, exists := c.Get(key); exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
	} else {
		http.Error(w, "Key not found", http.StatusNotFound)
	}
}

// Handler for setting a value in the cache
func (c *Cache) setHandler(w http.ResponseWriter, r *http.Request) {
	// var body map[string]string
	// if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	// 	http.Error(w, "Invalid request", http.StatusBadRequest)
	// 	return
	// }
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	c.Set(key, value)
	fmt.Printf("\nc:%p-%v", c, c)
	w.WriteHeader(http.StatusNoContent) // No content response
}

func main() {
	cache := NewCache()

	http.HandleFunc("/get", cache.getHandler)
	http.HandleFunc("/set", cache.setHandler)

	fmt.Println("Starting server on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
*/
