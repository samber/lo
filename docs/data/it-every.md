---
name: Every
slug: every
sourceRef: it/intersect.go#L25
category: it
subCategory: intersect
signatures:
  - "func Every[T comparable](collection iter.Seq[T], subset ...T) bool"
playUrl: ""
variantHelpers:
  - it#intersect#every
similarHelpers:
  - core#slice#every
  - it#intersect#some
  - it#intersect#none
position: 30
---

Returns true if all elements of a subset are contained in a collection.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
hasAll := it.Every(seq, 2, 4)
// hasAll == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
hasAll := it.Every(seq, "apple", "orange")
// hasAll == false (orange is not in collection)
```