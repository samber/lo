---
id: glossary
title: "Go Generics & Functional Programming Glossary"
description: "Definitions for the Go and functional-programming terms used throughout samber/lo's documentation — generics, iterators, predicates, and more."
sidebar_position: 100
---

# Go Generics & Functional Programming Glossary

## Generics
A Go 1.18+ language feature allowing functions and types to work with any type while
keeping compile-time type safety, via type parameters (`func Map[T, R any](...)`).
`lo` uses generics throughout instead of `interface{}` and reflection, so mismatched
types are caught by the compiler, not at runtime. See [Getting started](/docs/getting-started).

## Type Parameter
The `[T any]` part of a generic function's signature — a placeholder for a concrete
type, filled in either explicitly (`lo.Map[int, string](...)`) or, more commonly,
inferred by the compiler from the arguments.

## Type Constraint
A restriction on what types a type parameter can accept, such as `comparable` (usable
as a map key or with `==`) or `constraints.Ordered` (supports `<`, `>`). `lo.Max`
requires an ordered type; `lo.Uniq` requires a comparable one.

## `comparable`
A built-in Go type constraint meaning a type supports `==` and `!=`, and can be used
as a map key. Helpers like `lo.Contains`, `lo.Uniq`, and `lo.GroupBy` require their
element or key type to satisfy `comparable`.

