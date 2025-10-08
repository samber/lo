---
name: SampleBy
slug: sampleby
sourceRef: it/find.go#L455
category: it
subCategory: find
signatures:
  - "func SampleBy[T any](collection iter.Seq[T], randomIntGenerator func(int) int) T"
playUrl: "https://go.dev/play/p/5WqJN9-zv"
variantHelpers:
  - it#find#sampleby
similarHelpers:
  - core#slice#sample
  - it#find#sample
  - it#find#samples
  - it#find#samplesby
position: 160
---

Returns a random item from collection, using a custom random index generator.

Example:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
// Use custom RNG for predictable results (returns first element)
item := it.SampleBy(seq, func(max int) int { return 0 })
// item == 1
```