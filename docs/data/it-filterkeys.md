---
name: FilterKeys
slug: filterkeys
sourceRef: it/map.go#L238
category: it
subCategory: map
signatures:
  - "func FilterKeys[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) []K"
variantHelpers:
  - it#map#filterkeys
similarHelpers:
  - core#map#filterkeys
  - it#map#filtervalues
  - it#map#keys
  - it#map#values
position: 56
---

Filters map keys based on a predicate function that takes both key and value. Returns a slice of keys that satisfy the predicate.

```go
m := map[string]int{
    "apple":  3,
    "banana": 5,
    "cherry": 2,
    "date":   0,
}
result := lo.FilterKeys(m, func(key string, value int) bool {
    return value > 2
})
// []string{"apple", "banana"} (keys with values > 2)

numberMap := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
result = lo.FilterKeys(numberMap, func(key int, value string) bool {
    return len(value) == 3
})
// []int{1, 2, 3} (keys where value length is 3)

personMap := map[string]int{"alice": 25, "bob": 30, "charlie": 17}
result = lo.FilterKeys(personMap, func(key string, age int) bool {
    return strings.HasPrefix(key, "a") && age >= 20
})
// []string{"alice"} (keys starting with "a" and age >= 20)

emptyMap := map[string]int{}
result = lo.FilterKeys(emptyMap, func(key string, value int) bool {
    return true
})
// []string{} (empty map)
```