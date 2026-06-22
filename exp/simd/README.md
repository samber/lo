# SIMD experiment (Go 1.26+)

This package requires **Go 1.26** with `GOEXPERIMENT=simd` and **amd64**.

See [benchmarks](./BENCHMARK.md).

## CPU compatibility (avoiding SIGILL)

If you see **SIGILL: illegal instruction** when running tests, the CPU or VM does not support the SIMD instructions used by that code.

### Check support on Linux

```bash
# List SIMD-related flags
grep -E 'avx' /proc/cpuinfo

# Or with lscpu
lscpu | grep -i avx
```

**Rough mapping:**

| Tests / code      | Required flag(s)           | Typical CPUs                                                            |
| ----------------- | -------------------------- | ----------------------------------------------------------------------- |
| AVX (128-bit)     | `avx` (baseline on amd64)  | All amd64                                                               |
| AVX2 (256-bit)    | `avx2`                     | Intel Haswell+, AMD Excavator+                                          |
| AVX-512 (512-bit) | `avx512f`                  | Intel Skylake-X+, some Xeons; many AMD/consumer CPUs do **not** have it |

### What the tests do

- **AVX tests** (128-bit) call `requireAVX(t)` and are **skipped** if the CPU does not support AVX.
- **AVX2 tests** call `requireAVX2(t)` and are **skipped** if the CPU does not support AVX2 (no SIGILL).
- **AVX-512 tests** (when enabled) should call `requireAVX512(t)` and skip when AVX-512 is not available.

So on a machine without AVX2, AVX2 tests will show as skipped instead of crashing.

### Run only AVX tests

If your environment does not support AVX2/AVX-512, you can still run the AVX (128-bit) tests:

```bash
GOEXPERIMENT=simd go test -run AVX ./...
```
