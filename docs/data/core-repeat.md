---
name: Repeat
slug: repeat
sourceRef: slice.go#L601
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

Builds a slice with N copies of initial value. The value type must implement `Clone() T`.

```go
type Point struct {
    X, Y int
}

func (p Point) Clone() Point {
    return p
}

result := lo.Repeat(3, Point{X: 1, Y: 2})
// []Point{{1, 2}, {1, 2}, {1, 2}}
```


