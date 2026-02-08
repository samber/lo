---
name: FilterValues
slug: filtervalues
sourceRef: it/map.go#L253
category: it
subCategory: map
signatures:
  - "func FilterValues[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) []V"
variantHelpers:
  - it#map#filtervalues
similarHelpers:
  - core#map#filtervalues
  - it#map#filterkeys
  - it#map#keys
  - it#map#values
position: 58
---

Filters map values based on a predicate function that takes both key and value. Returns a slice of values that satisfy the predicate.

```go
m := map[string]int{
    "apple":  3,
    "banana": 5,
    "cherry": 2,
    "date":   0,
}
result := it.FilterValues(m, func(key string, value int) bool {
    return value > 2
})
// []int{3, 5} (values > 2, corresponds to "apple" and "banana")

numberMap := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
result = it.FilterValues(numberMap, func(key int, value string) bool {
    return len(value) == 3
})
// []string{"one", "two", "three"} (values with length 3)

personMap := map[string]int{"alice": 25, "bob": 30, "charlie": 17}
result = it.FilterValues(personMap, func(key string, age int) bool {
    return strings.HasPrefix(key, "a") && age >= 20
})
// []int{25} (value for "alice" only)

emptyMap := map[string]int{}
result = it.FilterValues(emptyMap, func(key string, value int) bool {
    return true
})
// []int{} (empty map)

dataMap := map[string]float64{"a": 1.5, "b": -2.0, "c": 0.0, "d": 3.14}
result = it.FilterValues(dataMap, func(key string, value float64) bool {
    return value > 0
})
// []float64{1.5, 3.14} (positive values only)
```