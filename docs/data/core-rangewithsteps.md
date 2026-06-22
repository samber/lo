---
name: RangeWithSteps
slug: rangewithsteps
sourceRef: math.go#L34
category: core
subCategory: math
playUrl: https://go.dev/play/p/0r6VimXAi9H
variantHelpers:
  - core#math#rangewithsteps
similarHelpers:
  - core#math#range
  - core#math#rangefrom
  - core#math#times
  - core#slice#repeat
  - core#slice#repeatby
position: 20
signatures:
  - "func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) []T"
---

Creates a slice progressing from `start` up to, but not including, `end` by `step`. Returns empty if `step` is 0 or direction mismatches.

```go
lo.RangeWithSteps(0, 20, 5)
// []int{0, 5, 10, 15}
```


