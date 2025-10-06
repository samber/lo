---
name: FilterMapToSeq
slug: filtermaptoseq
sourceRef: it/map.go#L178
category: it
subCategory: map
signatures:
  - "func FilterMapToSeq[K comparable, V, R any](in map[K]V, transform func(key K, value V) (R, bool)) iter.Seq[R]"
variantHelpers:
  - it#map#filtermaptoseq
similarHelpers:
  - it#map#maptoseq
  - core#map#filtermaptoslice
  - it#map#filterkeys
  - it#map#filtervalues
position: 54
---

Transforms a map into a sequence by applying a transform function to each key-value pair, but only includes values where the transform function returns true as the second value.

```go
m := map[string]int{
    "apple":  3,
    "banana": 0,
    "cherry": 2,
    "date":   0,
}
result := lo.FilterMapToSeq(m, func(key string, value int) (string, bool) {
    if value > 0 {
        return fmt.Sprintf("%s:%d", key, value), true
    }
    return "", false
})
// iter.Seq[string] yielding "apple:3", "cherry:2" (only entries with value > 0)

personMap := map[string]int{"alice": 25, "bob": 30, "charlie": 15}
type Person struct {
    Name string
    Age  int
}
result = lo.FilterMapToSeq(personMap, func(name string, age int) (Person, bool) {
    person := Person{Name: name, Age: age}
    return person, age >= 18
})
// iter.Seq[Person] yielding {Name: "alice", Age: 25}, {Name: "bob", Age: 30} (only adults)

dataMap := map[string]float64{"a": 1.5, "b": -2.0, "c": 3.14}
result = lo.FilterMapToSeq(dataMap, func(key string, value float64) (int, bool) {
    if value > 0 {
        return int(value * 100), true
    }
    return 0, false
})
// iter.Seq[int] yielding 150, 314 (1.5*100, 3.14*100 rounded)
```