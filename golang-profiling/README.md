## Setting up benchmarks

Benchmarks are just functions.
These functions need to be inside of a test file meaning the file should end with `_test`.
Then the benchmark function should follow the format below:

```golang
func BenchmarkMyFunction(b *testing.B) {
    ...
}
```

## Running the benchmark

```bash
go test -bench=. -benchmem
```
