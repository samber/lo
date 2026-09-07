---
title: "Go SIMD Slice Helpers (experimental)"
description: "8 experimental, high-performance slice operations using AVX/AVX2/AVX512 in Go 1.26+ (GOEXPERIMENT=simd, amd64 only)."
sidebar_position: 0
hide_table_of_contents: true
---

# SIMD - Experimental helpers

8 experimental, high-performance slice operations using AVX/AVX2/AVX512 in Go 1.26+ (GOEXPERIMENT=simd, amd64 only).

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
  category="experimental"
  subCategory="simd"
/>

---

Found a mistake or a missing detail? [Open an issue](https://github.com/samber/lo/issues) — feedback helps us improve this page.
