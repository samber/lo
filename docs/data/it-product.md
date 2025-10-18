---
name: Product / ProductBy
slug: product
sourceRef: it/math.go#L70
category: it
subCategory: math
signatures:
  - "func Product[T constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T]) T"
  - "func ProductBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T], iteratee func(item T) R) R"
playUrl: "https://go.dev/play/p/ebgxKxJmhLj"
variantHelpers:
  - it#math#product
  - it#math#productby
similarHelpers:
  - core#slice#product
  - core#slice#productby
position: 20
---

Multiplies values from a sequence. `ProductBy` applies a transform then multiplies. Returns 1 for empty sequences.

Examples:

```go
seq := it.RangeFrom(1, 4) // 1,2,3,4
p := it.Product(seq)
// p == 24
```

```go
nums := it.RangeFrom(2, 3) // 2,3,4
p := it.ProductBy(nums, func(n int) int { return n - 1 })
// (1*2*3) == 6
```


