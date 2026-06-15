package memory

import (
	"fmt"
)

func NewCache[K ~string, V any](t CacheType) (Cache[K, V], error) {
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
	case SyncMap:
		return NewSyncMapCache[K, V](), nil
	case MutexMap:
		return NewMutexMapCache[K, V](), nil
	default:
		return nil, fmt.Errorf("unknown cache type: %v", t)
	}
}
