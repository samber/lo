---
name: Values
slug: values
sourceRef: it/map.go#L44
category: it
subCategory: map
signatures:
  - "func Values[K comparable, V any](in ...map[K]V) iter.Seq[V]"
playUrl: "https://go.dev/play/p/L9KcJ3h8E4f"
variantHelpers:
  - it#map#values
similarHelpers:
  - core#slice#values
  - it#map#keys
  - it#map#uniqvalues
position: 10
---

Creates a sequence of the map values. Accepts multiple maps and concatenates their values.

Examples:

```go
m1 := map[string]int{
    "apple":  1,
    "banana": 2,
}
m2 := map[string]int{
    "cherry": 3,
    "date":   4,
}
valuesSeq := it.Values(m1, m2)
var result []int
for v := range valuesSeq {
    result = append(result, v)
}
// result contains values from both maps
```