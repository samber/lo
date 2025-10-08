---
name: Repeat
slug: repeat
sourceRef: it/seq.go#L384
category: it
subCategory: sequence
signatures:
  - "func Repeat[T lo.Clonable[T]](count int, initial T) iter.Seq[T]"
playUrl:
variantHelpers:
  - it#slice#repeatby
similarHelpers:
  - core#slice#repeat
  - core#slice#repeatby
position: 75
---

Creates a sequence that repeats the initial value count times.

```go
result := it.Repeat(3, "hello")
// ["hello", "hello", "hello"]
```