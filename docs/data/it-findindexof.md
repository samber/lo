---
name: FindIndexOf
slug: findindexof
sourceRef: it/find.go#L112
category: it
subCategory: find
signatures:
  - "func FindIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool)"
playUrl: ""
variantHelpers:
  - it#find#findindexof
similarHelpers:
  - core#slice#findindexof
  - it#find#find
  - it#find#findlastindexof
position: 50
---

Searches for an element based on a predicate and returns the element, its index, and true if found. Returns zero value, -1, and false if not found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(40)
}
found, index, ok := it.FindIndexOf(seq, func(x int) bool {
    return x > 25
})
// found == 30, index == 2, ok == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
found, index, ok := it.FindIndexOf(seq, func(s string) bool {
    return s == "orange"
})
// found == "", index == -1, ok == false
```