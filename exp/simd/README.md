# SIMD experiment (Go 1.27+)

This package requires **Go 1.27** with `GOEXPERIMENT=simd`. Unlike `simd/archsimd`, the
stdlib `simd` package it is built on is portable: no architecture build tag is required.

See [benchmarks](./BENCHMARK.md).

## Platform support

| Architecture     | Backing hardware                       |
| ---------------- | --------------------------------------- |
| amd64            | AVX / AVX2 / AVX-512, picked at runtime |
| arm64            | NEON                                    |
| wasm             | SIMD128                                 |
| everything else  | pure-Go emulation                       |

On the emulated tier, this package transparently falls back to the equivalent scalar
`github.com/samber/lo` function instead of paying the emulation overhead.

## Controlling the vector width

`GODEBUG=simd=<0|128|256|512>` forces a specific behaviour for the current process:

```bash
GODEBUG=simd=0   go test ./...  # force the scalar/lo.* fallback everywhere
GODEBUG=simd=128 go test ./...  # force 128-bit vectors
GODEBUG=simd=256 go test ./...  # force 256-bit vectors (panics if unsupported by the CPU)
```

Without `GODEBUG=simd`, the widest vector the CPU supports (with the required feature set)
is used automatically. Building without `GOEXPERIMENT=simd` at all compiles this package to
an empty stub with no exported symbols, same as before this package existed.

## Running the tests

```bash
GOEXPERIMENT=simd go test ./...
```
