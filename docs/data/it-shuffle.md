---
name: Shuffle
slug: shuffle
sourceRef: it/seq.go#L357
category: it
subCategory: sequence
signatures:
  - "func Shuffle[T any, I ~func(func(T) bool)](collection I) I"
playUrl: ""
variantHelpers:
  - it#sequence#shuffle
similarHelpers:
  - core#slice#shuffle
  - it#sequence#reverse
position: 130
---

Returns a sequence of shuffled values using Fisher-Yates algorithm. Note: this requires collecting all elements in memory.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
shuffled := it.Shuffle(seq)
var result []int
for v := range shuffled {
    result = append(result, v)
}
// result contains the same elements in random order
```