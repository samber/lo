---
id: glossary
title: 📚 Glossary
description: Comprehensive glossary of samber/lo library terms and concepts
sidebar_position: 100
---

# Glossary

## Generics
Go 1.18+ feature that allows writing functions and types that work with any type while maintaining type safety. The `lo` library extensively uses generics to provide type-safe utility functions.

## Reflection
Language reflection is a programming technique where a program can examine and manipulate its own structure and behavior at runtime.

## Type Safety
The principle of ensuring that operations are performed on the correct data types, preventing runtime errors. `lo` leverages Go's type system and generics to provide type-safe operations and type-safe APIs.

## Immutability
The concept of data that cannot be changed after creation. Most `lo` functions follow immutable patterns, returning new collections rather than modifying existing ones.

## Predicate Function
A function that returns a boolean value, typically used for filtering or testing conditions. In `lo`, predicates are commonly used with `Filter`, `Find`, `Contains`, etc.

## Transformer Function
A function that transforms one value into another. Used with `Map`, `MapValues`, and other transformation operations.

## Reducer Function
A function that combines two values into one, used with `Reduce` and similar operations.

## Comparator Function
A function that compares two values and returns their relative order. Used with sorting operations.

## Higher-Order Functions
Functions that take other functions as parameters or return functions as results. Most `lo` utilities are higher-order functions.

## Mutable Operations
Functions that modify collections in-place rather than creating new ones. Mutable operations can be more memory efficient but less safe in concurrent scenarios and when data sharing is involved.

## Lazy Evaluation
Some `lo` operations implement lazy evaluation patterns for improved performance with large datasets.

## Memory Efficiency
Considerations for memory usage when choosing between immutable and mutable operations. Mutable operations can be more memory efficient but less safe in concurrent scenarios and when data sharing is involved.

## Concurrency Safety
How different operations behave in concurrent contexts and thread safety guarantees.
