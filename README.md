# IPInfo Privacy Database Benchmarks

- Apple M1 Pro, 32GB
- Two different naive implementations, no real attempts to optimize memory usage
- Ingests `standard_privacy.csv.gz` in ~3 seconds
- Only works with single-address IPv4 rows from the CSV
- Performance is fine, memory usage would need scrutiny

```console
❯ go test -v -bench=.  ./...

====== MAP ======

csv ingest time = 3.219808834s

memory usage:
  virtual = 560 MB
  in use  = 532 MB

=== RUN   TestLookup_ViaMap
--- PASS: TestLookup_ViaMap (0.00s)
goos: darwin
goarch: arm64
pkg: github.com/jmalloc/ipinfo-benchmark/mapped
BenchmarkLookup_ViaMap_ArbitraryRecord
BenchmarkLookup_ViaMap_ArbitraryRecord-10      	46414405	        25.90 ns/op
BenchmarkLookup_ViaMap_NonExistentRecord
BenchmarkLookup_ViaMap_NonExistentRecord-10    	55942641	        21.44 ns/op
PASS
ok  	github.com/jmalloc/ipinfo-benchmark/mapped	6.690s

====== SLICE ======

csv ingest time = 2.65844425s

memory usage:
  virtual = 470 MB
  in use  = 435 MB

=== RUN   TestLookup_ViaSlice
--- PASS: TestLookup_ViaSlice (0.00s)
goos: darwin
goarch: arm64
pkg: github.com/jmalloc/ipinfo-benchmark/sliced
BenchmarkLookup_ViaSlice_ArbitraryRecord
BenchmarkLookup_ViaSlice_ArbitraryRecord-10      	18488263	        65.02 ns/op
BenchmarkLookup_ViaSlice_NonExistentRecord
BenchmarkLookup_ViaSlice_NonExistentRecord-10    	19994862	        60.44 ns/op
PASS
ok  	github.com/jmalloc/ipinfo-benchmark/sliced	5.605s
```
