---
title: 🤝 Contributing
description: Join the community of contributors.
sidebar_position: 110
---

# 🤝 Contributing

Hey! We are happy to have you as a new contributor. ✌️

For your first contribution please follow some guidelines:

## Function Naming
Function naming: helpers must be self-explanatory and respect standards (other languages, libraries...). Feel free to suggest many names in your contributions.

## Variadic functions
Many functions accept variadic parameters (like `lo.Keys(...map[K]V)` accepting multiple maps), providing flexibility while maintaining type safety.

## Slice type Parameters
Functions use `~[]T` constraints to accept any slice type, including named slice types, not just `[]T`. This design choice makes the library more flexible in real-world usage.

## Variants
When applicable, some functions can be added to sub-package as well: `mutable`, `it` and `parallel`. Add a documentation for each helper.

## Testing
We try to maintain code coverage above 90%.

## Benchmark and performance
Write performant helpers and limit extra memory consumption. Build an helper for general purpose and don't optimize for a particular use-case.

Feel free to write benchmarks.

Iterators can be unbounded and run for a very long time. If you expect a big memory footprint, please warn developers in the function comment.

## Documentation
Functions must be properly commented, with a Go Playground link. New helpers must be created with a markdown documentation in `docs/data/`. In markdown header, please link to similar helpers (and update other markdowns accordingly).

Add your helper to `docs/static/llms.txt`.

## Examples
Every function includes a "Play" link to the Go Playground, allowing developers to quickly experiment and understand behavior without setting up a local environment.

Please add an example of your helper in the file named `xxxx_example_test.go`. It will be available from Godoc website: https://pkg.go.dev/github.com/samber/lo
