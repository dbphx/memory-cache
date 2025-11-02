package memory

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type RistrettoCache struct {
	cache *ristretto.Cache
}

func NewRistretto() (*RistrettoCache, error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &RistrettoCache{cache: c}, nil
}

func (r *RistrettoCache) Set(key string, value []byte, ttl time.Duration) error {
	cost := int64(len(value))
	r.cache.SetWithTTL(key, value, cost, ttl)
	// ristretto set async → return nil immediately
	return nil
}

func (r *RistrettoCache) Get(key string) ([]byte, bool) {
	val, ok := r.cache.Get(key)
	if !ok {
		return nil, false
	}
	return val.([]byte), true
}

func (r *RistrettoCache) Delete(key string) error {
	r.cache.Del(key)
	return nil
}

func (r *RistrettoCache) Clear() error {
	r.cache.Clear()
	return nil
}

func (r *RistrettoCache) Close() error {
	r.cache.Close()
	return nil
}
