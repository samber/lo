---
name: Mode
slug: mode
sourceRef: math.go#L149
category: core
subCategory: math
playUrl: 
variantHelpers:
  - core#math#mode
similarHelpers:
  - core#math#mean
  - core#math#meanby
  - core#math#sum
  - core#math#sumby
  - core#math#product
  - core#math#productby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
  - core#math#countvalues
  - core#math#countvaluesby
position: 100
signatures:
  - "func Mode[T constraints.Integer | constraints.Float](collection []T) []T"
---

Returns the mode(s), i.e., the most frequent value(s) in a collection. If multiple values share the highest frequency, returns all. Empty input yields an empty slice.

```go
lo.Mode([]int{2, 2, 3, 3})
// []int{2, 3}
```


