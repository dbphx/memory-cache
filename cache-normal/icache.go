package memory

import (
	"time"

	"github.com/mdaliyan/icache"
)

type ICachePot struct {
	pot icache.Pot
}

func NewICachePot(defaultTTL time.Duration) (*ICachePot, error) {
	return &ICachePot{
		pot: icache.NewPot(defaultTTL),
	}, nil
}

func (c *ICachePot) Set(key string, value []byte, ttl time.Duration) error {
	c.pot.Set(key, value)
	return nil
}

func (c *ICachePot) Get(key string) ([]byte, bool) {
	var b []byte
	err := c.pot.Get(key, &b)
	if err != nil {
		return nil, false
	}
	return b, true
}

func (c *ICachePot) Delete(key string) error {
	c.pot.Drop(key)
	return nil
}

func (c *ICachePot) Clear() error {
	return nil
}

func (c *ICachePot) Close() error {
	return nil
}
