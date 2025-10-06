---
name: FindLastIndexOf
slug: findlastindexof
sourceRef: find.go#L101
category: core
subCategory: find
playUrl: https://go.dev/play/p/dPiMRtJ6cUx
variantHelpers:
  - core#find#findlastindexof
similarHelpers:
  - core#find#findindexof
  - core#find#find
  - core#find#lastindexof
  - core#find#findby
position: 60
signatures:
  - "func FindLastIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool)"
---

Searches for the last element matching the predicate and returns the element, its index, and a boolean indicating success.

```go
val, idx, ok := lo.FindLastIndexOf([]string{"a", "b", "a", "b"}, func(i string) bool {
    return i == "b"
})
// "b", 3, true
```


