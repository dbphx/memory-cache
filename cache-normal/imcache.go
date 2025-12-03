package memory

import (
	"time"

	"github.com/erni27/imcache"
)

type ImcacheAdapter struct {
	c *imcache.Cache[string, []byte]
}

func NewImcacheAdapter(opts ...imcache.Option[string, []byte]) *ImcacheAdapter {
	c := imcache.New[string, []byte](opts...)
	return &ImcacheAdapter{c: c}
}

func (m *ImcacheAdapter) Set(key string, value []byte, ttl time.Duration) error {
	var exp imcache.Expiration

	if ttl <= 0 {
		exp = imcache.WithNoExpiration()
	} else {
		exp = imcache.WithExpiration(ttl)
	}

	m.c.Set(key, value, exp)
	return nil
}

func (m *ImcacheAdapter) Get(key string) ([]byte, bool) {
	v, ok := m.c.Get(key)
	return v, ok
}

func (m *ImcacheAdapter) Delete(key string) error {
	_ = m.c.Remove(key)
	return nil
}

func (m *ImcacheAdapter) Clear() error {
	m.c.RemoveAll()
	return nil
}

func (m *ImcacheAdapter) Close() error {
	m.c.Close()
	return nil
}
