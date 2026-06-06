package memory

import (
	"time"

	"github.com/mdaliyan/icache"
)

type icacheEntry struct {
	Value     []byte
	ExpiresAt time.Time
}

type ICachePot struct {
	pot icache.Pot
}

func NewICachePot(defaultTTL time.Duration) (*ICachePot, error) {
	return &ICachePot{
		pot: icache.NewPot(defaultTTL),
	}, nil
}

func (c *ICachePot) Set(key string, value []byte, ttl time.Duration) error {
	entry := icacheEntry{Value: value}
	if ttl > 0 {
		entry.ExpiresAt = time.Now().Add(ttl)
	}
	c.pot.Set(key, entry)
	return nil
}

func (c *ICachePot) Get(key string) ([]byte, bool) {
	var entry icacheEntry
	err := c.pot.Get(key, &entry)
	if err != nil {
		return nil, false
	}
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		c.pot.Drop(key)
		return nil, false
	}
	return entry.Value, true
}

func (c *ICachePot) Delete(key string) error {
	c.pot.Drop(key)
	return nil
}

func (c *ICachePot) Clear() error {
	c.pot.Purge()
	return nil
}

func (c *ICachePot) Close() error {
	c.pot.Purge()
	return nil
}
