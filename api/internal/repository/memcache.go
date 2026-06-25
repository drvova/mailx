package repository

import (
	"context"
	"strconv"
	"sync"
	"time"
)

type cacheEntry struct {
	value     string
	expiresAt time.Time
}

// MemCache is an in-memory key-value cache that implements service.Cache.
type MemCache struct {
	mu    sync.RWMutex
	items map[string]cacheEntry
}

func NewMemCache() *MemCache {
	c := &MemCache{items: make(map[string]cacheEntry)}
	go c.evict()
	return c
}

func (c *MemCache) Set(_ context.Context, key string, value any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = cacheEntry{
		value:     toString(value),
		expiresAt: time.Now().Add(expiration),
	}
	return nil
}

func (c *MemCache) Get(_ context.Context, key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.items[key]
	if !ok || time.Now().After(e.expiresAt) {
		return "", nil
	}
	return e.value, nil
}

func (c *MemCache) Del(_ context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
	return nil
}

func (c *MemCache) Incr(_ context.Context, key string, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[key]
	var n int64
	if ok && time.Now().Before(e.expiresAt) {
		n, _ = strconv.ParseInt(e.value, 10, 64)
	}
	n++
	c.items[key] = cacheEntry{
		value:     strconv.FormatInt(n, 10),
		expiresAt: time.Now().Add(expiration),
	}
	return nil
}

func (c *MemCache) evict() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, e := range c.items {
			if now.After(e.expiresAt) {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}

func toString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return ""
	}
}
