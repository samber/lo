---
name: Reduce
slug: reduce
sourceRef: slice.go#L87
category: core
subCategory: slice
playUrl: https://go.dev/play/p/CgHYNUpOd1I
variantHelpers:
  - core#slice#reduce
similarHelpers:
  - core#slice#reduceright
  - core#slice#sum
  - core#slice#sumby
  - core#slice#product
  - core#slice#productby
  - core#slice#mean
  - core#slice#meanby
position: 50
signatures:
  - "func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R"
---

Reduces a collection to a single value by accumulating results of an accumulator function. Each call receives the previous result value.

```go
sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
    return agg + item
}, 0)
// 10
```


