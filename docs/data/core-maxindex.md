---
name: MaxIndex
slug: maxindex
sourceRef: find.go#L432
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#maxindex
similarHelpers:
  - core#find#max
  - core#find#maxby
  - core#find#maxindexby
  - core#find#min
  - core#find#minby
  - core#find#minindex
  - core#find#minindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 210
signatures:
  - "func MaxIndex[T constraints.Ordered](collection []T) (T, int)"
---

Returns the maximum value and its index. Returns (zero value, -1) when the collection is empty.

```go
value, idx := lo.MaxIndex([]int{2, 5, 3})
// value == 5, idx == 1
```


