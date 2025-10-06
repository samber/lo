---
name: ProductBy
slug: productby
sourceRef: math.go#L108
category: core
subCategory: math
playUrl: https://go.dev/play/p/wadzrWr9Aer
variantHelpers:
  - core#math#productby
similarHelpers:
  - core#math#product
  - core#math#sumby
  - core#math#meanby
  - core#find#minby
  - core#find#maxby
  - core#find#minindexby
  - core#find#maxindexby
  - core#math#sum
  - core#math#mean
position: 70
signatures:
  - "func ProductBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R"
---

Calculates the product of values computed by a predicate. Returns 1 for nil or empty collections.

```go
strings := []string{"foo", "bar"}
lo.ProductBy(strings, func(item string) int {
    return len(item)
})
// 9
```


