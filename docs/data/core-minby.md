---
name: MinBy
slug: minby
sourceRef: find.go#L314
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#minby
similarHelpers:
  - core#find#min
  - core#find#minindex
  - core#find#minindexby
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#maxindexby
  - core#math#sum
  - core#math#mean
  - core#math#product
  - core#math#mode
position: 160
signatures:
  - "func MinBy[T any](collection []T, comparison func(a T, b T) bool) T"
---

Searches the minimum value of a collection using the given comparison function. Returns the first minimal value; zero value when empty.

```go
type Point struct{ X int }
min := lo.MinBy([]Point{{1}, {5}, {3}}, func(a, b Point) bool {
    return a.X < b.X
})
// {1}
```


