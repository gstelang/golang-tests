# benchmarks

```
go test ./benchmarks -bench=. --benchmem

goos: darwin
goarch: arm64
pkg: github.com/gstelang/golang-tests/benchmarks
cpu: Apple M1
BenchmarkFactorialRecursive-8           55954006                20.98 ns/op            0 B/op          0 allocs/op
BenchmarkFactorialIterative-8           165931000                7.201 ns/op           0 B/op          0 allocs/op
BenchmarkStringConcat_Plus-8             4885476               244.2 ns/op           616 B/op          9 allocs/op
BenchmarkStringConcat_Builder-8         11632556               101.3 ns/op           240 B/op          4 allocs/op
PASS
ok      github.com/gstelang/golang-tests/benchmarks     6.680s

notes: 
* Columns
    1. 2nd column is number of iterations
    2. 3rd column is avg time per operation in nanoseconds
    3. number of bytes per operation
    4. Number of allocs per operation
* without benchmem won't report 3rd and 4th column

```

# Other things to experiment
```
go test ./benchmarks -bench=BenchmarkFactorialRecursive -cpu=1   // how many CPUs
go test ./benchmarks -bench=BenchmarkFactorialRecursive -count=5 // Run each benchmark 5 times.
go test ./benchmarks -bench=BenchmarkFactorialRecursive -benchtime=100000x // changes b.N OR 
        // -benchtime=10s: Run each benchmark for 10 seconds
```