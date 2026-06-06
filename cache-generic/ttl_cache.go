package memory

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type TTLCache[K ~string, V any] struct {
	cache *ttlcache.Cache[string, V]
}

func NewTTLCache[K ~string, V any]() (*TTLCache[K, V], error) {
	c := ttlcache.New[string, V](
		ttlcache.WithTTL[string, V](time.Minute),
		ttlcache.WithDisableTouchOnHit[string, V](),
	)
	return &TTLCache[K, V]{cache: c}, nil
}

func (t *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	t.cache.Set(string(key), value, ttl)
	return nil
}

func (t *TTLCache[K, V]) Get(key K) (V, bool) {
	item := t.cache.Get(string(key))
	if item == nil {
		var zero V
		return zero, false
	}
	return item.Value(), true
}

func (t *TTLCache[K, V]) Delete(key K) error {
	t.cache.Delete(string(key))
	return nil
}

func (t *TTLCache[K, V]) Clear() error {
	t.cache.DeleteAll()
	return nil
}

func (t *TTLCache[K, V]) Close() error {
	return nil
}
