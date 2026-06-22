---
name: Mean
slug: mean
sourceRef: math.go#L126
category: core
subCategory: math
playUrl: https://go.dev/play/p/tPURSuteUsP
variantHelpers:
  - core#math#mean
similarHelpers:
  - core#math#meanby
  - core#math#mode
  - core#math#sum
  - core#math#sumby
  - core#math#product
  - core#math#productby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
position: 80
signatures:
  - "func Mean[T constraints.Float | constraints.Integer](collection []T) T"
---

Calculates the arithmetic mean of a collection of numbers. Returns 0 for an empty collection.

```go
lo.Mean([]int{2, 3, 4, 5})
// 3
```


