package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mutex    *sync.RWMutex
	duration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.entries[key] = cacheEntry{time.Now(), val}
	//fmt.Println("Added key:", key)
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	if entry, ok := cache.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (cache *Cache) Delete(key string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	delete(cache.entries, key)
	//fmt.Println("Deleted key:", key)
}

func (cache *Cache) Reap() {
	for key := range cache.entries {
		life := time.Now().Sub(cache.entries[key].createdAt)
		if life >= cache.duration {
			cache.Delete(key)
		}
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		cache.Reap()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
		duration: interval,
		mutex:    &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}
