package memory

import (
	"encoding/json"
	"time"

	"github.com/coocood/freecache"
)

type FreeCacheWrapper[K comparable, V any] struct {
	cache *freecache.Cache
}

func NewFreeCache[K comparable, V any]() (*FreeCacheWrapper[K, V], error) {
	cacheSize := 100 * 1024 * 1024 // 100MB
	return &FreeCacheWrapper[K, V]{cache: freecache.NewCache(cacheSize)}, nil
}

func (f *FreeCacheWrapper[K, V]) Set(key K, value V, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expire := int(ttl.Seconds())
	return f.cache.Set([]byte(any(key).(string)), data, expire)
}

func (f *FreeCacheWrapper[K, V]) Get(key K) (V, bool) {
	var v V
	data, err := f.cache.Get([]byte(any(key).(string)))
	if err != nil {
		return v, false
	}
	_ = json.Unmarshal(data, &v)
	return v, true
}

func (f *FreeCacheWrapper[K, V]) Delete(key K) error {
	f.cache.Del([]byte(any(key).(string)))
	return nil
}

func (f *FreeCacheWrapper[K, V]) Clear() error {
	f.cache.Clear()
	return nil
}

func (f *FreeCacheWrapper[K, V]) Close() error { return nil }
