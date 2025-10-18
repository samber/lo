---
name: FindOrElse
slug: findorelse
sourceRef: it/find.go#L147
category: it
subCategory: find
signatures:
  - "func FindOrElse[T any](collection iter.Seq[T], fallback T, predicate func(item T) bool) T"
playUrl: "https://go.dev/play/p/8VgPO9mi-t"
variantHelpers:
  - it#find#findorelse
similarHelpers:
  - core#slice#findorelse
  - it#find#find
position: 70
---

Searches for an element using a predicate or returns a fallback value if not found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(40)
}
result := it.FindOrElse(seq, 99, func(x int) bool {
    return x > 25
})
// result == 30
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
result := it.FindOrElse(seq, "unknown", func(s string) bool {
    return len(s) > 10
})
// result == "unknown" (fallback value)
```