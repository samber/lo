---
name: Find
slug: find
sourceRef: find.go#L73
category: core
subCategory: find
playUrl: https://go.dev/play/p/Eo7W0lvKTky
variantHelpers:
  - core#find#find
similarHelpers:
  - core#find#findorelse
  - core#find#findkey
  - core#find#findindexof
  - core#slice#filter
  - core#slice#first
  - core#slice#last
position: 40
signatures:
  - "func Find[T any](collection []T, predicate func(item T) bool) (T, bool)"
---

Searches for an element in a slice based on a predicate and returns the element with a boolean indicating success.

```go
value, ok := lo.Find([]string{"a", "b", "c", "d"}, func(i string) bool {
    return i == "b"
})
// "b", true

value, ok = lo.Find([]string{"foobar"}, func(i string) bool {
    return i == "b"
})
// "", false
```


