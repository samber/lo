---
name: Without
slug: without
sourceRef: it/intersect.go#L155
category: it
subCategory: intersect
signatures:
  - "func Without[T comparable, I ~func(func(T) bool)](collection I, exclude ...T) I"
playUrl: "https://go.dev/play/p/eAOoUsQnrZf"
variantHelpers:
  - it#intersect#without
similarHelpers:
  - core#slice#without
  - it#intersect#intersect
position: 50
---

Returns a sequence excluding all given values.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
filtered := it.Without(seq, 2, 4)
var result []int
for v := range filtered {
    result = append(result, v)
}
// result contains 1, 3, 5
```