---
name: MaxBy
slug: maxby
sourceRef: find.go#L459
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#maxby
similarHelpers:
  - core#find#max
  - core#find#maxindex
  - core#find#maxindexby
  - core#find#min
  - core#find#minby
  - core#find#minindex
  - core#find#minindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 220
signatures:
  - "func MaxBy[T any](collection []T, comparison func(a T, b T) bool) T"
---

Searches the maximum value of a collection using the given comparison function. Returns zero value when the collection is empty.

```go
type Point struct{ X int }
max := lo.MaxBy([]Point{{1}, {5}, {3}}, func(a, b Point) bool {
    return a.X > b.X
})
// {5}
```


