---
name: UnionBy
slug: unionby
sourceRef: intersect.go#L260
category: core
subCategory: intersect
playUrl:
variantHelpers:
  - core#intersect#unionby
similarHelpers:
  - core#intersect#union
  - core#intersect#intersect
  - core#intersect#intersectby
  - core#slice#uniq
  - core#slice#uniqby
position: 110
signatures:
  - "func UnionBy[T any, V comparable, Slice ~[]T](iteratee func(item T) V, lists ...Slice) Slice"
---

Returns all distinct elements from multiple collections based on a key function. The result maintains the relative order of first occurrences.

```go
lo.UnionBy(func(i int) int { return i / 2 }, []int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
// []int{0, 2, 4, 10}
```

```go
lo.UnionBy(func(s string) string { return s[:1] }, []string{"foo", "bar"}, []string{"baz"})
// []string{"foo", "baz"}
```
