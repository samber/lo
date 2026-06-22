---
name: Range
slug: range
sourceRef: it/math.go#L12
category: iter
subCategory: math
signatures:
  - "func Range(elementNum int) iter.Seq[int]"
playUrl: "https://go.dev/play/p/79QUZBa8Ukn"
variantHelpers:
  - iter#math#range
similarHelpers:
  - core#slice#range
  - iter#math#rangefrom
  - iter#math#rangewithsteps
position: 0
---

Creates a sequence of integers starting from 0. Yields `elementNum` integers, stepping by ±1 depending on sign.

```go
seq := it.Range(4)
var out []int
for v := range seq { out = append(out, v) }
// out == []int{0, 1, 2, 3}
```