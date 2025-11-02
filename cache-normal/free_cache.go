package memory

import (
	"time"

	"github.com/coocood/freecache"
)

type FreeCacheWrapper struct {
	cache *freecache.Cache
}

func NewFreeCache() (*FreeCacheWrapper, error) {
	cacheSize := 100 * 1024 * 1024 // 100MB
	return &FreeCacheWrapper{cache: freecache.NewCache(cacheSize)}, nil
}

func (f *FreeCacheWrapper) Set(key string, value []byte, ttl time.Duration) error {
	expire := int(ttl.Seconds())
	return f.cache.Set([]byte(key), value, expire)
}

func (f *FreeCacheWrapper) Get(key string) ([]byte, bool) {
	data, err := f.cache.Get([]byte(key))
	if err != nil {
		return nil, false
	}
	return data, true
}

func (f *FreeCacheWrapper) Delete(key string) error {
	f.cache.Del([]byte(key))
	return nil
}

func (f *FreeCacheWrapper) Clear() error {
	f.cache.Clear()
	return nil
}

func (f *FreeCacheWrapper) Close() error {
	return nil
}
