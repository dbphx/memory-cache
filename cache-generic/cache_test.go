package memory

import (
	"fmt"
	"testing"
	"time"
)

type aliasKey string

func genericCacheTypes() []CacheType {
	return []CacheType{
		BigCache,
		FreeCache,
		Ristretto,
		Theine,
		TttlCache,
	}
}

func TestGenericCachesSupportStringAliases(t *testing.T) {
	for _, typ := range genericCacheTypes() {
		t.Run(string(typ), func(t *testing.T) {
			t.Parallel()

			cache, err := NewCache[aliasKey, string](typ)
			if err != nil {
				t.Fatalf("NewCache(%s): %v", typ, err)
			}
			defer cache.Close()

			key := aliasKey(fmt.Sprintf("%s-key", typ))
			if err := cache.Set(key, "value", 5*time.Second); err != nil {
				t.Fatalf("Set(%s): %v", typ, err)
			}

			got, ok := cache.Get(key)
			if !ok || got != "value" {
				t.Fatalf("Get(%s): got=%q ok=%v", typ, got, ok)
			}

			if err := cache.Clear(); err != nil {
				t.Fatalf("Clear(%s): %v", typ, err)
			}
			if _, ok := cache.Get(key); ok {
				t.Fatalf("Get(%s) after clear: expected miss", typ)
			}

			if err := cache.Set(key, "value-2", 5*time.Second); err != nil {
				t.Fatalf("Set after clear (%s): %v", typ, err)
			}
			if got, ok := cache.Get(key); !ok || got != "value-2" {
				t.Fatalf("Get(%s) after clear reuse: got=%q ok=%v", typ, got, ok)
			}
		})
	}
}

func TestGenericCachesTTLExpiry(t *testing.T) {
	for _, typ := range genericCacheTypes() {
		t.Run(string(typ), func(t *testing.T) {
			t.Parallel()

			cache, err := NewCache[string, string](typ)
			if err != nil {
				t.Fatalf("NewCache(%s): %v", typ, err)
			}
			defer cache.Close()

			key := fmt.Sprintf("%s-ttl", typ)
			if err := cache.Set(key, "value", 1100*time.Millisecond); err != nil {
				t.Fatalf("Set(%s): %v", typ, err)
			}
			if _, ok := cache.Get(key); !ok {
				t.Fatalf("Get(%s) immediately after set: expected hit", typ)
			}

			time.Sleep(1500 * time.Millisecond)

			if _, ok := cache.Get(key); ok {
				t.Fatalf("Get(%s) after ttl: expected miss", typ)
			}
		})
	}
}
