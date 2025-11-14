package memory

import (
	"time"

	cache2go "github.com/muesli/cache2go"
)

type cache2goAdapter struct {
	cache *cache2go.CacheTable
}

func NewCache2Go() (Cache, error) {
	return &cache2goAdapter{
		cache: cache2go.Cache("cache2go"),
	}, nil
}

func (c *cache2goAdapter) Set(key string, value []byte, ttl time.Duration) error {
	c.cache.Add(key, ttl, value)
	return nil
}

func (c *cache2goAdapter) Get(key string) ([]byte, bool) {
	item, err := c.cache.Value(key)
	if err != nil || item == nil {
		return nil, false
	}
	v, ok := item.Data().([]byte)
	if !ok {
		return nil, false
	}
	return v, true
}

func (c *cache2goAdapter) Delete(key string) error {
	c.cache.Delete(key)
	return nil
}

func (c *cache2goAdapter) Clear() error {
	c.cache.Flush()
	return nil
}

func (c *cache2goAdapter) Close() error {
	c.cache.Flush()
	return nil
}
