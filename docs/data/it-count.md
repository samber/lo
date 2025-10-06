---
name: Count / CountBy
slug: count
sourceRef: it/seq.go#L630
category: it
subCategory: sequence
signatures:
  - "func Count[T comparable](collection iter.Seq[T], value T) int"
  - "func CountBy[T any](collection iter.Seq[T], predicate func(item T) bool) int"
playUrl: ""
variantHelpers:
  - it#sequence#count
  - it#sequence#countby
similarHelpers:
  - core#slice#count
  - it#sequence#countvalues
position: 110
---

Counts elements in a collection. Count counts elements equal to a value, CountBy counts elements matching a predicate.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(2)
    _ = yield(3)
    _ = yield(2)
}
cnt := it.Count(seq, 2)
// cnt == 3
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("apricot")
}
cnt := it.CountBy(seq, func(s string) bool {
    return len(s) > 5
})
// cnt == 2 (banana, apricot)
```