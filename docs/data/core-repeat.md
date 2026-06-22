---
name: Repeat
slug: repeat
sourceRef: slice.go#L362
category: core
subCategory: slice
playUrl: https://go.dev/play/p/g3uHXbmc3b6
variantHelpers:
  - core#slice#repeat
similarHelpers:
  - core#slice#times
  - core#slice#repeatby
position: 220
signatures:
  - "func Repeat[T any](count int, initial T) []T"
---

Builds a slice with N copies of initial value.

```go
lo.Repeat(5, "42")
// []string{"42", "42", "42", "42", "42"}
```


