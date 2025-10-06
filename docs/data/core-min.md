---
name: Min
slug: min
sourceRef: find.go#L265
category: core
subCategory: find
playUrl: https://go.dev/play/p/r6e-Z8JozS8
variantHelpers:
  - core#find#min
similarHelpers:
  - core#find#minby
  - core#find#minindex
  - core#find#minindexby
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#maxindexby
position: 140
signatures:
  - "func Min[T constraints.Ordered](collection []T) T"
---

Returns the minimum value of a collection. Returns the zero value when the collection is empty.

```go
lo.Min([]int{1, 2, 3})
// 1
```


