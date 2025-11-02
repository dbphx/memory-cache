package memory

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
)

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
	return b.cache.Set(key, value)
}

func (b *BigCacheWrapper) Get(key string) ([]byte, bool) {
	val, err := b.cache.Get(key)
	if err != nil {
		return nil, false
	}
	return val, true
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
