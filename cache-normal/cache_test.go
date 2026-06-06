package memory

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func normalCacheTypes() []CacheType {
	return []CacheType{
		BigCache,
		FreeCache,
		Ristretto,
		Theine,
		TttlCache,
		FastCache,
		GoCache,
		ICache,
		Go2Cache,
		Imcache,
		Otter,
	}
}

func TestNormalCachesBasicOperations(t *testing.T) {
	for _, typ := range normalCacheTypes() {
		t.Run(string(typ), func(t *testing.T) {
			t.Parallel()

			cache, err := NewCache(typ)
			if err != nil {
				t.Fatalf("NewCache(%s): %v", typ, err)
			}
			defer cache.Close()

			key := fmt.Sprintf("%s-basic", typ)
			value := []byte("value")

			if err := cache.Set(key, value, 5*time.Second); err != nil {
				t.Fatalf("Set(%s): %v", typ, err)
			}

			got, ok := cache.Get(key)
			if !ok {
				t.Fatalf("Get(%s): expected hit", typ)
			}
			if !bytes.Equal(got, value) {
				t.Fatalf("Get(%s): got %q want %q", typ, got, value)
			}

			if err := cache.Delete(key); err != nil {
				t.Fatalf("Delete(%s): %v", typ, err)
			}
			if _, ok := cache.Get(key); ok {
				t.Fatalf("Get(%s) after delete: expected miss", typ)
			}
		})
	}
}

func TestNormalCachesClearKeepsCacheReusable(t *testing.T) {
	for _, typ := range normalCacheTypes() {
		t.Run(string(typ), func(t *testing.T) {
			t.Parallel()

			cache, err := NewCache(typ)
			if err != nil {
				t.Fatalf("NewCache(%s): %v", typ, err)
			}
			defer cache.Close()

			if err := cache.Set("before-clear", []byte("v1"), 5*time.Second); err != nil {
				t.Fatalf("Set before clear (%s): %v", typ, err)
			}
			if err := cache.Clear(); err != nil {
				t.Fatalf("Clear(%s): %v", typ, err)
			}
			if _, ok := cache.Get("before-clear"); ok {
				t.Fatalf("Get(%s) after clear: expected miss", typ)
			}

			if err := cache.Set("after-clear", []byte("v2"), 5*time.Second); err != nil {
				t.Fatalf("Set after clear (%s): %v", typ, err)
			}
			if got, ok := cache.Get("after-clear"); !ok || !bytes.Equal(got, []byte("v2")) {
				t.Fatalf("Get(%s) after clear reuse: expected hit", typ)
			}
		})
	}
}

func TestNormalCachesTTLExpiry(t *testing.T) {
	for _, typ := range normalCacheTypes() {
		t.Run(string(typ), func(t *testing.T) {
			t.Parallel()

			cache, err := NewCache(typ)
			if err != nil {
				t.Fatalf("NewCache(%s): %v", typ, err)
			}
			defer cache.Close()

			key := fmt.Sprintf("%s-ttl", typ)
			if err := cache.Set(key, []byte("value"), 1100*time.Millisecond); err != nil {
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
