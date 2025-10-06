---
name: MapToSeq
slug: maptoseq
sourceRef: it/map.go#L164
category: it
subCategory: map
signatures:
  - "func MapToSeq[K comparable, V, R any](in map[K]V, transform func(key K, value V) R) iter.Seq[R]"
variantHelpers:
  - it#map#maptoseq
similarHelpers:
  - core#map#maptoslice
  - it#map#values
  - it#map#keys
  - it#map#entries
  - it#map#filtermaptoseq
position: 52
---

Transforms a map into a sequence by applying a transform function to each key-value pair. The transform function determines what values are yielded in the output sequence.

```go
m := map[string]int{
    "apple":  3,
    "banana": 5,
    "cherry": 2,
}
result := lo.MapToSeq(m, func(key string, value int) string {
    return fmt.Sprintf("%s:%d", key, value)
})
// iter.Seq[string] yielding "apple:3", "banana:5", "cherry:2"

numberMap := map[int]string{1: "one", 2: "two", 3: "three"}
result = lo.MapToSeq(numberMap, func(key int, value string) int {
    return key * len(value)
})
// iter.Seq[int] yielding 3, 6, 15 (1*3, 2*3, 3*5)

personMap := map[string]int{"alice": 25, "bob": 30}
type Person struct {
    Name string
    Age  int
}
result = lo.MapToSeq(personMap, func(name string, age int) Person {
    return Person{Name: name, Age: age}
})
// iter.Seq[Person] yielding {Name: "alice", Age: 25}, {Name: "bob", Age: 30}
```