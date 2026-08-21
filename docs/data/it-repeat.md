---
name: Repeat
slug: repeat
sourceRef: it/seq.go#L504
category: iter
subCategory: sequence
signatures:
  - "func Repeat[T lo.Clonable[T]](count int, initial T) iter.Seq[T]"
playUrl: https://go.dev/play/p/xs-aq0p_uDP
variantHelpers:
  - iter#sequence#repeatby
similarHelpers:
  - core#slice#repeat
  - core#slice#repeatby
position: 75
---

Creates a sequence that repeats the initial value count times.

```go
type Point struct {
    X, Y int
}

func (p Point) Clone() Point {
    return p
}

result := it.Repeat(3, Point{X: 1, Y: 2})
// iter.Seq[Point] yielding {1, 2}, {1, 2}, {1, 2}
```