## Reflection
A runtime technique (Go's `reflect` package) for examining and manipulating types and
values without knowing them at compile time. Libraries predating generics (like
`go-funk`) relied on reflection; `lo` avoids it, trading a small amount of flexibility
for compile-time safety and better performance. See
[samber/lo vs go-funk](/compare/samber-lo-vs-go-funk).

## Type Safety
Catching type errors at compile time instead of at runtime. `lo` achieves this through
generics rather than through runtime type assertions.

## Zero Value
The default value Go assigns to a variable that hasn't been explicitly initialized —
`0` for numbers, `""` for strings, `nil` for pointers/slices/maps/interfaces.
`lo.Compact` removes zero-value elements from a slice; `lo.IsEmpty` checks whether a
value equals its type's zero value.

## Nil Safety
Avoiding panics from dereferencing a `nil` pointer or operating on a `nil`
slice/map. `lo.FromPtrOr`, `lo.IsNil`/`lo.IsNotNil`, and `lo.Coalesce` exist
specifically to handle pointers and optional values without a `nil` check at every
call site.

## Immutability
Data that isn't changed after creation. Most core `lo` functions are immutable: they
return a new slice or map rather than modifying the input. `lo/mutable` provides
explicit in-place variants for when that allocation cost isn't acceptable.

## Predicate Function
A function returning a boolean, used to test or filter elements — the `fn` in
`lo.Filter(collection, fn)`, `lo.Find(collection, fn)`, and `lo.Every(collection, fn)`.

## Iteratee
A general term (borrowed from Lodash) for the callback function passed to a
collection helper — a predicate, transformer, or reducer, depending on the operation.
See [Lodash to Go](/guides/from-lodash) for the full vocabulary mapping.

## Transformer Function
A function that maps one value to another. The `fn` in `lo.Map(collection, fn)` and
`lo.MapValues(m, fn)`.

## Reducer Function
A function that combines an accumulator and an element into a new accumulator, used by
`lo.Reduce` and `lo.ReduceRight` to fold a collection down to a single value.

## Comparator Function
A function comparing two values to determine their order, used by `lo.MaxBy`,
`lo.MinBy`, and the sorting-adjacent helpers.

## Higher-Order Function
A function that takes another function as a parameter or returns one. Nearly every
`lo` helper is higher-order — that's what makes `Map`/`Filter`/`Reduce` reusable across
any element type and any transformation.

## Mutable Operations
Functions that modify a collection in place instead of allocating a new one —
`lo/mutable`'s `Filter`, `Map`, `Reverse`, `Shuffle`, `Fill`. Trades the safety of
never observing a partially-mutated value (relevant when a slice is shared across
goroutines) for avoiding an allocation. See [In-place slice mutation](/docs/mutable/slice).

## Memory Efficiency
How much a helper allocates relative to a hand-written equivalent. Verified
allocation-profile data for `lo`'s slice helpers is in
[Is samber/lo slow?](/compare/performance) — the short version: `lo.Map`/`lo.Filter`
allocate identically to a pre-sized, hand-written loop.

## Lazy Evaluation
Computing a value only when it's actually consumed, rather than eagerly up front.
`lo/it` implements every core helper lazily over Go 1.23+ iterators (`iter.Seq`), so
`it.Filter(it.Map(seq, f), p)` never allocates an intermediate slice — each element
flows through the whole pipeline before the next one starts. Contrast with **Eager
Evaluation** below.

## Eager Evaluation
Computing a value immediately, in full, before it's needed — what core `lo` does.
`lo.Filter` allocates and fully populates a result slice before returning it. Eager
evaluation is simpler to reason about; lazy evaluation (`lo/it`) avoids intermediate
allocations on long pipelines.

## `iter.Seq` / `iter.Seq2`
Go 1.23+ standard library types representing a lazy sequence: `iter.Seq[V]` is
`func(yield func(V) bool)`, and `iter.Seq2[K, V]` is the two-value form used for
key-value pairs. `lo/it` provides `Map`/`Filter`/`Reduce`-style operators over both.
See [Go iterator helpers](/docs/iter/slice).

## Range-over-func
The Go 1.23+ language feature (`for v := range seq`) that lets a function matching
`iter.Seq`'s shape be consumed with a normal `for range` loop — the mechanism
`iter.Seq` and `lo/it` are built on.

## Concurrency Safety
Whether an operation is safe to call from multiple goroutines simultaneously, and
what happens to shared data if it isn't. Core `lo` functions are safe to call
concurrently on independent inputs, since they don't mutate shared state; `lo/mutable`
functions mutate their input in place and are not safe to call concurrently on the
*same* slice.

## Worker Pool
A fixed or dynamic set of goroutines pulling work from a shared queue, used to bound
concurrency. `lo/parallel` implements a simpler variant — one goroutine per item —
appropriate when the transform is slow or I/O-bound; see
[Parallel slice processing](/docs/parallel/slice) and
[when to reach for it](/compare/samber-lo-vs-go-stdlib).

## Debounce
Delaying a function call until a burst of triggering events has settled — calling it
once after the last event, not once per event. `lo.NewDebounce` implements this; see
[Concurrency helpers](/docs/core/concurrency).

## Throttle
Limiting a function to run at most once per fixed time window, regardless of how many
times it's triggered. `lo.NewThrottle` implements this — the fan-in counterpart to
debounce's fan-out delay.

## Set Operations
Operations treating slices as mathematical sets: union, intersection, and difference.
`lo.Union`, `lo.Intersect`, and `lo.Difference` implement these over Go slices; see
[Set operations on slices](/docs/core/intersect).

## Chunking
Splitting a collection into fixed-size sub-collections. `lo.Chunk` returns
`[][]T`; the Go 1.23+ standard library's `slices.Chunk` returns a lazy
`iter.Seq[[]T]` instead — see the
[signature comparison](/compare/samber-lo-vs-go-stdlib) for the distinction.

## Tuple
A fixed-size, heterogeneous grouping of values — Go has no native tuple type, so `lo`
provides `Tuple2`..`Tuple9` generic structs, plus `Zip2`..`Zip9`/`Unzip2`..`Unzip9` to
convert between parallel slices and slices of tuples. See [Tuple helpers](/docs/core/tuple).

## Currying / Partial Application
Fixing some of a function's arguments ahead of time, producing a new function that
takes the rest. `lo.PartialX` implements partial application for a fixed set of
arities, since Go's generics can't express arbitrary arity the way a dynamically
typed language can.
