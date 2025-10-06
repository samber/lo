---
name: FindIndexOf
slug: findindexof
sourceRef: find.go#L87
category: core
subCategory: find
playUrl: https://go.dev/play/p/XWSEM4Ic_t0
variantHelpers:
  - core#find#findindexof
similarHelpers:
  - core#find#find
  - core#slice#indexof
  - core#find#findlastindexof
position: 50
signatures:
  - "func FindIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool)"
---

Searches for an element based on a predicate and returns the element, its index, and a boolean indicating success.

```go
val, idx, ok := lo.FindIndexOf([]string{"a", "b", "a", "b"}, func(i string) bool {
    return i == "b"
})
// "b", 1, true
```


