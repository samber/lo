---
name: Splice
slug: splice
sourceRef: it/seq.go#L744
category: it
subCategory: sequence
signatures:
  - "func Splice[T any, I ~func(func(T) bool)](collection I, index int, elements ...T) I"
variantHelpers:
  - it#sequence#splice
similarHelpers:
  - it#sequence#slice
  - it#sequence#replace
  - it#sequence#replaceall
  - core#slice#splice
position: 122
---

Inserts elements into a sequence at the specified index. Returns a new sequence with the elements inserted.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(5)
}
result := it.Splice(seq, 2, 3, 4)
// iter.Seq[int] yielding 1, 2, 3, 4, 5

result = it.Splice(seq, 0, 0)
// iter.Seq[int] yielding 0, 1, 2, 5 (insert at beginning)

result = it.Splice(seq, 3, 6, 7)
// iter.Seq[int] yielding 1, 2, 5, 6, 7 (insert at end)

seq = func(yield func(string) bool) {
    yield("a")
    yield("c")
}
result = it.Splice(seq, 1, "b")
// iter.Seq[string] yielding "a", "b", "c"

result = it.Splice(seq, 1, "x", "y")
// iter.Seq[string] yielding "a", "x", "y", "c"
```
