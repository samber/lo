---
name: FindKey
slug: findkey
sourceRef: find.go#L128
category: core
subCategory: find
playUrl: https://go.dev/play/p/Bg0w1VDPYXx
variantHelpers:
  - core#find#findkey
similarHelpers:
  - core#find#findkeyby
  - core#find#find
  - core#find#findby
position: 80
signatures:
  - "func FindKey[K comparable, V comparable](object map[K]V, value V) (K, bool)"
---

Returns the first key whose value equals the provided value.

```go
k, ok := lo.FindKey(map[string]int{"foo":1, "bar":2, "baz":3}, 2)
// "bar", true
```


