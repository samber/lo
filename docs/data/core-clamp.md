---
name: Clamp
slug: clamp
sourceRef: math.go#L59
category: core
subCategory: math
playUrl: https://go.dev/play/p/RU4lJNC2hlI
variantHelpers:
  - core#math#clamp
similarHelpers:
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
  - core#math#mean
  - core#math#sum
  - core#math#product
position: 30
signatures:
  - "func Clamp[T constraints.Ordered](value T, mIn T, mAx T) T"
---

Clamps a number within inclusive lower and upper bounds.

```go
lo.Clamp(42, -10, 10)
// 10
```


