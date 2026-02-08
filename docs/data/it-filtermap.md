---
name: FilterMap
slug: filtermap
sourceRef: it/seq.go#L86
category: it
subCategory: sequence
signatures:
  - "func FilterMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R]"
  - "func FilterMapI[T, R any](collection iter.Seq[T], callback func(item T, index int) (R, bool)) iter.Seq[R]"
variantHelpers:
  - it#sequence#filtermap
  - it#sequence#filtermapi
similarHelpers:
  - it#sequence#map
  - it#sequence#filter
  - it#sequence#filtermaptoslice
  - it#sequence#rejectmap
position: 40
---

Maps elements of a sequence to new values and filters out elements where the callback returns false. Only elements where the second return value is true are included in the result.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}
result := it.FilterMap(seq, func(x int) (string, bool) {
    if x%2 == 0 {
        return fmt.Sprintf("even-%d", x), true
    }
    return "", false
})
// iter.Seq[string] yielding "even-2", "even-4"

seq = func(yield func(string) bool) {
    yield("a")
    yield("")
    yield("c")
    yield("d")
}
result = it.FilterMap(seq, func(s string) (int, bool) {
    if s != "" {
        return len(s), true
    }
    return 0, false
})
// iter.Seq[int] yielding 1, 1, 1 (length of "a", "c", "d")
```

### FilterMapI

Maps elements of a sequence to new values and filters out elements where the callback returns false. The callback receives both the item and its index.

```go
seq := func(yield func(string) bool) {
    yield("apple")
    yield("banana")
    yield("cherry")
}
result := it.FilterMapI(seq, func(s string, index int) (string, bool) {
    if index%2 == 0 {
        return fmt.Sprintf("%s-%d", s, index), true
    }
    return "", false
})
// iter.Seq[string] yielding "apple-0", "cherry-2"
```