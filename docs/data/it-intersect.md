---
name: Intersect
slug: intersect
sourceRef: it/intersect.go#L78
category: it
subCategory: intersect
signatures:
  - "func Intersect[T comparable, I ~func(func(T) bool)](lists ...I) I"
playUrl: "https://go.dev/play/p/kz3cGhGZZWF"
variantHelpers:
  - it#intersect#intersect
similarHelpers:
  - core#slice#intersect
  - it#intersect#union
position: 10
---

Returns the intersection between given collections (elements present in all collections).

Examples:

```go
seq1 := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
seq2 := func(yield func(int) bool) {
    _ = yield(2)
    _ = yield(3)
    _ = yield(5)
}
seq3 := func(yield func(int) bool) {
    _ = yield(3)
    _ = yield(2)
    _ = yield(6)
}
intersection := it.Intersect(seq1, seq2, seq3)
var result []int
for v := range intersection {
    result = append(result, v)
}
// result contains 2, 3 (elements present in all sequences)
```