---
name: Contains
slug: contains
sourceRef: it/intersect.go#L13
category: iter
subCategory: intersect
signatures:
  - "func Contains[T comparable](collection iter.Seq[T], element T) bool"
playUrl: "https://go.dev/play/p/1edj7hH3TS2"
variantHelpers:
  - iter#intersect#contains
similarHelpers:
  - core#slice#contains
  - iter#intersect#containsby
  - iter#intersect#some
position: 0
---

Returns true if an element is present in a collection.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
}
has := it.Contains(seq, 20)
// has == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
has := it.Contains(seq, "orange")
// has == false
```