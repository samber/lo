---
name: UnzipX
slug: unzipx
sourceRef: tuples.go#L514
category: core
subCategory: tuple
signatures:
  - "func Unzip2[A, B any](tuples []Tuple2[A, B]) ([]A, []B)"
  - "func Unzip3[A, B, C any](tuples []Tuple3[A, B, C]) ([]A, []B, []C)"
  - "func Unzip4[A, B, C, D any](tuples []Tuple4[A, B, C, D]) ([]A, []B, []C, []D)"
  - "func Unzip5[A, B, C, D, E any](tuples []Tuple5[A, B, C, D, E]) ([]A, []B, []C, []D, []E)"
  - "func Unzip6[A, B, C, D, E, F any](tuples []Tuple6[A, B, C, D, E, F]) ([]A, []B, []C, []D, []E, []F)"
  - "func Unzip7[A, B, C, D, E, F, G any](tuples []Tuple7[A, B, C, D, E, F, G]) ([]A, []B, []C, []D, []E, []F, []G)"
  - "func Unzip8[A, B, C, D, E, F, G, H any](tuples []Tuple8[A, B, C, D, E, F, G, H]) ([]A, []B, []C, []D, []E, []F, []G, []H)"
  - "func Unzip9[A, B, C, D, E, F, G, H, I any](tuples []Tuple9[A, B, C, D, E, F, G, H, I]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I)"
playUrl: https://go.dev/play/p/ciHugugvaAW
variantHelpers:
  - core#tuple#unzip2
  - core#tuple#unzip3
  - core#tuple#unzip4
  - core#tuple#unzip5
  - core#tuple#unzip6
  - core#tuple#unzip7
  - core#tuple#unzip8
  - core#tuple#unzip9
similarHelpers:
  - core#tuple#tuplex
  - core#tuple#unpackx
  - core#tuple#zipx
  - core#tuple#zipbyx
  - core#tuple#unzipbyx
  - core#slice#mapkeys
  - core#slice#mapvalues
position: 40
---

Splits a slice of tuples back into multiple parallel slices. Variants support tuple sizes from 2 to 9.

Variants: `Unzip2..Unzip9`

```go
pairs := []lo.Tuple2[int, string]{
  lo.T2(1, "a"),
  lo.T2(2, "b"),
}
xs, ys := lo.Unzip2(pairs)
```


