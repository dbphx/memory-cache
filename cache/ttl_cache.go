package memory

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type TTLCache[K comparable, V any] struct {
	cache *ttlcache.Cache[K, V]
}

func NewTTLCache[K comparable, V any]() (*TTLCache[K, V], error) {
	c := ttlcache.New[K, V](
		ttlcache.WithTTL[K, V](time.Minute),
		ttlcache.WithDisableTouchOnHit[K, V](),
	)
	return &TTLCache[K, V]{cache: c}, nil
}

func (t *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	t.cache.Set(key, value, ttl)
	return nil
}

func (t *TTLCache[K, V]) Get(key K) (V, bool) {
	item := t.cache.Get(key)
	if item == nil {
		var zero V
		return zero, false
	}
	return item.Value(), true
}

func (t *TTLCache[K, V]) Delete(key K) error {
	t.cache.Delete(key)
	return nil
}

func (t *TTLCache[K, V]) Clear() error {
	t.cache.DeleteAll()
	return nil
}

func (t *TTLCache[K, V]) Close() error {
	return nil
}
