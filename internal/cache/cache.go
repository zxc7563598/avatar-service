package cache

import (
	"sync"
	"time"
)

type MemoryCache struct {
	mu    sync.RWMutex
	store map[string]*cacheItem
	ttl   time.Duration
}

type cacheItem struct {
	data []byte
	time time.Time
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		store: make(map[string]*cacheItem),
		ttl:   10 * time.Minute,
	}
}

func (c *MemoryCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.store[key]
	if !ok || time.Since(item.time) > c.ttl {
		return nil, false
	}
	return item.data, true
}

func (c *MemoryCache) Set(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = &cacheItem{data: data, time: time.Now()}
}
