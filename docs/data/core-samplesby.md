---
name: SamplesBy
slug: samplesby
sourceRef: find.go#L686
category: core
subCategory: find
playUrl: https://go.dev/play/p/HDmKmMgq0XN
variantHelpers:
  - core#find#samplesby
similarHelpers:
  - core#find#sample
  - core#find#samples
  - core#find#shuffle
position: 380
signatures:
  - "func SamplesBy[T any, Slice ~[]T](collection Slice, count int, randomIntGenerator randomIntGenerator) Slice"
---

Returns N random unique items from a collection, using the provided random index generator.

```go
v := lo.SamplesBy(
    []int{10, 20, 30},
    2,
    func(n int) int {
        return 0
    },
)
```


