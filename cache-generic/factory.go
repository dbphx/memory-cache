package memory

import (
	"fmt"
)

func NewCache[K comparable, V any](t CacheType) (Cache[K, V], error) {
	switch t {
	case BigCache:
		return NewBigCache[K, V]()
	case FreeCache:
		return NewFreeCache[K, V]()
	case Ristretto:
		return NewRistretto[K, V]()
	case Theine:
		return NewTheineCache[K, V](1_000_000)
	case TttlCache:
		return NewTTLCache[K, V]()
	default:
		return nil, fmt.Errorf("unknown cache type: %v", t)
	}
}
