package pokecache

import (
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() Cache {
	return Cache{
		entries: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Reap(time.Hour)
	result := c.entries[key].val

	if result != nil {
		return result, true
	}

	return nil, false
}

func (c *Cache) Reap(t time.Duration) {
	for i := range c.entries {
		if time.Since(c.entries[i].createdAt) > t {
			delete(c.entries, i)
			//fmt.Println("deleted cached entry")
		}
	}
}
