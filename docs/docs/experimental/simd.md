---
title: SIMD
description: High-performance slice operations using Go 1.27's portable simd package, built with GOEXPERIMENT=simd on amd64, arm64, and wasm.
sidebar_position: 0
hide_table_of_contents: true
---

# SIMD - Experimental helpers

This page lists all operations on slices, available in the `exp/simd` sub-package. These helpers use Go 1.27's portable `simd` package, built with the `GOEXPERIMENT=simd` flag. The same code compiles to AVX/AVX2/AVX512 on amd64, NEON on arm64, and SIMD128 on wasm, picking the widest vector the CPU supports at runtime, with an automatic fallback to the scalar `lo.*` implementation everywhere else (including `GODEBUG=simd=0`).

:::warning Help improve this documentation
This documentation is still new and evolving. If you spot any mistakes, unclear explanations, or missing details, please [open an issue](https://github.com/samber/lo/issues).

Your feedback helps us improve!
:::

:::warning Unstable API
SIMD helpers are experimental. The API may break in the future.
:::

## Performance

Benchmarks show that running SIMD operations on small datasets (below one vector's worth of
data) is slower, since the fixed vector setup cost dominates:

```txt
BenchmarkSumInt8/small/lo-2      344053036    3.576 ns/op
BenchmarkSumInt8/small/simd-2     48484386   24.80 ns/op
```

But much faster on large datasets, where the win scales with the CPU's vector width (up to
16 int32 lanes on AVX-512, 4 on NEON):

```txt
BenchmarkSumInt8/xlarge/lo-2         466400   2576 ns/op
BenchmarkSumInt8/xlarge/simd-2     13143464   92.25 ns/op
```

See [BENCHMARK.md](https://github.com/samber/lo/blob/master/exp/simd/BENCHMARK.md) for a full
amd64 (AVX-512) vs arm64 (NEON) comparison across every helper.

import HelperList from '@site/plugins/helpers-pages/components/HelperList';

<HelperList
  category="experimental"
  subCategory="simd"
/>
