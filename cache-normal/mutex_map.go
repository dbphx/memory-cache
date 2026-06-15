package memory

import (
	"sync"
	"time"
)

type mutexMapEntry struct {
	value     []byte
	expiresAt time.Time
}

type mutexMapCache struct {
	mu   sync.RWMutex
	data map[string]mutexMapEntry
}

func NewMutexMapCache() Cache {
	return &mutexMapCache{
		data: make(map[string]mutexMapEntry),
	}
}

func (c *mutexMapCache) Set(key string, value []byte, ttl time.Duration) error {
	entry := mutexMapEntry{
		value: append([]byte(nil), value...),
	}
	if ttl > 0 {
		entry.expiresAt = time.Now().Add(ttl)
	}

	c.mu.Lock()
	c.data[key] = entry
	c.mu.Unlock()
	return nil
}

func (c *mutexMapCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	entry, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	if !entry.expiresAt.IsZero() && time.Now().After(entry.expiresAt) {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return nil, false
	}
	return append([]byte(nil), entry.value...), true
}

func (c *mutexMapCache) Delete(key string) error {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
	return nil
}

func (c *mutexMapCache) Clear() error {
	c.mu.Lock()
	c.data = make(map[string]mutexMapEntry)
	c.mu.Unlock()
	return nil
}

func (c *mutexMapCache) Close() error {
	return nil
}
