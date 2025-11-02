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

type Cache interface {
	Set(key string, value []byte, ttl time.Duration) error
	Get(key string) ([]byte, bool)
	Delete(key string) error
	Clear() error
	Close() error
}
