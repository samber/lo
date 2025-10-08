---
name: Range
slug: range
sourceRef: it/math.go#L12
category: it
subCategory: math
signatures:
  - "func Range(elementNum int) iter.Seq[int]"
playUrl: "https://go.dev/play/p/6ksL0W6KEuQ"
variantHelpers:
  - it#math#range
similarHelpers:
  - core#slice#range
  - it#math#rangefrom
  - it#math#rangewithsteps
position: 0
---

Creates a sequence of integers starting from 0. Yields `elementNum` integers, stepping by ±1 depending on sign.

```go
seq := it.Range(4)
var out []int
for v := range seq { out = append(out, v) }
// out == []int{0, 1, 2, 3}
```