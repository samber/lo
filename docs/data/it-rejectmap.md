---
name: RejectMap
slug: rejectmap
sourceRef: it/seq.go#L608
category: it
subCategory: sequence
signatures:
  - "func RejectMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R]"
variantHelpers:
  - it#sequence#rejectmap
similarHelpers:
  - it#sequence#filtermap
  - it#sequence#map
  - it#sequence#filter
  - it#sequence#reject
position: 42
---

Maps elements of a sequence to new values and rejects elements where the callback returns true. Only elements where the second return value is false are included in the result.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}
result := lo.RejectMap(seq, func(x int) (string, bool) {
    if x%2 == 0 {
        return fmt.Sprintf("even-%d", x), true // reject even numbers
    }
    return fmt.Sprintf("odd-%d", x), false
})
// iter.Seq[string] yielding "odd-1", "odd-3"

seq = func(yield func(string) bool) {
    yield("a")
    yield("")
    yield("c")
    yield("d")
}
result = lo.RejectMap(seq, func(s string) (int, bool) {
    if s == "" {
        return 0, true // reject empty strings
    }
    return len(s), false
})
// iter.Seq[int] yielding 1, 1, 1 (length of "a", "c", "d")
```