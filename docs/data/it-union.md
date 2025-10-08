---
name: Union
slug: union
sourceRef: it/intersect.go#L136
category: it
subCategory: intersect
signatures:
  - "func Union[T comparable, I ~func(func(T) bool)](lists ...I) I"
playUrl: "https://go.dev/play/p/ImIoFNpSUUB"
variantHelpers:
  - it#intersect#union
similarHelpers:
  - core#slice#union
  - it#intersect#intersect
position: 20
---

Returns all distinct elements from given collections (union of all collections).

Examples:

```go
seq1 := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
seq2 := func(yield func(int) bool) {
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
seq3 := func(yield func(int) bool) {
    _ = yield(3)
    _ = yield(5)
}
union := it.Union(seq1, seq2, seq3)
var result []int
for v := range union {
    result = append(result, v)
}
// result contains 1, 2, 3, 4, 5 (all distinct elements)
```