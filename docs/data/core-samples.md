---
name: Samples
slug: samples
sourceRef: find.go#L679
category: core
subCategory: find
playUrl: https://go.dev/play/p/QYRD8aufD0C
variantHelpers:
  - core#find#samples
similarHelpers:
  - core#find#sample
  - core#find#samplesby
  - core#find#shuffle
position: 370
signatures:
  - "func Samples[T any, Slice ~[]T](collection Slice, count int) Slice"
---

Returns N random unique items from a collection.

```go
v := lo.Samples(
    []int{10, 20, 30},
    2,
)
```


