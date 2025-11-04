package memory

import (
	"errors"
	"time"

	"github.com/VictoriaMetrics/fastcache"
)

type FastCacheWrapper struct {
	cache *fastcache.Cache
}

func NewFastCacheWrapper(sizeInBytes int) (*FastCacheWrapper, error) {
	return &FastCacheWrapper{
		cache: fastcache.New(sizeInBytes),
	}, nil
}

func (f *FastCacheWrapper) Set(key string, value []byte, ttl time.Duration) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	f.cache.Set([]byte(key), value)
	return nil
}

func (f *FastCacheWrapper) Get(key string) ([]byte, bool) {
	val := f.cache.Get(nil, []byte(key))
	if val == nil {
		return nil, false
	}
	return val, true
}

func (f *FastCacheWrapper) Delete(key string) error {
	f.cache.Del([]byte(key))
	return nil
}

func (f *FastCacheWrapper) Clear() error {
	f.cache.Reset()
	return nil
}

func (f *FastCacheWrapper) Close() error {
	f.cache = nil
	return nil
}
