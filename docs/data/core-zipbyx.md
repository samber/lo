---
name: ZipByX
slug: zipbyx
sourceRef: tuples.go#L335
category: core
subCategory: tuple
signatures:
  - "func ZipBy2[A any, B any, Out any](a []A, b []B, predicate func(a A, b B) Out) []Out"
  - "func ZipBy3[A any, B any, C any, Out any](a []A, b []B, c []C, predicate func(a A, b B, c C) Out) []Out"
  - "func ZipBy4[A any, B any, C any, D any, Out any](a []A, b []B, c []C, d []D, predicate func(a A, b B, c C, d D) Out) []Out"
  - "func ZipBy5[A any, B any, C any, D any, E any, Out any](a []A, b []B, c []C, d []D, e []E, predicate func(a A, b B, c C, d D, e E) Out) []Out"
  - "func ZipBy6[A any, B any, C any, D any, E any, F any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, predicate func(a A, b B, c C, d D, e E, f F) Out) []Out"
  - "func ZipBy7[A any, B any, C any, D any, E any, F any, G any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, predicate func(a A, b B, c C, d D, e E, f F, g G) Out) []Out"
  - "func ZipBy8[A any, B any, C any, D any, E any, F any, G any, H any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, predicate func(a A, b B, c C, d D, e E, f F, g G, h H) Out) []Out"
  - "func ZipBy9[A any, B any, C any, D any, E any, F any, G any, H any, I any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, predicate func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Out) []Out"
playUrl: https://go.dev/play/p/wlHur6yO8rR
variantHelpers:
  - core#tuple#zipbyx
similarHelpers:
  - core#tuple#tuplex
  - core#tuple#unpackx
  - core#tuple#zipx
  - core#tuple#unzipx
  - core#tuple#unzipbyx
  - core#slice#map
  - core#slice#filtermap
position: 30
---

Zips multiple slices and projects each grouped set through a function. Variants support 2 up to 9 input slices.

Variants: `ZipBy2..ZipBy9`

```go
xs := []int{1,2}
ys := []string{"a","b"}
pairs := lo.ZipBy2(xs, ys, func(x int, y string) string {
    return fmt.Sprintf("%d-%s", x, y)
})
```


