package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration

	cacheMu sync.RWMutex
	cache   map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		interval: interval,
		cache:    make(map[string]cacheEntry),
	}
	// Start the read loop to clean the cache regularly.
	c.readLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)

	go func() {
		for t := range ticker.C {
			c.cacheMu.Lock()
			for key, entry := range c.cache {
				if t.Sub(entry.createdAt) > c.interval {
					delete(c.cache, key)
				}
			}
			c.cacheMu.Unlock()
		}
	}()
}
