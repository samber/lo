---
title: 🚀 Getting started
description: Let's discover samber/lo in less than 5 minutes.
sidebar_position: 1
---

# Getting started

**samber/lo** is a Lodash-style utility library for Go 1.18+ that leverages generics to provide type-safe helper functions. The library is organized into several packages, each serving different use cases.

## 🚀 Install

```bash
go get -u github.com/samber/lo@v1
```

**AI Agent Skill:**

```bash
npx skills add https://github.com/samber/cc-skills-golang --skill golang-samber-lo
```

## 🧢 Core Package (`lo`)

The main package provides immutable utility functions for slices, maps, strings, math operations, and more. It's the core of the library with 274 documented functions.

```go
import "github.com/samber/lo"

// Example: Map a slice of numbers to their squares
numbers := []int{1, 2, 3, 4, 5}
squared := lo.Map(numbers, func(x int, _ int) int {
    return x * x
})
// Result: [1, 4, 9, 16, 25]
```

## 🔄 Iter Package (`lo/it`)

The `it` package provides Go 1.23+ sequence helpers (`iter.Seq`) with lazy evaluation, offering over 100 functions for efficient iteration without buffering.

```go
import (
    "slices"
    loi "github.com/samber/lo/it"
)

seqIn := loi.Range(1000) // iter.Seq[int] of 0..999

// Lazy iteration without buffering
seqOut := loi.Filter(seqIn, func(x int) bool {
    return x%2 == 0
})

result := slices.Collect(seqOut)
// Result: [0, 2, 4, 6, ...]
```

## 👣 Mutable Package (`lo/mutable`)

The mutable package provides in-place operations that modify collections directly, useful for performance-critical scenarios.

```go
import lom "github.com/samber/lo/mutable"

// Filter in-place (modifies the original slice, returns the shortened view)
numbers := []int{1, 2, 3, 4, 5}
kept := lom.Filter(numbers, func(x int) bool {
    return x%2 == 0
})
// kept: [2, 4], backed by the same array as numbers
```

## 🏎️ Parallel Package (`lo/parallel`)

The parallel package enables concurrent processing of collections, transforming each item in its own goroutine and collecting results in the original order.

```go
import lop "github.com/samber/lo/parallel"

numbers := []int{1, 2, 3, 4, 5}

// Process items concurrently, one goroutine per item
results := lop.Map(numbers, func(x int, index int) int {
    // Some expensive operation
    return x * x
})
```

## ✅ Key Benefits

- **Type-safe** with generics
- **Immutable** by default (main package)
- **Performance** optimized with parallel and mutable variants
- **Comprehensive** with 449 documented utility functions
- **Lazy evaluation** with `iter` std package (Go >= 1.23)
- **Minimal dependencies** zero dependencies outside the Go standard library

## 👀 Next Steps

- Check the [Go documentation](https://pkg.go.dev/github.com/samber/lo) for complete API reference
- Explore examples in the repository
- Choose the right sub-package for your use case
