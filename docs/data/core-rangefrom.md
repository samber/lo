---
name: RangeFrom
slug: rangefrom
sourceRef: math.go#L21
category: core
subCategory: math
playUrl: https://go.dev/play/p/0r6VimXAi9H
variantHelpers:
  - core#math#rangefrom
similarHelpers:
  - core#math#range
  - core#math#rangewithsteps
  - core#math#times
  - core#slice#repeat
  - core#slice#repeatby
position: 10
signatures:
  - "func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) []T"
---

Creates a slice starting at `start` with the specified length. Negative length yields a descending sequence.

```go
lo.RangeFrom(1, 5)
// []int{1, 2, 3, 4, 5}
```


