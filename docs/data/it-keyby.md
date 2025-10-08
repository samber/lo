---
name: KeyBy
slug: keyby
sourceRef: it/seq.go#L398
category: it
subCategory: sequence
signatures:
  - "func KeyBy[K comparable, V any](collection iter.Seq[V], transform func(item V) K) map[K]V"
playUrl:
variantHelpers:
  - it#map#associate
similarHelpers:
  - core#slice#keyby
  - core#slice#associate
position: 15
---

Transforms a sequence into a map using a transform function to generate keys.

```go
result := it.KeyBy(it.Range(1, 5), func(item int) string {
    return fmt.Sprintf("key-%d", item)
})
// map[string]int{"key-1": 1, "key-2": 2, "key-3": 3, "key-4": 4}
```