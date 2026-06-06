package memory

import (
	"time"

	"github.com/Yiling-J/theine-go"
)

type TheineCache[K ~string, V any] struct {
	cache    *theine.Cache[string, V]
	capacity int
}

func NewTheineCache[K ~string, V any](capacity int) (*TheineCache[K, V], error) {
	builder := theine.NewBuilder[string, V](int64(capacity))
	cache, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return &TheineCache[K, V]{cache: cache, capacity: capacity}, nil
}

func (t *TheineCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	t.cache.SetWithTTL(string(key), value, 1, ttl)
	return nil
}

func (t *TheineCache[K, V]) Get(key K) (V, bool) {
	return t.cache.Get(string(key))
}

func (t *TheineCache[K, V]) Delete(key K) error {
	t.cache.Delete(string(key))
	return nil
}

func (t *TheineCache[K, V]) Clear() error {
	t.cache.Close()
	builder := theine.NewBuilder[string, V](int64(t.capacity))
	cache, err := builder.Build()
	if err != nil {
		return err
	}
	t.cache = cache
	return nil
}

func (t *TheineCache[K, V]) Close() error {
	t.cache.Close()
	return nil
}
