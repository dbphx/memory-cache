package memory

import "time"

type CacheType string

const (
	BigCache  CacheType = "bigcache"
	FreeCache CacheType = "freecache"
	Ristretto CacheType = "ristretto"
	Theine    CacheType = "theine"
	TttlCache CacheType = "ttlcache"
)

type Cache[K comparable, V any] interface {
	Set(key K, value V, ttl time.Duration) error
	Get(key K) (V, bool)
	Delete(key K) error
	Clear() error
	Close() error
}
