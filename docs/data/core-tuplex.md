---
name: TupleX
slug: tuplex
sourceRef: tuples.go#L5
category: core
subCategory: tuple
signatures:
  - "func T2[A, B any](a A, b B) lo.Tuple2[A, B]"
  - "func T3[A, B, C any](a A, b B, c C) lo.Tuple3[A, B, C]"
  - "func T4[A, B, C, D any](a A, b B, c C, d D) lo.Tuple4[A, B, C, D]"
  - "func T5[A, B, C, D, E any](a A, b B, c C, d D, e E) lo.Tuple5[A, B, C, D, E]"
  - "func T6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) lo.Tuple6[A, B, C, D, E, F]"
  - "func T7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) lo.Tuple7[A, B, C, D, E, F, G]"
  - "func T8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) lo.Tuple8[A, B, C, D, E, F, G, H]"
  - "func T9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) lo.Tuple9[A, B, C, D, E, F, G, H, I]"
playUrl: https://go.dev/play/p/IllL3ZO4BQm
variantHelpers:
  - core#tuple#tx
  - core#tuple#tuplex
similarHelpers:
  - core#tuple#unpackx
  - core#tuple#zipx
  - core#tuple#zipbyx
  - core#tuple#unzipx
  - core#tuple#unzipbyx
position: 0
---

Constructors for tuple values from 2 up to 9 elements.

Variants: `T2..T9`

```go
t := lo.T3(1, "a", true)
// lo.Tuple3[int, string, bool]{A:1, B:"a", C:true}
```


