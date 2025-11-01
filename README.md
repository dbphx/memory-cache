# 🧠 Go Cache Benchmark

Benchmark performance comparison between popular in-memory cache libraries in Go.

## 📋 Overview

This project benchmarks common Go caching libraries under identical workloads for both **Set** and **Get** operations.  
The goal is to evaluate **speed**, **memory efficiency**, and **allocation behavior**.

## 🧪 Tested Libraries

- [BigCache](https://github.com/allegro/bigcache)
- [FreeCache](https://github.com/coocood/freecache)
- [Ristretto](https://github.com/dgraph-io/ristretto)
- [Theine](https://github.com/yujunz/theine-go)
- [TTLCache](https://github.com/jellydator/ttlcache)

## ⚙️ Environment

```
goos: darwin
goarch: arm64
cpu: Apple M1 Pro
pkg: github.com/memcache/benchmark
```

## 🚀 Benchmark Results

| Library       | Operation | ns/op | B/op | allocs/op |
|----------------|------------|-------|------|------------|
| **BigCache**   | Set        | 329.5 | 32   | 2 |
|                | Get        | 356.5 | 199  | 5 |
| **FreeCache**  | Set        | 235.2 | 32   | 2 |
|                | Get        | 361.5 | 192  | 4 |
| **Ristretto**  | Set        | 428.6 | 113  | 3 |
|                | Get        | 176.3 | 24   | 1 |
| **Theine**     | Set        | 226.2 | 0    | 0 |
|                | Get        | 176.2 | 16   | 1 |
| **TTLCache**   | Set        | 486.5 | 1    | 0 |
|                | Get        | 132.5 | 0    | 0 |

✅ **Fastest Get:** TTLCache (132.5 ns/op)  
⚡ **Fastest Set:** Theine (226.2 ns/op)  
💾 **Lowest Allocations:** Theine / TTLCache

## 🧰 How to Run

```bash
# Clone the repo
git clone https://github.com/memcache/benchmark.git
cd benchmark

# Run all benchmarks
go test -bench=. -benchmem
```

To filter a specific cache:
```bash
go test -bench=BigCache -benchmem
```

## 📊 Notes

- All benchmarks run with `GOMAXPROCS=8`
- Data size and access patterns are uniform
- Each test includes memory allocation profiling (`-benchmem`)

## 🧭 Conclusion

| Use Case | Recommended Library |
|-----------|--------------------|
| Ultra-low latency reads | **TTLCache** |
| Balanced performance | **Theine** |
| Heavy concurrency, medium latency | **FreeCache** |
| Cache with admission/eviction policies | **Ristretto** |
| Simple key-value cache | **BigCache** |

