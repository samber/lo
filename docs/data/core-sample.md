---
name: Sample
slug: sample
sourceRef: find.go#L662
category: core
subCategory: find
playUrl: https://go.dev/play/p/vCcSJbh5s6l
variantHelpers:
  - core#find#sample
similarHelpers:
  - core#find#samples
  - core#find#samplesby
  - core#find#shuffle
position: 350
signatures:
  - "func Sample[T any](collection []T) T"
---

Returns a random item from a collection.

```go
v := lo.Sample(
    []int{10, 20, 30},
)
```


