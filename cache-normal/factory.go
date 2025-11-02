package memory

import (
	"fmt"
)

func NewCache(t CacheType) (Cache, error) {
	switch t {
	case BigCache:
		return NewBigCache()
	case FreeCache:
		return NewFreeCache()
	case Ristretto:
		return NewRistretto()
	case Theine:
		return NewTheineCache(1_000_000)
	case TttlCache:
		return NewTTLCache()
	default:
		return nil, fmt.Errorf("unknown cache type: %v", t)
	}
}
