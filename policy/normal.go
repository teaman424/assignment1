package store_cache

import (
	"fmt"
	"time"
)

// normal cache use map
type Cache struct {
	items    map[string]*entry
	capacity int
}

type entry struct {
	val       string
	createdAt time.Time
}

// NewCache creates a new non-thread safe cache.
func NewCache(capacity int) *Cache {
	return &Cache{
		items:    make(map[string]*entry, 0),
		capacity: capacity,
	}
}

// Set item to the cache. replacing or new any existing item.
// Accumulate key and value lengths and cannot exceed capacity value
func (c *Cache) Set(key string, value string) (ok bool) {
	c.capacity = c.capacity - len(key) - len(value)
	if c.capacity <= 0 {
		return false
	}

	c.items[key] = &entry{
		val:       value,
		createdAt: time.Now(),
	}
	c.capacity = c.capacity - len(key) - len(value)
	return true
}

// Get an item from the cache.
// Returns the item or zero value, and a bool indicating whether the key was found.
func (c *Cache) Get(key string) (value string, ok bool) {
	item, found := c.items[key]
	if !found {
		return
	}
	fmt.Println(len(c.items))
	c.capacity = len(item.val) + len(key) + c.capacity
	return item.val, true
}
