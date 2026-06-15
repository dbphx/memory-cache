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
- SyncMap
- MutexMap

Generic adapters:

- BigCache
- FreeCache
- Ristretto
- Theine
- TTLCache
- SyncMap
- MutexMap

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
| BigCache | Set | 868.7 | 176 | 3 |
| BigCache | Get | 1018 | 368 | 8 |
| FreeCache | Set | 87.58 | 0 | 0 |
| FreeCache | Get | 139.7 | 16 | 1 |
| Ristretto | Set | 1391 | 221 | 5 |
| Ristretto | Get | 161.4 | 23 | 1 |
| Theine | Set | 183.9 | 0 | 0 |
| Theine | Get | 144.3 | 16 | 1 |
| TTLCache | Set | 403.3 | 1 | 0 |
| TTLCache | Get | 103.8 | 0 | 0 |
| FastCache | Set | 588.1 | 176 | 3 |
| FastCache | Get | 888.1 | 360 | 7 |
| GoCache | Set | 108.2 | 24 | 1 |
| GoCache | Get | 63.15 | 0 | 0 |
| ICache | Set | 249.6 | 123 | 2 |
| ICache | Get | 164.8 | 80 | 2 |
| Go2Cache | Set | 197.1 | 208 | 5 |
| Go2Cache | Get | 107.5 | 16 | 1 |
| Imcache | Set | 106.0 | 16 | 1 |
| Imcache | Get | 97.12 | 0 | 0 |
| Otter | Set | 390.3 | 100 | 1 |
| Otter | Get | 96.05 | 0 | 0 |
| SyncMap | Set | 177.8 | 128 | 4 |
| SyncMap | Get | 91.42 | 16 | 1 |
| MutexMap | Set | 86.11 | 16 | 1 |
| MutexMap | Get | 98.18 | 16 | 1 |

Raw output:

```text
goos: darwin
goarch: arm64
pkg: github.com/memcache/cache-normal/bench-mark
cpu: Apple M1 Pro
BenchmarkAllCaches/BigCache_Set-8         	 1404957	       868.7 ns/op	     176 B/op	       3 allocs/op
BenchmarkAllCaches/BigCache_Get-8         	  996652	      1018 ns/op	     368 B/op	       8 allocs/op
BenchmarkAllCaches/FreeCache_Set-8        	11476572	        87.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/FreeCache_Get-8        	 8442313	       139.7 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/Ristretto_Set-8        	  757081	      1391 ns/op	     221 B/op	       5 allocs/op
BenchmarkAllCaches/Ristretto_Get-8        	 7633647	       161.4 ns/op	      23 B/op	       1 allocs/op
BenchmarkAllCaches/Theine_Set-8           	 6380754	       183.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Theine_Get-8           	 7496114	       144.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/TTLCache_Set-8         	 3036314	       403.3 ns/op	       1 B/op	       0 allocs/op
BenchmarkAllCaches/TTLCache_Get-8         	11531992	       103.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/FastCache_Set-8        	 2384817	       588.1 ns/op	     176 B/op	       3 allocs/op
BenchmarkAllCaches/FastCache_Get-8        	 1336828	       888.1 ns/op	     360 B/op	       7 allocs/op
BenchmarkAllCaches/GoCache_Set-8          	 9708838	       108.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkAllCaches/GoCache_Get-8          	19967455	        63.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Icache_Set-8           	 4969292	       249.6 ns/op	     123 B/op	       2 allocs/op
BenchmarkAllCaches/Icache_Get-8           	 7213393	       164.8 ns/op	      80 B/op	       2 allocs/op
BenchmarkAllCaches/Go2Cache_Set-8         	 6138344	       197.1 ns/op	     208 B/op	       5 allocs/op
BenchmarkAllCaches/Go2Cache_Get-8         	11163894	       107.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/Imcache_Set-8          	11422342	       106.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/Imcache_Get-8          	12205189	        97.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Otter_Set-8            	 3127407	       390.3 ns/op	     100 B/op	       1 allocs/op
BenchmarkAllCaches/Otter_Get-8            	12320733	        96.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/SyncMap_Set-8          	 6996861	       177.8 ns/op	     128 B/op	       4 allocs/op
BenchmarkAllCaches/SyncMap_Get-8          	13148961	        91.42 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/MutexMap_Set-8         	14224096	        86.11 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/MutexMap_Get-8         	11651427	        98.18 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/memcache/cache-normal/bench-mark	38.855s
```

### `cache-generic/benchmark`

| Library | Operation | ns/op | B/op | allocs/op |
|---|---:|---:|---:|---:|
| BigCache | Set | 867.4 | 176 | 3 |
| BigCache | Get | 1093 | 368 | 8 |
| FreeCache | Set | 227.4 | 32 | 2 |
| FreeCache | Get | 311.8 | 192 | 4 |
| Ristretto | Set | 1415 | 212 | 5 |
| Ristretto | Get | 158.4 | 23 | 1 |
| Theine | Set | 206.0 | 0 | 0 |
| Theine | Get | 140.7 | 16 | 1 |
| TTLCache | Set | 401.3 | 1 | 0 |
| TTLCache | Get | 103.2 | 0 | 0 |
| SyncMap | Set | 178.4 | 112 | 3 |
| SyncMap | Get | 74.85 | 0 | 0 |
| MutexMap | Set | 71.81 | 0 | 0 |
| MutexMap | Get | 65.15 | 0 | 0 |

Raw output:

```text
goos: darwin
goarch: arm64
pkg: github.com/memcache/cache-generic/benchmark
cpu: Apple M1 Pro
BenchmarkAllCaches/BigCache_Set-8         	 1229905	       867.4 ns/op	     176 B/op	       3 allocs/op
BenchmarkAllCaches/BigCache_Get-8         	 1000000	      1093 ns/op	     368 B/op	       8 allocs/op
BenchmarkAllCaches/FreeCache_Set-8        	 5407934	       227.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkAllCaches/FreeCache_Get-8        	 3913832	       311.8 ns/op	     192 B/op	       4 allocs/op
BenchmarkAllCaches/Ristretto_Set-8        	  886327	      1415 ns/op	     212 B/op	       5 allocs/op
BenchmarkAllCaches/Ristretto_Get-8        	 7462737	       158.4 ns/op	      23 B/op	       1 allocs/op
BenchmarkAllCaches/Theine_Set-8           	 5948756	       206.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/Theine_Get-8           	 8552840	       140.7 ns/op	      16 B/op	       1 allocs/op
BenchmarkAllCaches/TTLCache_Set-8         	 2905779	       401.3 ns/op	       1 B/op	       0 allocs/op
BenchmarkAllCaches/TTLCache_Get-8         	11595308	       103.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/SyncMap_Set-8          	 8237158	       178.4 ns/op	     112 B/op	       3 allocs/op
BenchmarkAllCaches/SyncMap_Get-8          	15123158	        74.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/MutexMap_Set-8         	17905436	        71.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllCaches/MutexMap_Get-8         	18414356	        65.15 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/memcache/cache-generic/benchmark	21.112s
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

- Fastest read path in the normal suite: `GoCache` at `63.15 ns/op`.
- Fastest write path in the normal suite: `MutexMap` at `86.11 ns/op`, narrowly ahead of `FreeCache`.
- Fastest generic adapter in both read and write: `MutexMap` (`65.15 ns/op` get, `71.81 ns/op` set).
- Lowest-allocation read path in both suites includes `TTLCache`, and in the generic suite also `SyncMap` and `MutexMap` (`0 B/op`, `0 allocs/op`).
- Generic wrappers are materially slower for `BigCache` and `FreeCache` because they serialize values.
