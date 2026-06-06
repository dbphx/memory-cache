# Go Cache Benchmark

Benchmark comparison for the cache adapters in this repo.

## Overview

The repo has two benchmark suites:

- `cache-normal/bench-mark`: `[]byte` values with the full adapter set
- `cache-generic/benchmark`: generic `string -> string` wrappers

All numbers below were re-run after fixing adapter semantics around TTL, `Clear`, and sync visibility.

## Tested Libraries

Normal adapters:

- BigCache
- FreeCache
- Ristretto
- Theine
- TTLCache
- FastCache
- GoCache
- ICache
- Go2Cache
- Imcache
- Otter

Generic adapters:

- BigCache
- FreeCache
- Ristretto
- Theine
- TTLCache

## Environment

```text
goos: darwin
goarch: arm64
cpu: Apple M1 Pro
```

## Latest Results

### `cache-normal/bench-mark`

| Library | Operation | ns/op | B/op | allocs/op |
|---|---:|---:|---:|---:|
| BigCache | Set | 828.4 | 176 | 3 |
| BigCache | Get | 1116 | 368 | 8 |
| FreeCache | Set | 87.16 | 0 | 0 |
| FreeCache | Get | 127.5 | 16 | 1 |
| Ristretto | Set | 1140 | 219 | 5 |
| Ristretto | Get | 154.1 | 23 | 1 |
| Theine | Set | 191.3 | 0 | 0 |
| Theine | Get | 142.6 | 16 | 1 |
| TTLCache | Set | 387.1 | 0 | 0 |
| TTLCache | Get | 103.1 | 0 | 0 |
| FastCache | Set | 507.4 | 176 | 3 |
| FastCache | Get | 824.2 | 360 | 7 |
| GoCache | Set | 101.2 | 24 | 1 |
| GoCache | Get | 58.87 | 0 | 0 |
| ICache | Set | 241.3 | 122 | 2 |
| ICache | Get | 158.6 | 80 | 2 |
| Go2Cache | Set | 200.6 | 208 | 5 |
| Go2Cache | Get | 107.0 | 16 | 1 |
| Imcache | Set | 106.2 | 16 | 1 |
| Imcache | Get | 93.23 | 0 | 0 |
| Otter | Set | 371.4 | 101 | 1 |
| Otter | Get | 100.8 | 0 | 0 |

Raw output:

```text
goos: darwin
goarch: arm64
pkg: github.com/memcache/cache-normal/bench-mark
cpu: Apple M1 Pro
BenchmarkAllCaches/BigCache_Set-8         	 1428632	       828.4 ns/op	     176 B/op	       3 allocs/op
BenchmarkAllCaches/BigCache_Get-8         	 1000000	      1116 ns/op	     368 B/op	       8 allocs/op
BenchmarkAllCaches/FreeCache_Set-8        	14128090	        87.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/FreeCache_Get-8        	 9321328	       127.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/Ristretto_Set-8        	 1000000	      1140 ns/op	     219 B/op	       5 allocs/op
BenchmarkAllCaches/Ristretto_Get-8        	 7964700	       154.1 ns/op	      23 B/op	       1 allocs/op
BenchmarkAllCaches/Theine_Set-8           	 5954248	       191.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Theine_Get-8           	 8476034	       142.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/TTLCache_Set-8         	 3149944	       387.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/TTLCache_Get-8         	11721253	       103.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/FastCache_Set-8        	 2391687	       507.4 ns/op	     176 B/op	       3 allocs/op
BenchmarkAllCaches/FastCache_Get-8        	 1475431	       824.2 ns/op	     360 B/op	       7 allocs/op
BenchmarkAllCaches/GoCache_Set-8          	11860740	       101.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAllCaches/GoCache_Get-8          	20456942	        58.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Icache_Set-8           	 5102410	       241.3 ns/op	     122 B/op	       2 allocs/op
BenchmarkAllCaches/Icache_Get-8           	 7733790	       158.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkAllCaches/Go2Cache_Set-8         	 6385521	       200.6 ns/op	     208 B/op	       5 allocs/op
BenchmarkAllCaches/Go2Cache_Get-8         	11285367	       107.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/Imcache_Set-8          	11378505	       106.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/Imcache_Get-8          	12905098	        93.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Otter_Set-8            	 3183435	       371.4 ns/op	     101 B/op	       1 allocs/op
BenchmarkAllCaches/Otter_Get-8            	12131055	       100.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/memcache/cache-normal/bench-mark	33.047s
```

### `cache-generic/benchmark`

| Library | Operation | ns/op | B/op | allocs/op |
|---|---:|---:|---:|---:|
| BigCache | Set | 987.6 | 176 | 3 |
| BigCache | Get | 1058 | 368 | 8 |
| FreeCache | Set | 219.9 | 32 | 2 |
| FreeCache | Get | 303.3 | 192 | 4 |
| Ristretto | Set | 1346 | 212 | 5 |
| Ristretto | Get | 153.7 | 23 | 1 |
| Theine | Set | 202.6 | 0 | 0 |
| Theine | Get | 138.6 | 16 | 1 |
| TTLCache | Set | 387.4 | 0 | 0 |
| TTLCache | Get | 103.3 | 0 | 0 |

Raw output:

```text
goos: darwin
goarch: arm64
pkg: github.com/memcache/cache-generic/benchmark
cpu: Apple M1 Pro
BenchmarkAllCaches/BigCache_Set-8         	 1519365	       987.6 ns/op	     176 B/op	       3 allocs/op
BenchmarkAllCaches/BigCache_Get-8         	 1130392	      1058 ns/op	     368 B/op	       8 allocs/op
BenchmarkAllCaches/FreeCache_Set-8        	 5784136	       219.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkAllCaches/FreeCache_Get-8        	 4359615	       303.3 ns/op	     192 B/op	       4 allocs/op
BenchmarkAllCaches/Ristretto_Set-8        	  789294	      1346 ns/op	     212 B/op	       5 allocs/op
BenchmarkAllCaches/Ristretto_Get-8        	 7669881	       153.7 ns/op	      23 B/op	       1 allocs/op
BenchmarkAllCaches/Theine_Set-8           	 6115185	       202.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Theine_Get-8           	 8534083	       138.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/TTLCache_Set-8         	 3129807	       387.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/TTLCache_Get-8         	11861325	       103.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/memcache/cache-generic/benchmark	17.264s
```

## How to Run

```bash
go test ./cache-normal/bench-mark -bench=. -benchmem -run=^$
go test ./cache-generic/benchmark -bench=. -benchmem -run=^$
```

To filter a specific cache:

```bash
go test ./cache-normal/bench-mark -bench=BigCache -benchmem -run=^$
```

## Notes

- Benchmarks were run with `go test -bench=. -benchmem -run=^$`.
- The `Get` benchmark now verifies preload succeeded before timing begins.
- Several adapters now honor TTL in the wrapper layer, so the numbers are not directly comparable to the old README results.
- Generic wrappers for `BigCache` and `FreeCache` include JSON serialization overhead.

## Takeaways

- Fastest read path in the normal suite: `GoCache` at `58.87 ns/op`.
- Fastest write path in the normal suite: `FreeCache` at `87.16 ns/op`.
- Lowest-allocation read path in both suites: `TTLCache` (`0 B/op`, `0 allocs/op`).
- Generic wrappers are materially slower for `BigCache` and `FreeCache` because they serialize values.
