---
name: CrossJoinX
slug: crossjoinx
sourceRef: tuples.go#L891
category: core
subCategory: tuple
signatures:
  - "func CrossJoin2[A, B any](listA []A, listB []B) []Tuple2[A, B]"
  - "func CrossJoin3[A, B, C any](listA []A, listB []B, listC []C) []Tuple3[A, B, C]"
  - "func CrossJoin4[A, B, C, D any](listA []A, listB []B, listC []C, listD []D) []Tuple4[A, B, C, D]"
  - "func CrossJoin5[A, B, C, D, E any](listA []A, listB []B, listC []C, listD []D, listE []E) []Tuple5[A, B, C, D, E]"
  - "func CrossJoin6[A, B, C, D, E, F any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F) []Tuple6[A, B, C, D, E, F]"
  - "func CrossJoin7[A, B, C, D, E, F, G any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G) []Tuple7[A, B, C, D, E, F, G]"
  - "func CrossJoin8[A, B, C, D, E, F, G, H any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H) []Tuple8[A, B, C, D, E, F, G, H]"
  - "func CrossJoin9[A, B, C, D, E, F, G, H, I any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, listI []I) []Tuple9[A, B, C, D, E, F, G, H, I]"
playUrl: https://go.dev/play/p/3VFppyL9FDU
variantHelpers:
  - core#tuple#crossjoinx
similarHelpers:
  - core#tuple#tuplex
  - core#intersect#product
  - core#intersect#productby
  - core#map#entries
position: 50
---

Computes the cartesian product of input slices, returning tuples of all combinations. Variants support 2 up to 9 input slices.

Variants: `CrossJoin2..CrossJoin9`

```go
a := []int{1,2}
b := []string{"x","y"}
pairs := lo.CrossJoin2(a, b)
```


