---
name: MeanBy
slug: meanby
sourceRef: math.go#L137
category: core
subCategory: math
playUrl: https://go.dev/play/p/j7TsVwBOZ7P
variantHelpers:
  - core#math#meanby
similarHelpers:
  - core#math#mean
  - core#math#mode
  - core#math#sum
  - core#math#sumby
  - core#math#product
  - core#math#productby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
position: 90
signatures:
  - "func MeanBy[T any, R constraints.Float | constraints.Integer](collection []T, iteratee func(item T) R) R"
---

Calculates the mean of values computed by a predicate. Returns 0 for an empty collection.

```go
list := []string{"aa", "bbb", "cccc", "ddddd"}
lo.MeanBy(list, func(item string) float64 {
    return float64(len(item))
})
// 3.5
```


