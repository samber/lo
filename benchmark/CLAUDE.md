# Benchmark Guidelines

## File Organization

Benchmark files follow the naming convention:

```
benchmark/{package}_{category}_bench_test.go
```

- **package**: `core`, `it`, `mutable`, `parallel`
- **category**: `slice`, `map`, `find`, `intersect`, `math`, `string`, `type_manipulation`, `condition`, `tuples`

Shared data generators live in `helpers_test.go` (and `it_helpers_test.go` for `go1.23` iter helpers).

## Performance PRs

Every performance improvement PR **must** include a `benchstat` comparison in the PR description. Without before/after numbers, the PR will not be merged.

### How to produce a benchstat report

1. Check out `master` and run the "before" benchmarks:
   ```bash
   git stash && git switch master
   go test ./benchmark/... -bench=BenchmarkXxx -benchmem -count=6 -cpu=1 | tee /tmp/before.txt
   ```

2. Switch to your branch and run the "after" benchmarks:
   ```bash
   git switch my-branch && git stash pop
   go test ./benchmark/... -bench=BenchmarkXxx -benchmem -count=6 -cpu=1 | tee /tmp/after.txt
   ```

3. Compare with `benchstat`:
   ```bash
   benchstat /tmp/before.txt /tmp/after.txt
   ```

4. Paste the full `benchstat` output in the PR description.

### What to include in the PR description

- The optimization technique (pre-allocation, direct indexing, value receivers, etc.)
- The `benchstat` table showing time/op, allocs/op, and bytes/op deltas
- An explanation of **why** the change is faster, not just **what** changed

### When NOT to submit a performance PR

- If `benchstat` shows no statistically significant improvement (p >= 0.05)
- If the improvement is < 5% and adds code complexity
- If the change regresses other benchmarks — always run the full suite, not just the targeted benchmark

## Adding New Benchmarks

When adding a new helper function to the library, add a corresponding benchmark in the appropriate `{package}_{category}_bench_test.go` file. Use the standard parametric pattern:

```go
func BenchmarkMyFunc(b *testing.B) {
    for _, n := range lengths {
        data := genSliceInt(n)
        b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                _ = lo.MyFunc(data, ...)
            }
        })
    }
}
```

Use shared generators from `helpers_test.go` — do not create local generator functions.
