---
name: KeyBy
slug: keyby
sourceRef: slice.go#L374
category: core
subCategory: slice
playUrl: https://go.dev/play/p/ccUiUL_Lnel
variantHelpers:
  - core#slice#keyby
similarHelpers:
  - core#slice#groupby
  - core#slice#partitionby
  - core#map#associate
  - core#slice#keyify
position: 230
signatures:
  - "func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V"
---

Transforms a slice to a map using a pivot callback to compute keys.

```go
m := lo.KeyBy(
    []string{"a", "aa", "aaa"},
    func(str string) int {
        return len(str)
    },
)
// map[int]string{1: "a", 2: "aa", 3: "aaa"}
```


