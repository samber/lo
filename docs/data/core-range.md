---
name: Range
slug: range
sourceRef: math.go#L9
category: core
subCategory: math
playUrl: https://go.dev/play/p/0r6VimXAi9H
variantHelpers:
  - core#math#range
similarHelpers:
  - core#math#rangefrom
  - core#math#rangewithsteps
  - core#math#times
  - core#slice#repeat
  - core#slice#repeatby
position: 0
signatures:
  - "func Range(elementNum int) []int"
---

Creates a slice of integers of the given length starting at 0. Negative length produces a descending sequence.

```go
lo.Range(4)
// []int{0, 1, 2, 3}

lo.Range(-4)
// []int{0, -1, -2, -3}
```


