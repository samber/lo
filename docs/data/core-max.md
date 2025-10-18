---
name: Max
slug: max
sourceRef: find.go#L410
category: core
subCategory: find
playUrl: https://go.dev/play/p/r6e-Z8JozS8
variantHelpers:
  - core#find#max
similarHelpers:
  - core#find#min
  - core#find#maxby
  - core#find#minby
  - core#find#maxindex
  - core#find#minindex
  - core#find#maxindexby
  - core#find#minindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 200
signatures:
  - "func Max[T constraints.Ordered](collection []T) T"
---

Searches the maximum value of a collection. Returns zero value when the collection is empty.

```go
max := lo.Max([]int{2, 5, 3})
// 5
```


