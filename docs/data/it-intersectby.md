---
name: IntersectBy
slug: intersectby
sourceRef: it/intersect.go#L78
category: it
subCategory: intersect
signatures:
  - "func IntersectBy[T any, K comparable, I ~func(func(T) bool)](func(T) K, lists ...I) I"
playUrl:
variantHelpers:
  - it#intersect#intersectby
similarHelpers:
  - it#intersect#intersect
  - core#slice#intersect
  - core#slice#intersectby
  - it#intersect#union
position: 10
---

Returns the intersection between given collections using a custom key selector function.

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

transform := func(v int) string {
  return strconv.Itoa(v)
}

intersection := it.IntersectBy(transform, seq1, seq2, seq3)

var result []int
for v := range intersection {
    result = append(result, v)
}
// result contains 2, 3 (elements present in all sequences)
```