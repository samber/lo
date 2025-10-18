---
name: DropLast
slug: droplast
sourceRef: it/seq.go#L512
category: it
subCategory: sequence
signatures:
  - "func DropLast[T any, I ~func(func(T) bool)](collection I, n int) I"
variantHelpers:
  - it#sequence#droplast
similarHelpers:
  - it#sequence#drop
  - it#sequence#dropwhile
  - it#sequence#droplastwhile
  - it#sequence#trim
  - it#sequence#trimsuffix
position: 78
---

Drops the last n elements from a sequence. Returns a new sequence without the specified number of trailing elements.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}
result := lo.DropLast(seq, 2)
// iter.Seq[int] yielding 1, 2, 3

result = lo.DropLast(seq, 0)
// iter.Seq[int] yielding 1, 2, 3, 4, 5 (unchanged)

result = lo.DropLast(seq, 10)
// iter.Seq[int] yielding nothing (all elements dropped)

seq = func(yield func(string) bool) {
    yield("a")
    yield("b")
    yield("c")
}
result = lo.DropLast(seq, 1)
// iter.Seq[string] yielding "a", "b"
```