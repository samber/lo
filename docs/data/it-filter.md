---
name: Filter
slug: filter
sourceRef: it/seq.go#L33
category: it
subCategory: sequence
signatures:
  - "func Filter[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
playUrl: "https://go.dev/play/p/psenko2KKsX"
variantHelpers:
  - it#sequence#filter
  - it#sequence#filteri
similarHelpers:
  - core#slice#filter
  - it#sequence#reject
position: 10
---

Returns a sequence of all elements for which the predicate function returns true.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
filtered := it.Filter(seq, func(x int) bool {
    return x%2 == 0
})
var result []int
for v := range filtered {
    result = append(result, v)
}
// result contains 2, 4 (even numbers)
```