---
name: FindOrElse
slug: findorelse
sourceRef: find.go#L116
category: core
subCategory: find
playUrl: https://go.dev/play/p/Eo7W0lvKTky
variantHelpers:
  - core#find#findorelse
similarHelpers:
  - core#find#find
  - core#slice#firstor
  - core#slice#lastor
  - core#slice#nthor
position: 70
signatures:
  - "func FindOrElse[T any](collection []T, fallback T, predicate func(item T) bool) T"
---

Searches for an element based on a predicate and returns it if found, otherwise returns the fallback.

```go
value := lo.FindOrElse([]string{"a", "b", "c", "d"}, "x", func(i string) bool {
    return i == "b"
})
// "b"

value = lo.FindOrElse([]string{"foobar"}, "x", func(i string) bool {
    return i == "b"
})
// "x"
```


