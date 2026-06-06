package memory

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type RistrettoCache[K ~string, V any] struct {
	cache *ristretto.Cache
}

func NewRistretto[K ~string, V any]() (*RistrettoCache[K, V], error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &RistrettoCache[K, V]{cache: c}, nil
}

func (r *RistrettoCache[K, V]) Set(key K, value V, ttl time.Duration) error {
	if ok := r.cache.SetWithTTL(string(key), value, 1, ttl); !ok {
		return nil
	}
	r.cache.Wait()
	return nil
}

func (r *RistrettoCache[K, V]) Get(key K) (V, bool) {
	val, ok := r.cache.Get(string(key))
	if !ok {
		var zero V
		return zero, false
	}
	return val.(V), true
}

func (r *RistrettoCache[K, V]) Delete(key K) error {
	r.cache.Del(string(key))
	return nil
}

func (r *RistrettoCache[K, V]) Clear() error {
	r.cache.Clear()
	return nil
}

func (r *RistrettoCache[K, V]) Close() error {
	r.cache.Close()
	return nil
}
