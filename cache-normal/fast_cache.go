package memory

import (
	"encoding/json"
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
	entry := bytesCacheEntry{Value: value}
	if ttl > 0 {
		entry.ExpiresAt = time.Now().Add(ttl)
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	f.cache.Set([]byte(key), data)
	return nil
}

func (f *FastCacheWrapper) Get(key string) ([]byte, bool) {
	data := f.cache.Get(nil, []byte(key))
	if data == nil {
		return nil, false
	}
	var entry bytesCacheEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, false
	}
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		f.cache.Del([]byte(key))
		return nil, false
	}
	return entry.Value, true
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
