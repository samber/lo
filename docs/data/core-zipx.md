---
name: ZipX
slug: zipx
sourceRef: tuples.go#L103
category: core
subCategory: tuple
signatures:
  - "func Zip2[A, B any](a []A, b []B) []Tuple2[A, B]"
playUrl: https://go.dev/play/p/jujaA6GaJTp
variantHelpers:
  - core#tuple#zipx
similarHelpers:
  - core#tuple#tuplex
  - core#tuple#unpackx
  - core#tuple#zipbyx
  - core#tuple#unzipx
  - core#tuple#unzipbyx
  - core#slice#interleave
position: 20
---

Zips multiple slices into a slice of tuples. Variants support 2 up to 9 input slices.

Variants: `Zip2..Zip9`

```go
xs := []int{1,2}
ys := []string{"a","b"}
pairs := lo.Zip2(xs, ys)
```


