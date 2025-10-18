---
name: CrossJoinByX
slug: crossjoinbyx
sourceRef: tuples.go#L956
category: core
subCategory: tuple
signatures:
  - "func CrossJoinBy2[A, B, Out any](listA []A, listB []B, project func(a A, b B) Out) []Out"
  - "func CrossJoinBy3[A, B, C, Out any](listA []A, listB []B, listC []C, project func(a A, b B, c C) Out) []Out"
  - "func CrossJoinBy4[A, B, C, D, Out any](listA []A, listB []B, listC []C, listD []D, project func(a A, b B, c C, d D) Out) []Out"
  - "func CrossJoinBy5[A, B, C, D, E, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, project func(a A, b B, c C, d D, e E) Out) []Out"
  - "func CrossJoinBy6[A, B, C, D, E, F, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, project func(a A, b B, c C, d D, e E, f F) Out) []Out"
  - "func CrossJoinBy7[A, B, C, D, E, F, G, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, project func(a A, b B, c C, d D, e E, f F, g G) Out) []Out"
  - "func CrossJoinBy8[A, B, C, D, E, F, G, H, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, project func(a A, b B, c C, d D, e E, f F, g G, h H) Out) []Out"
  - "func CrossJoinBy9[A, B, C, D, E, F, G, H, I, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, listI []I, project func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Out) []Out"
playUrl: https://go.dev/play/p/8Y7btpvuA-C
variantHelpers:
  - core#tuple#crossjoinbyx
similarHelpers:
  - core#tuple#zipx
  - core#tuple#unzipx
  - core#tuple#zipbyx
  - core#tuple#unzipbyx
  - core#slice#product
  - core#slice#productby
position: 60
---

Computes a cartesian product and projects each combination through a function. Variants support 2 up to 9 input slices.

Variants: `CrossJoinBy2..CrossJoinBy9`

```go
a := []int{1,2}
b := []string{"x","y"}
out := lo.CrossJoinBy2(a, b, func(x int, y string) string {
    return fmt.Sprintf("%d-%s", x, y)
})
```


