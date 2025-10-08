---
name: Length
slug: length
sourceRef: it/seq.go#L16
category: it
subCategory: sequence
signatures:
  - "func Length[T any](collection iter.Seq[T]) int"
playUrl: "https://go.dev/play/p/3dnbOjTbL-o"
variantHelpers:
  - it#sequence#length
similarHelpers:
  - core#slice#length
  - it#sequence#isempty
  - it#sequence#isnotempty
position: 0
---

Returns the length of a collection by iterating through the entire sequence.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
length := it.Length(seq)
// length == 3
```

```go
seq := func(yield func(string) bool) {
    // empty sequence
}
length := it.Length(seq)
// length == 0
```