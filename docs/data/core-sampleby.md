---
name: SampleBy
slug: sampleby
sourceRef: find.go#L669
category: core
subCategory: find
playUrl: https://go.dev/play/p/HDmKmMgq0XN
variantHelpers:
  - core#find#sampleby
similarHelpers:
  - core#find#sample
  - core#find#samples
  - core#find#samplesby
  - core#find#shuffle
position: 360
signatures:
  - "func SampleBy[T any](collection []T, randomIntGenerator randomIntGenerator) T"
---

Returns a random item from a collection, using the provided random index generator.

```go
v := lo.SampleBy([]int{10, 20, 30}, func(n int) int {
    return 0
})
// v == 10
```


