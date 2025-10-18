---
name: Find
slug: find
sourceRef: it/find.go#L105
category: it
subCategory: find
signatures:
  - "func Find[T any](collection iter.Seq[T], predicate func(item T) bool) (T, bool)"
playUrl: "https://go.dev/play/p/5SdLM6jf-q"
variantHelpers:
  - it#find#find
similarHelpers:
  - core#slice#find
  - it#find#findindexof
  - it#find#findorelse
position: 40
---

Searches for an element in a sequence based on a predicate function. Returns the element and true if found, zero value and false if not found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(40)
}
found, ok := it.Find(seq, func(x int) bool {
    return x > 25
})
// found == 30, ok == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
found, ok := it.Find(seq, func(s string) bool {
    return len(s) > 10
})
// found == "", ok == false (no element longer than 10 chars)
```