---
name: IsEmpty
slug: isempty
sourceRef: it/type_manipulation.go#L50
category: iter
subCategory: type
signatures:
  - "func IsEmpty[T any](collection iter.Seq[T]) bool"
playUrl: https://go.dev/play/p/krZ-laaVi2C
variantHelpers:
  - iter#type#isempty
similarHelpers:
  - iter#type#isnotempty
  - iter#type#empty
  - iter#sequence#length
position: 10
---

Returns true if the sequence is empty. Will consume the entire sequence to check.

Examples:

```go
seq := func(yield func(int) bool) {
    // empty sequence
}
empty := it.IsEmpty(seq)
// empty == true
```

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
}
empty := it.IsEmpty(seq)
// empty == false
```