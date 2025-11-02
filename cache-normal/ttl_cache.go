package memory

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type TTLCache struct {
	cache *ttlcache.Cache[string, []byte]
}

func NewTTLCache() (*TTLCache, error) {
	c := ttlcache.New[string, []byte](
		ttlcache.WithTTL[string, []byte](time.Minute),
		ttlcache.WithDisableTouchOnHit[string, []byte](),
	)
	return &TTLCache{cache: c}, nil
}

func (t *TTLCache) Set(key string, value []byte, ttl time.Duration) error {
	t.cache.Set(key, value, ttl)
	return nil
}

func (t *TTLCache) Get(key string) ([]byte, bool) {
	item := t.cache.Get(key)
	if item == nil {
		return nil, false
	}
	return item.Value(), true
}

func (t *TTLCache) Delete(key string) error {
	t.cache.Delete(key)
	return nil
}

func (t *TTLCache) Clear() error {
	t.cache.DeleteAll()
	return nil
}

func (t *TTLCache) Close() error {
	return nil
}
