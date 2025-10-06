---
name: MaxIndexBy
slug: maxindexby
sourceRef: find.go#L482
category: core
subCategory: find
playUrl:
variantHelpers:
  - core#find#maxindexby
similarHelpers:
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#min
  - core#find#minby
  - core#find#minindex
  - core#find#minindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 230
signatures:
  - "func MaxIndexBy[T any](collection []T, comparison func(a T, b T) bool) (T, int)"
---

Returns the maximum value and its index using the given comparison function. Returns (zero value, -1) when the collection is empty.

```go
type Point struct{ X int }
value, idx := lo.MaxIndexBy([]Point{{1}, {5}, {3}}, func(a, b Point) bool {
    return a.X > b.X
})
// value == {5}, idx == 1
```


