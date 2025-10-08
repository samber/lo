---
name: First
slug: first
sourceRef: it/find.go#L362
category: it
subCategory: find
signatures:
  - "func First[T any](collection iter.Seq[T]) (T, bool)"
playUrl: "https://go.dev/play/p/5MgZD9-zl"
variantHelpers:
  - it#find#first
similarHelpers:
  - core#slice#first
  - it#find#last
  - it#find#firstor
position: 120
---

Returns the first element of a collection and a boolean indicating availability. Returns zero value and false if the collection is empty.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
}
first, ok := it.First(seq)
// first == 10, ok == true
```

```go
seq := func(yield func(string) bool) {
    // empty sequence
}
first, ok := it.First(seq)
// first == "", ok == false (zero value for string)
```