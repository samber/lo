---
name: Subset
slug: subset
sourceRef: it/seq.go#L667
category: it
subCategory: sequence
signatures:
  - "func Subset[T any, I ~func(func(T) bool)](collection I, offset, length int) I"
variantHelpers:
  - it#sequence#subset
similarHelpers:
  - it#sequence#slice
  - it#sequence#drop
  - it#sequence#dropright
  - core#slice#slice
position: 120
---

Returns a subset of a sequence starting from the specified offset with the given length.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}
result := lo.Subset(seq, 1, 3)
// iter.Seq[int] yielding 2, 3, 4

result = lo.Subset(seq, 0, 2)
// iter.Seq[int] yielding 1, 2

result = lo.Subset(seq, 3, 10)
// iter.Seq[int] yielding 4, 5 (returns available elements)

result = lo.Subset(seq, 10, 5)
// iter.Seq[int] yielding nothing (offset beyond sequence)

seq = func(yield func(string) bool) {
    yield("a")
    yield("b")
    yield("c")
    yield("d")
}
result = lo.Subset(seq, 1, 2)
// iter.Seq[string] yielding "b", "c"
```