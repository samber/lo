---
name: MinIndexBy
slug: minindexby
sourceRef: find.go#L337
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#minindexby
similarHelpers:
  - core#find#min
  - core#find#minby
  - core#find#minindex
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#maxindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 170
signatures:
  - "func MinIndexBy[T any](collection []T, comparison func(a T, b T) bool) (T, int)"
---

Searches the minimum value using a comparison function and returns the value and its index. Returns (zero value, -1) when empty.

```go
type Point struct{ X int }
value, idx := lo.MinIndexBy([]Point{{1}, {5}, {3}}, func(a, b Point) bool {
    return a.X < b.X
})
// value == {1}, idx == 0
```


