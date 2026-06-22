---
name: Sum
slug: sum
sourceRef: math.go#L70
category: core
subCategory: math
playUrl: https://go.dev/play/p/upfeJVqs4Bt
variantHelpers:
  - core#math#sum
similarHelpers:
  - core#math#product
  - core#math#mean
  - core#math#sumby
  - core#math#productby
  - core#math#meanby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
  - core#math#mode
position: 40
signatures:
  - "func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T"
---

Sums the values in a collection. Returns 0 for an empty collection.

```go
lo.Sum([]int{1, 2, 3, 4, 5})
// 15
```


