# 🧠 Go Cache Benchmark

Benchmark performance comparison between popular in-memory cache libraries in Go.

## 📋 Overview

This project benchmarks common Go caching libraries under identical workloads for both **Set** and **Get** operations.  
The goal is to evaluate **speed**, **memory efficiency**, and **allocation behavior**.

## 🧪 Tested Libraries

- BigCache
- FreeCache
- Ristretto
- Theine
- TTLCache
- FastCache
- GoCache

## ⚙️ Environment

```
goos: darwin
goarch: arm64
cpu: Apple M1 Pro
pkg: github.com/memcache/benchmark
```

## 🚀 Benchmark Results

| Library       | Operation | ns/op | B/op | allocs/op |
|---------------|-----------|-------|------|-----------|
| **BigCache**  | Set       | 208.1 | 1    | 0         |
|               | Get       | 93.47 | 23   | 2         |
| **FreeCache** | Set       | 71.20 | 0    | 0         |
|               | Get       | 100.6 | 16   | 1         |
| **Ristretto** | Set       | 349.1 | 121  | 3         |
|               | Get       | 141.4 | 23   | 1         |
| **Theine**    | Set       | 168.0 | 0    | 0         |
|               | Get       | 134.9 | 16   | 1         |
| **TTLCache**  | Set       | 392.0 | 0    | 0         |
|               | Get       | 100.1 | 0    | 0         |
| **FastCache** | Set       | 59.55 | 0    | 0         |
|               | Get       | 100.8 | 16   | 1         |
| **GoCache**   | Set       | 131.7 | 24   | 1         |
|               | Get       | 81.41 | 0    | 0         |
| **ICache**    | Set       | 223.8 | 80   | 2         |
|               | Get       | 132.6 | 40   | 2         |
| **Go2Cache**  | Set       | 188.6 | 208  | 5         |
|               | Get       | 102.7 | 16   | 1         |
raw benchmark:

```
BenchmarkAllCaches/BigCache_Set-8                5033186               208.1 ns/op             1 B/op          0 allocs/op
BenchmarkAllCaches/BigCache_Get-8               12522745                93.47 ns/op           23 B/op          2 allocs/op
BenchmarkAllCaches/FreeCache_Set-8              16266094                71.20 ns/op            0 B/op          0 allocs/op
BenchmarkAllCaches/FreeCache_Get-8              11706307               100.6 ns/op            16 B/op          1 allocs/op
BenchmarkAllCaches/Ristretto_Set-8               2962323               349.1 ns/op           121 B/op          3 allocs/op
BenchmarkAllCaches/Ristretto_Get-8               8604910               141.4 ns/op            23 B/op          1 allocs/op
BenchmarkAllCaches/Theine_Set-8                  6353618               168.0 ns/op             0 B/op          0 allocs/op
BenchmarkAllCaches/Theine_Get-8                  8893598               134.9 ns/op            16 B/op          1 allocs/op
BenchmarkAllCaches/TTLCache_Set-8                3159577               392.0 ns/op             0 B/op          0 allocs/op
BenchmarkAllCaches/TTLCache_Get-8               11683659               100.1 ns/op             0 B/op          0 allocs/op
BenchmarkAllCaches/FastCache_Set-8              19292216                58.70 ns/op            0 B/op          0 allocs/op
BenchmarkAllCaches/FastCache_Get-8              11788620               102.1 ns/op            16 B/op          1 allocs/op
BenchmarkAllCaches/GoCache_Set-8                 8923922               131.7 ns/op            24 B/op          1 allocs/op
BenchmarkAllCaches/GoCache_Get-8                15443277                81.41 ns/op            0 B/op          0 allocs/op
BenchmarkAllCaches/Icache_Set-8                  5346549               223.8 ns/op            80 B/op          2 allocs/op
BenchmarkAllCaches/Icache_Get-8                  9755575               132.6 ns/op            40 B/op          2 allocs/op
BenchmarkAllCaches/Go2Cache_Set-8                6685914               188.6 ns/op           208 B/op          5 allocs/op
BenchmarkAllCaches/Go2Cache_Get-8               11563836               102.7 ns/op            16 B/op          1 allocs/op
```

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
|-----------|---------------------|
| Ultra-low latency reads | **TTLCache**        |
| Balanced performance | **FastCache**       |
| Heavy concurrency, medium latency | **FreeCache**       |
| Cache with admission/eviction policies | **Ristretto**       |
| Simple key-value cache | **BigCache**        |
