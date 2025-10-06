---
name: Reverse
slug: reverse
sourceRef: mutable/slice.go#L65
category: mutable
subCategory: slice
playUrl: https://go.dev/play/p/O-M5pmCRgzV
variantHelpers:
  - mutable#slice#reverse
similarHelpers:
  - core#slice#reverse
position: 30
signatures:
  - "func Reverse[T any, Slice ~[]T](collection Slice)"
---

Reverses the slice in place so the first element becomes the last, the second becomes the second-to-last, and so on.

```go
import lom "github.com/samber/lo/mutable"

list := []int{0, 1, 2, 3, 4, 5}
lom.Reverse(list)
// list -> []int{5, 4, 3, 2, 1, 0}
```

With custom types:

```go
type Point struct{ X, Y int }
pts := []Point{{0,0}, {1,1}, {2,2}}
lom.Reverse(pts)
// pts -> []Point{{2,2}, {1,1}, {0,0}}
```

