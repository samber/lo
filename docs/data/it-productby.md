---
name: ProductBy
slug: productby
sourceRef: it/math.go#L90
category: it
subCategory: math
signatures:
  - "func ProductBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T], transform func(item T) R) R"
playUrl:
variantHelpers:
  - it#math#product
similarHelpers:
  - core#slice#productby
  - core#slice#product
position: 67
---

Returns the product of values in the collection using the given transform function.

```go
result := it.ProductBy(it.Range(1, 5), func(item int) int {
    return item * 2
})
// 384 (2 * 4 * 6 * 8)
```