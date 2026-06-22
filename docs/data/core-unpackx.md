---
name: UnpackX
slug: unpackx
sourceRef: tuples.go#L53
category: core
subCategory: tuple
signatures:
  - "func Unpack2[A, B any](tuple Tuple2[A, B]) (A, B)"
playUrl: https://go.dev/play/p/xVP_k0kJ96W
variantHelpers:
  - core#tuple#unpackx
similarHelpers:
  - core#tuple#tuplex
  - core#tuple#zipx
  - core#tuple#zipbyx
  - core#tuple#unzipx
  - core#tuple#unzipbyx
position: 10
---

Extracts values from tuples. Variants support tuple sizes from 2 to 9.

Variants: `Unpack2..Unpack9`

```go
a, b, c := lo.Unpack3(lo.T3(1, "a", true))
```


