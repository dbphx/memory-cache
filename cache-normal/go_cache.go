package memory

import (
	"time"

	gocache "github.com/yuseferi/gocache"
)

type goCacheAdapter struct {
	cache *gocache.Cache
}

func NewGoCache() (Cache, error) {
	return &goCacheAdapter{
		cache: gocache.NewCache(time.Minute),
	}, nil
}

func (g *goCacheAdapter) Set(key string, value []byte, ttl time.Duration) error {
	g.cache.Set(key, value, ttl)
	return nil
}

func (g *goCacheAdapter) Get(key string) ([]byte, bool) {
	v, ok := g.cache.Get(key)
	if !ok {
		return nil, false
	}
	return v.([]byte), true
}

func (g *goCacheAdapter) Delete(key string) error {
	g.cache.Delete(key)
	return nil
}

func (g *goCacheAdapter) Clear() error {
	g.cache.Clear()
	return nil
}

func (g *goCacheAdapter) Close() error {
	return nil
}
