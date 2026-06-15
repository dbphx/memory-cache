package memory

import (
	"sync"
	"time"
)

type syncMapEntry struct {
	value     []byte
	expiresAt time.Time
}

type syncMapCache struct {
	data sync.Map
}

func NewSyncMapCache() Cache {
	return &syncMapCache{}
}

func (c *syncMapCache) Set(key string, value []byte, ttl time.Duration) error {
	entry := syncMapEntry{
		value: append([]byte(nil), value...),
	}
	if ttl > 0 {
		entry.expiresAt = time.Now().Add(ttl)
	}
	c.data.Store(key, entry)
	return nil
}

func (c *syncMapCache) Get(key string) ([]byte, bool) {
	raw, ok := c.data.Load(key)
	if !ok {
		return nil, false
	}
	entry := raw.(syncMapEntry)
	if !entry.expiresAt.IsZero() && time.Now().After(entry.expiresAt) {
		c.data.Delete(key)
		return nil, false
	}
	return append([]byte(nil), entry.value...), true
}

func (c *syncMapCache) Delete(key string) error {
	c.data.Delete(key)
	return nil
}

func (c *syncMapCache) Clear() error {
	c.data.Range(func(key, _ any) bool {
		c.data.Delete(key)
		return true
	})
	return nil
}

func (c *syncMapCache) Close() error {
	return nil
}
