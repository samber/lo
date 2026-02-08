---
name: Distinct
slug: distinct
sourceRef: channel.go#L329
category: core
subCategory: channel
signatures:
  - "func Distinct[T comparable](upstream <-chan T) <-chan T"
similarHelpers:
  - core#slice#uniq
  - core#slice#uniqby
position: 258
---

Removes duplicate values from a channel while preserving order.

```go
ch := lo.SliceToChannel(10, []int{1, 2, 1, 3, 2, 4})
distinct := lo.Distinct(ch)
var result []int
for v := range distinct {
    result = append(result, v)
}
// result contains [1, 2, 3, 4]
```
