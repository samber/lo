---
name: Product
slug: product
sourceRef: math.go#L90
category: core
subCategory: math
playUrl: https://go.dev/play/p/2_kjM_smtAH
variantHelpers:
  - core#math#product
similarHelpers:
  - core#math#sum
  - core#math#mean
  - core#math#productby
  - core#math#sumby
  - core#math#meanby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
  - core#math#mode
position: 60
signatures:
  - "func Product[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T"
---

Calculates the product of the values in a collection. Returns 1 for nil or empty collections.

```go
lo.Product([]int{1, 2, 3, 4, 5})
// 120
```


