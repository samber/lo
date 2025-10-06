---
name: UnzipByX
slug: unzipbyx
sourceRef: tuples.go#L698
category: core
subCategory: tuple
signatures:
  - "func UnzipBy2[In any, A any, B any](items []In, predicate func(In) (a A, b B)) ([]A, []B)"
  - "func UnzipBy3[In any, A any, B any, C any](items []In, predicate func(In) (a A, b B, c C)) ([]A, []B, []C)"
  - "func UnzipBy4[In any, A any, B any, C any, D any](items []In, predicate func(In) (a A, b B, c C, d D)) ([]A, []B, []C, []D)"
  - "func UnzipBy5[In any, A any, B any, C any, D any, E any](items []In, predicate func(In) (a A, b B, c C, d D, e E)) ([]A, []B, []C, []D, []E)"
  - "func UnzipBy6[In any, A any, B any, C any, D any, E any, F any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F)) ([]A, []B, []C, []D, []E, []F)"
  - "func UnzipBy7[In any, A any, B any, C any, D any, E any, F any, G any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, g G)) ([]A, []B, []C, []D, []E, []F, []G)"
  - "func UnzipBy8[In any, A any, B any, C any, D any, E any, F any, G any, H any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, g G, h H)) ([]A, []B, []C, []D, []E, []F, []G, []H)"
  - "func UnzipBy9[In any, A any, B any, C any, D any, E any, F any, G any, H any, I any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, g G, h H, i I)) ([]A, []B, []C, []D, []E, []F, []G, []H, []I)"
playUrl: https://go.dev/play/p/tN8yqaRZz0r
variantHelpers:
  - core#tuple#unzipbyx
similarHelpers:
  - core#tuple#tuplex
  - core#tuple#unpackx
  - core#tuple#zipx
  - core#tuple#zipbyx
  - core#tuple#unzipx
  - core#slice#map
  - core#slice#filtermap
position: 50
---

Transforms each input element into a tuple and splits results into parallel slices. Variants support arities from 2 to 9.

Variants: `UnzipBy2..UnzipBy9`

```go
type User struct{ ID int; Name string }
ids, names := lo.UnzipBy2([]User{{1,"a"},{2,"b"}}, func(u User) (int, string) {
    return u.ID, u.Name
})
```


