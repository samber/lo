---
name: FindLastIndexOf
slug: findlastindexof
sourceRef: it/find.go#L127
category: it
subCategory: find
signatures:
  - "func FindLastIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool)"
playUrl: ""
variantHelpers:
  - it#find#findlastindexof
similarHelpers:
  - core#slice#findlastindexof
  - it#find#findindexof
  - it#find#find
position: 60
---

Searches for the last element matching a predicate and returns the element, its index, and true if found. Returns zero value, -1, and false if not found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(20)
    _ = yield(40)
}
found, index, ok := it.FindLastIndexOf(seq, func(x int) bool {
    return x == 20
})
// found == 20, index == 3, ok == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
found, index, ok := it.FindLastIndexOf(seq, func(s string) bool {
    return len(s) > 10
})
// found == "", index == -1, ok == false
```