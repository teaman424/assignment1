package store_cache

import (
	"time"
)

type FifoCache struct {
	Cache // the normal cache is embedded
	qeueu []string
}

// NewFifoCache creates a new non-thread safe cache.
func NewFifoCache(capacity int) *FifoCache {
	return &FifoCache{
		Cache: Cache{
			items:    make(map[string]*entry, 0),
			capacity: capacity,
		},
		qeueu: make([]string, 0),
	}
}

// Set item to the cache. replacing or new any existing item and item's key add qeueu
// Accumulate key and value lengths and cannot exceed capacity value
func (c *FifoCache) Set(key string, value string) (ok bool) {
	c.capacity = c.capacity - len(key) - len(value)
	if c.capacity <= 0 {
		return false
	}

	c.items[key] = &entry{
		val:       value,
		createdAt: time.Now(),
	}
	c.capacity = c.capacity - (len(key) + len(value))
	c.qeueu = append(c.qeueu, key)
	return true
}

//=============================
// Get an item from the cache,not need remove from qeueu.
// So use normal cache Get func
//=============================

// First In Last Out,pop frist item
func (c *FifoCache) Pop() (value string, ok bool) {
	if len(c.qeueu) <= 0 {
		return
	}

	item_key := c.qeueu[0]
	c.qeueu = c.qeueu[1:]

	item, found := c.items[item_key]
	if !found {
		return
	}
	delete(c.items, item_key)
	c.capacity = len(item.val) + len(item_key) + c.capacity
	return item.val, true
}
