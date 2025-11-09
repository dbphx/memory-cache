package memory

import (
	"time"

	"github.com/Yiling-J/theine-go"
)

type TheineCache[K comparable, V any] struct {
	cache *theine.Cache[K, V]
}

func NewTheineCache[K comparable, V any](capacity int) (*TheineCache[K, V], error) {
	builder := theine.NewBuilder[K, V](int64(capacity))
	cache, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return &TheineCache[K, V]{cache: cache}, nil
}

func (t *TheineCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	t.cache.SetWithTTL(key, value, 1, ttl)
	return nil
}

func (t *TheineCache[K, V]) Get(key K) (V, bool) {
	return t.cache.Get(key)
}

func (t *TheineCache[K, V]) Delete(key K) error {
	t.cache.Delete(key)
	return nil
}

func (t *TheineCache[K, V]) Clear() error {
	t.cache.Close()
	return nil
}

func (t *TheineCache[K, V]) Close() error {
	t.cache.Close()
	return nil
}
