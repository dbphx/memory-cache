package memory

import (
	"context"
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

type bytesCacheEntry struct {
	Value     []byte    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}

type BigCacheWrapper struct {
	cache *bigcache.BigCache
}

func NewBigCache() (*BigCacheWrapper, error) {
	c, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return nil, err
	}
	return &BigCacheWrapper{cache: c}, nil
}

func (b *BigCacheWrapper) Set(key string, value []byte, ttl time.Duration) error {
	entry := bytesCacheEntry{Value: value}
	if ttl > 0 {
		entry.ExpiresAt = time.Now().Add(ttl)
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return b.cache.Set(key, data)
}

func (b *BigCacheWrapper) Get(key string) ([]byte, bool) {
	data, err := b.cache.Get(key)
	if err != nil {
		return nil, false
	}
	var entry bytesCacheEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, false
	}
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		_ = b.cache.Delete(key)
		return nil, false
	}
	return entry.Value, true
}

func (b *BigCacheWrapper) Delete(key string) error {
	return b.cache.Delete(key)
}

func (b *BigCacheWrapper) Clear() error {
	return b.cache.Reset()
}

func (b *BigCacheWrapper) Close() error {
	// bigcache không có Close
	return nil
}
