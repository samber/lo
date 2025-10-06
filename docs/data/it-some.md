---
name: Some
slug: some
sourceRef: it/intersect.go#L52
category: it
subCategory: intersect
signatures:
  - "func Some[T comparable](collection iter.Seq[T], subset ...T) bool"
playUrl: ""
variantHelpers:
  - it#intersect#some
similarHelpers:
  - core#slice#some
  - it#intersect#every
  - it#intersect#none
position: 40
---

Returns true if at least one element of a subset is contained in a collection.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
hasAny := it.Some(seq, 2, 6)
// hasAny == true (2 is in collection)
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
hasAny := it.Some(seq, "orange", "grape")
// hasAny == false (neither is in collection)
```