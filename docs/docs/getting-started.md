---
title: 🚀 Getting started
description: Let's discover samber/lo in less than 5 minutes.
sidebar_position: 1
---

# Getting started

**samber/lo** is a Lodash-style utility library for Go 1.18+ that leverages generics to provide type-safe helper functions. The library is organized into several packages, each serving different use cases.

## Installation

```bash
go get -u github.com/samber/lo
```

## Base Package (`lo`)

The main package provides immutable utility functions for slices, maps, strings, math operations, and more. It's the core of the library with over 300+ functions.

```go
import "github.com/samber/lo"

// Example: Map a slice of numbers to their squares
numbers := []int{1, 2, 3, 4, 5}
squared := lo.Map(numbers, func(x int, _ int) int {
    return x * x
})
// Result: [1, 4, 9, 16, 25]
```

## Iter Package (`lo/it`)

The  `it` package provides Go 1.23+ sequence helpers with lazy evaluation, offering over 100 functions for efficient iteration without buffering.

```go
// Future usage (Go 1.23+)
import (
    "iter"
    loit "github.com/samber/lo/it"
)

// Lazy iteration without buffering
seq := loit.Filter(iter.Range(0, 1000), func(x int) bool {
    return x%2 == 0
})
```

## Mutable Package (`lo/mutable`)

The mutable package provides in-place operations that modify collections directly, useful for performance-critical scenarios.

```go
import lom "github.com/samber/lo/mutable"

// Filter in-place (modifies the original slice)
numbers := []int{1, 2, 3, 4, 5}
lom.Filter(&numbers, func(x int) bool {
    return x%2 == 0
})
// Result: [2, 4]
```

## Parallel Package (`lo/parallel`)

The parallel package enables concurrent processing of collections with built-in worker pools, perfect for CPU-intensive operations.

```go
import lop "github.com/samber/lo/parallel"

// Process items concurrently (4 workers by default)
results := lop.Map(numbers, 4, func(x int) int {
    // Some expensive operation
    return expensiveOperation(x)
})
```

## Key Benefits

- **Type-safe** with generics
- **Immutable** by default (main package)
- **Performance** optimized with parallel and mutable variants
- **Comprehensive** with 500+ utility functions
- **Lazy evaluation** with `iter` std package (Go >= 1.23)

## Next Steps

- Check the [Go documentation](https://pkg.go.dev/github.com/samber/lo) for complete API reference
- Explore examples in the repository
- Choose the right sub-package for your use case
