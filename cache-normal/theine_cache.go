package memory

import (
	"time"

	"github.com/Yiling-J/theine-go"
)

type TheineCache struct {
	cache *theine.Cache[string, []byte]
}

func NewTheineCache(capacity int) (*TheineCache, error) {
	builder := theine.NewBuilder[string, []byte](int64(capacity))
	cache, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return &TheineCache{cache: cache}, nil
}

func (t *TheineCache) Set(key string, value []byte, ttl time.Duration) error {
	cost := int64(len(value))
	t.cache.SetWithTTL(key, value, cost, ttl)
	return nil
}

func (t *TheineCache) Get(key string) ([]byte, bool) {
	return t.cache.Get(key)
}

func (t *TheineCache) Delete(key string) error {
	t.cache.Delete(key)
	return nil
}

func (t *TheineCache) Clear() error {
	t.cache.Close()
	return nil
}

func (t *TheineCache) Close() error {
	t.cache.Close()
	return nil
}
