---
name: MinIndex
slug: minindex
sourceRef: find.go#L287
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#minindex
similarHelpers:
  - core#find#min
  - core#find#minby
  - core#find#minindexby
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#maxindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 150
signatures:
  - "func MinIndex[T constraints.Ordered](collection []T) (T, int)"
---

Returns the minimum value and its index. Returns (zero value, -1) when the collection is empty.

```go
value, idx := lo.MinIndex([]int{2, 5, 3})
// value == 2, idx == 0
```


