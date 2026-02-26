---
title: SIMD operations
description: High-performance slice operations using AVX, AVX2 and AVX512 SIMD when built with Go 1.26+ and GOEXPERIMENT=simd on amd64.
sidebar_position: 0
hide_table_of_contents: true
---

:::warning Help improve this documentation
This documentation is still new and evolving. If you spot any mistakes, unclear explanations, or missing details, please [open an issue](https://github.com/samber/lo/issues).

Your feedback helps us improve!
:::

#
## SIMD helpers

This page lists all operations on slices, available in the `exp/simd` sub-package. These helpers use **AVX** (128-bit), **AVX2** (256-bit) or **AVX512** (512-bit) SIMD when built with Go 1.26+, the `GOEXPERIMENT=simd` flag, and on amd64.

:::warning Unstable API
SIMD helpers are experimental. The API may break in the future.
:::

## Performance

Benchmarks show that running SIMD operators on small datasets is slower:

```txt
BenchmarkSumInt8/small/Fallback-lo-4             203616572        5.875 ns/op
BenchmarkSumInt8/small/AVX-x16-4                 100000000        12.04 ns/op
BenchmarkSumInt8/small/AVX2-x32-4                 64041816        17.93 ns/op
BenchmarkSumInt8/small/AVX512-x64-4               26947528        44.75 ns/op
```

But much much faster on big datasets:

```txt
BenchmarkSumInt8/xlarge/Fallback-lo-4               247677       4860 ns/op
BenchmarkSumInt8/xlarge/AVX-x16-4                  3851040      311.4 ns/op
BenchmarkSumInt8/xlarge/AVX2-x32-4                 7100002      169.2 ns/op
BenchmarkSumInt8/xlarge/AVX512-x64-4              10107534      118.1 ns/op
```

import HelperList from '@site/plugins/helpers-pages/components/HelperList';

<HelperList
  category="exp"
  subCategory="simd"
/>
