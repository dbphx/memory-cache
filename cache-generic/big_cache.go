package memory

import (
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

type BigCacheWrapper[K comparable, V any] struct {
	cache *bigcache.BigCache
}

func NewBigCache[K comparable, V any]() (*BigCacheWrapper[K, V], error) {
	c, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		return nil, err
	}
	return &BigCacheWrapper[K, V]{cache: c}, nil
}

func (b *BigCacheWrapper[K, V]) Set(key K, value V, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return b.cache.Set(any(key).(string), data)
}

func (b *BigCacheWrapper[K, V]) Get(key K) (V, bool) {
	var v V
	data, err := b.cache.Get(any(key).(string))
	if err != nil {
		return v, false
	}
	_ = json.Unmarshal(data, &v)
	return v, true
}

func (b *BigCacheWrapper[K, V]) Delete(key K) error {
	return b.cache.Delete(any(key).(string))
}

func (b *BigCacheWrapper[K, V]) Clear() error {
	return b.cache.Reset()
}

func (b *BigCacheWrapper[K, V]) Close() error {
	return nil
}
