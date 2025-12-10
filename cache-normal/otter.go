package memory

import (
	"time"

	"github.com/maypok86/otter/v2"
)

type OtterV2Cache struct {
	back *otter.Cache[string, []byte]
}

func NewOtterV2Cache(maxEntries int) *OtterV2Cache {
	c := otter.Must(&otter.Options[string, []byte]{
		MaximumSize: maxEntries,
	})
	return &OtterV2Cache{back: c}
}

func (c *OtterV2Cache) Set(key string, value []byte, ttl time.Duration) error {
	c.back.Set(key, value)
	if ttl > 0 {
		c.back.SetExpiresAfter(key, ttl)
	}
	return nil
}

func (c *OtterV2Cache) Get(key string) ([]byte, bool) {
	v, ok := c.back.GetIfPresent(key)
	return v, ok
}

func (c *OtterV2Cache) Delete(key string) error {
	c.back.Invalidate(key)
	return nil
}

func (c *OtterV2Cache) Clear() error {
	c.back.InvalidateAll()
	return nil
}

func (c *OtterV2Cache) Close() error {
	c.back.InvalidateAll()
	return nil
}
