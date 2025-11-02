package memory

import (
	"fmt"
	"testing"
	"time"

	cachenormal "github.com/memcache/cache-normal"
)

const (
	testTTL  = 10 * time.Second
	testSize = 10_000
)

func generateKeys(n int) []string {
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		keys[i] = fmt.Sprintf("key-%d", i)
	}
	return keys
}

func BenchmarkAllCaches(b *testing.B) {
	keys := generateKeys(testSize)
	value := []byte("value-data")

	tests := []struct {
		name string
		typ  cachenormal.CacheType
	}{
		{"BigCache", cachenormal.BigCache},
		{"FreeCache", cachenormal.FreeCache},
		{"Ristretto", cachenormal.Ristretto},
		{"Theine", cachenormal.Theine},
		{"TTLCache", cachenormal.TttlCache},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("%s_Set", tt.name), func(b *testing.B) {
			c, err := cachenormal.NewCache(tt.typ)
			if err != nil {
				b.Fatalf("failed to init cache: %v", err)
			}
			defer c.Close()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				key := keys[i%testSize]
				if err := c.Set(key, value, testTTL); err != nil {
					b.Errorf("Set failed: %v", err)
				}
			}
		})

		b.Run(fmt.Sprintf("%s_Get", tt.name), func(b *testing.B) {
			c, err := cachenormal.NewCache(tt.typ)
			if err != nil {
				b.Fatalf("failed to init cache: %v", err)
			}
			defer c.Close()

			for _, k := range keys {
				_ = c.Set(k, value, testTTL)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				key := keys[i%testSize]
				_, _ = c.Get(key)
			}
		})
	}
}
