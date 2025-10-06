---
name: Last
slug: last
sourceRef: it/find.go#L389
category: it
subCategory: find
signatures:
  - "func Last[T any](collection iter.Seq[T]) (T, bool)"
playUrl: ""
variantHelpers:
  - it#find#last
similarHelpers:
  - core#slice#last
  - it#find#first
  - it#find#lastor
position: 130
---

Returns the last element of a collection and a boolean indicating availability. Returns zero value and false if the collection is empty.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
}
last, ok := it.Last(seq)
// last == 30, ok == true
```

```go
seq := func(yield func(string) bool) {
    // empty sequence
}
last, ok := it.Last(seq)
// last == "", ok == false (zero value for string)
```