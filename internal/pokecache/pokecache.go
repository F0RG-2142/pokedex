package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache struct {
	mu sync.Mutex
	v  map[string]cacheEntry
}

func NewCache(interval time.Duration) *cache {
	//Create New Cache
	var c = cache{
		mu: sync.Mutex{},
		v:  make(map[string]cacheEntry),
	}
	return &c
}

func (cache *cache) Add(key string, value []byte) {
	//Make Cache Entry
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.v[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (cache *cache) Get(key string) ([]byte, bool) {
	//Get entry from cache
	cache.mu.Lock()
	defer cache.mu.Unlock()

	value := cache.v[key]
	
}

func (cache *cache) reapLoop(createdAt time.Time) ([]byte, bool) {
	//Delete cache entries that are older than a set interval
	cache.mu.Lock()
	defer cache.mu.Unlock()

}
