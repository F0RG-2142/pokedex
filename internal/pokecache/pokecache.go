package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

type Cache struct {
	mu sync.Mutex
	V  map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	//Create New Cache
	var c = Cache{
		mu: sync.Mutex{},
		V:  make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return &c
}

func (cache *Cache) Add(key string, value []byte) {
	//Make Cache Entry
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.V[key] = cacheEntry{
		createdAt: time.Now(),
		Val:       value,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	//`Get entry from cache`
	cache.mu.Lock()
	defer cache.mu.Unlock()
	value, ok := cache.V[key]
	if !ok {
		return nil, false
	}
	return value.Val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// Remove expired entries directly in the loop
		cache.mu.Lock()
		now := time.Now()
		for key, entry := range cache.V {
			if now.Sub(entry.createdAt) > interval {
				delete(cache.V, key)
			}
		}
		cache.mu.Unlock()
	}
}
