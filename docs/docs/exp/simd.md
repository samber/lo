---
title: SIMD operations
description: High-performance slice operations using SSE, AVX2 and AVX512 SIMD when built with Go 1.26+ and GOEXPERIMENT=simd on amd64.
sidebar_position: 0
hide_table_of_contents: true
---

:::warning Help improve this documentation
This documentation is still new and evolving. If you spot any mistakes, unclear explanations, or missing details, please [open an issue](https://github.com/samber/lo/issues).

Your feedback helps us improve!
:::

#
## SIMD helpers

This page lists all operations on slices, available in the `exp/simd` sub-package. These helpers use **SSE** (128-bit), **AVX2** (256-bit) or **AVX512** (512-bit) SIMD when built with Go 1.26+, the `GOEXPERIMENT=simd` flag, and on amd64.

import HelperList from '@site/plugins/helpers-pages/components/HelperList';

<HelperList 
  category="exp"
  subCategory="simd"
/>
