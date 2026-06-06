package memory

import (
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

type BigCacheWrapper[K ~string, V any] struct {
	cache *bigcache.BigCache
}

type genericCacheEntry[V any] struct {
	Value     V         `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewBigCache[K ~string, V any]() (*BigCacheWrapper[K, V], error) {
	c, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		return nil, err
	}
	return &BigCacheWrapper[K, V]{cache: c}, nil
}

func (b *BigCacheWrapper[K, V]) Set(key K, value V, ttl time.Duration) error {
	entry := genericCacheEntry[V]{Value: value}
	if ttl > 0 {
		entry.ExpiresAt = time.Now().Add(ttl)
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return b.cache.Set(string(key), data)
}

func (b *BigCacheWrapper[K, V]) Get(key K) (V, bool) {
	var v V
	data, err := b.cache.Get(string(key))
	if err != nil {
		return v, false
	}
	var entry genericCacheEntry[V]
	if err := json.Unmarshal(data, &entry); err != nil {
		return v, false
	}
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		_ = b.cache.Delete(string(key))
		return v, false
	}
	return entry.Value, true
}

func (b *BigCacheWrapper[K, V]) Delete(key K) error {
	return b.cache.Delete(string(key))
}

func (b *BigCacheWrapper[K, V]) Clear() error {
	return b.cache.Reset()
}

func (b *BigCacheWrapper[K, V]) Close() error {
	return nil
}
