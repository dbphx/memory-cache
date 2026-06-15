package memory

import (
	"sync"
	"time"
)

type mutexMapEntry[V any] struct {
	value     V
	expiresAt time.Time
}

type mutexMapCache[K ~string, V any] struct {
	mu   sync.RWMutex
	data map[string]mutexMapEntry[V]
}

func NewMutexMapCache[K ~string, V any]() Cache[K, V] {
	return &mutexMapCache[K, V]{
		data: make(map[string]mutexMapEntry[V]),
	}
}

func (c *mutexMapCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	entry := mutexMapEntry[V]{
		value: value,
	}
	if ttl > 0 {
		entry.expiresAt = time.Now().Add(ttl)
	}

	c.mu.Lock()
	c.data[string(key)] = entry
	c.mu.Unlock()
	return nil
}

func (c *mutexMapCache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	entry, ok := c.data[string(key)]
	c.mu.RUnlock()
	if !ok {
		var zero V
		return zero, false
	}
	if !entry.expiresAt.IsZero() && time.Now().After(entry.expiresAt) {
		c.mu.Lock()
		delete(c.data, string(key))
		c.mu.Unlock()
		var zero V
		return zero, false
	}
	return entry.value, true
}

func (c *mutexMapCache[K, V]) Delete(key K) error {
	c.mu.Lock()
	delete(c.data, string(key))
	c.mu.Unlock()
	return nil
}

func (c *mutexMapCache[K, V]) Clear() error {
	c.mu.Lock()
	c.data = make(map[string]mutexMapEntry[V])
	c.mu.Unlock()
	return nil
}

func (c *mutexMapCache[K, V]) Close() error {
	return nil
}
