---
name: TupleX
slug: tuplex
sourceRef: tuples.go#L5
category: core
subCategory: tuple
signatures:
  - "func T2[A, B any](a A, b B) Tuple2[A, B]"
playUrl: https://go.dev/play/p/IllL3ZO4BQm
variantHelpers:
  - core#tuple#t2
  - core#tuple#t3
  - core#tuple#t4
  - core#tuple#t5
  - core#tuple#t6
  - core#tuple#t7
  - core#tuple#t8
  - core#tuple#t9
  - core#tuple#tuple2
  - core#tuple#tuple3
  - core#tuple#tuple4
  - core#tuple#tuple5
  - core#tuple#tuple6
  - core#tuple#tuple7
  - core#tuple#tuple8
  - core#tuple#tuple9
similarHelpers:
  - core#tuple#unpackx
  - core#tuple#zipx
  - core#tuple#zipbyx
  - core#tuple#unzipx
  - core#tuple#unzipbyx
position: 0
---

Constructors for tuple values from 2 up to 9 elements.

Variants:

- T2..T9: `TX` creates a `TupleX` from its `X` inputs.

```go
t := lo.T3(1, "a", true)
// Tuple3[int, string, bool]{A:1, B:"a", C:true}
```


