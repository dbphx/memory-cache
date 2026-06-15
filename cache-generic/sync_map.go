package memory

import (
	"sync"
	"time"
)

type syncMapEntry[K ~string, V any] struct {
	value     V
	expiresAt time.Time
}

type syncMapCache[K ~string, V any] struct {
	data sync.Map
}

func NewSyncMapCache[K ~string, V any]() Cache[K, V] {
	return &syncMapCache[K, V]{}
}

func (c *syncMapCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	entry := syncMapEntry[K, V]{
		value: value,
	}
	if ttl > 0 {
		entry.expiresAt = time.Now().Add(ttl)
	}
	c.data.Store(string(key), entry)
	return nil
}

func (c *syncMapCache[K, V]) Get(key K) (V, bool) {
	raw, ok := c.data.Load(string(key))
	if !ok {
		var zero V
		return zero, false
	}
	entry := raw.(syncMapEntry[K, V])
	if !entry.expiresAt.IsZero() && time.Now().After(entry.expiresAt) {
		c.data.Delete(string(key))
		var zero V
		return zero, false
	}
	return entry.value, true
}

func (c *syncMapCache[K, V]) Delete(key K) error {
	c.data.Delete(string(key))
	return nil
}

func (c *syncMapCache[K, V]) Clear() error {
	c.data.Range(func(key, _ any) bool {
		c.data.Delete(key)
		return true
	})
	return nil
}

func (c *syncMapCache[K, V]) Close() error {
	return nil
}
