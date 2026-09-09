// Package simd accelerates a subset of the samber/lo slice helpers (Sum, Mean, Min,
// Max, Clamp, SumBy, MeanBy, Contains) using the Go 1.27 stdlib "simd" package.
//
// Unlike simd/archsimd, the stdlib "simd" package is portable: it maps to AVX/AVX2/AVX512
// on amd64, NEON on arm64, SIMD128 on wasm, and falls back to a pure-Go emulation on every
// other architecture. This package therefore requires no per-architecture build tags.
//
// Building this package requires Go 1.27+ and GOEXPERIMENT=simd. Without the experiment
// enabled, this file is the only one compiled: the package exposes no symbols.
//
// # Fallback
//
// When the current process runs the emulated implementation (no hardware SIMD support, or
// GODEBUG=simd=0), every helper transparently falls back to the equivalent scalar
// github.com/samber/lo function, which is faster than paying the emulation overhead.
package simd
