---
name: Assign
slug: assign
sourceRef: it/map.go#L122
category: it
subCategory: map
signatures:
  - "func Assign[K comparable, V any, Map ~map[K]V](maps ...iter.Seq[Map]) Map"
variantHelpers:
  - it#map#assign
similarHelpers:
  - core#map#assign
  - it#map#fromentries
  - it#map#invert
position: 50
---

Merges multiple map sequences into a single map. Later maps overwrite values from earlier maps when keys conflict.

```go
map1 := func(yield func(map[string]int) bool) {
    yield(map[string]int{"a": 1, "b": 2})
}
map2 := func(yield func(map[string]int) bool) {
    yield(map[string]int{"b": 3, "c": 4})
}
map3 := func(yield func(map[string]int) bool) {
    yield(map[string]int{"d": 5, "e": 6})
}
result := lo.Assign(map1, map2, map3)
// map[string]int{"a": 1, "b": 3, "c": 4, "d": 5, "e": 6}
// Note: "b" is 3 (overwritten from map2)

singleMap := func(yield func(map[int]string) bool) {
    yield(map[int]string{1: "one", 2: "two"})
}
result = lo.Assign(singleMap)
// map[int]string{1: "one", 2: "two"}

emptyMap1 := func(yield func(map[string]bool) bool) {
    yield(map[string]bool{})
}
emptyMap2 := func(yield func(map[string]bool) bool) {
    yield(map[string]bool{"active": true})
}
result = lo.Assign(emptyMap1, emptyMap2)
// map[string]bool{"active": true}
```