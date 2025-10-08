---
name: Uniq
slug: uniq
sourceRef: it/seq.go#L216
category: it
subCategory: sequence
signatures:
  - "func Uniq[T comparable, I ~func(func(T) bool)](collection I) I"
playUrl: https://go.dev/play/p/0RlEI4-zq
variantHelpers:
  - it#sequence#uniq
similarHelpers:
  - core#slice#uniq
  - it#sequence#uniqby
position: 50
---

Returns a duplicate-free version of a sequence, removing consecutive duplicates.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(2)
    _ = yield(3)
    _ = yield(2)
    _ = yield(2)
}
uniqueSeq := it.Uniq(seq)
var result []int
for v := range uniqueSeq {
    result = append(result, v)
}
// result contains 1, 2, 3, 2 (consecutive duplicates removed)
```