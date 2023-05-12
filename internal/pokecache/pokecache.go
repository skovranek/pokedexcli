package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	Mux		*sync.Mutex
}

type cacheEntry struct {
	created time.Time
	val     []byte
}

func NewCache() Cache {
	newCache := Cache{
		entries: map[string]cacheEntry{},
		Mux:     &sync.Mutex{},
	}
	go newCache.reapLoop()
	return newCache
}

type cache interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
	reapLoop(interval time.Duration)
}

func (cache Cache) Add(key string, val []byte) {
	cache.Mux.Lock()
	defer cache.Mux.Unlock()
	cache.entries[key] = cacheEntry{
		created: time.Now(),
		val:	 val,
	}
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.Mux.Lock()
	defer cache.Mux.Unlock()
	entry, ok := cache.entries[key]
	return entry.val, ok
}

func (cache Cache) reapLoop() {
	ch := time.Tick(3 * time.Minute)
	previousInterval := time.Now()
	for currentIntervalNow := range ch {
		cache.Mux.Lock()
		for key, val := range cache.entries {
			if val.created.Before(previousInterval) {
				delete(cache.entries, key)
			}
		}
		cache.Mux.Unlock()
		previousInterval = currentIntervalNow
	}
}