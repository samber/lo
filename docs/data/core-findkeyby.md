---
name: FindKeyBy
slug: findkeyby
sourceRef: find.go#L140
category: core
subCategory: find
playUrl: https://go.dev/play/p/9IbiPElcyo8
variantHelpers:
  - core#find#findkeyby
similarHelpers:
  - core#find#findkey
  - core#find#findby
  - core#find#findorelse
position: 90
signatures:
  - "func FindKeyBy[K comparable, V any](object map[K]V, predicate func(key K, value V) bool) (K, bool)"
---

Returns the first key in the map for which the predicate returns true.

```go
k, ok := lo.FindKeyBy(map[string]int{"foo":1, "bar":2, "baz":3}, func(k string, v int) bool {
    return k == "foo"
})
// "foo", true
```


