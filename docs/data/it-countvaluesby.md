---
name: CountValuesBy
slug: countvaluesby
sourceRef: it/seq.go#L720
category: it
subCategory: slice
signatures:
  - "func CountValuesBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U]int"
variantHelpers: []
similarHelpers:
  - core#slice#countvaluesby
position: 204
---

CountValuesBy counts the number of each element returned from transform function.
Is equivalent to chaining Map and CountValues.

```go
type Person struct {
    Name string
    Age  int
}

collection := func(yield func(Person) bool) {
    yield(Person{"Alice", 25})
    yield(Person{"Bob", 30})
    yield(Person{"Charlie", 25})
    yield(Person{"Diana", 30})
}

countsByAge := it.CountValuesBy(collection, func(p Person) int {
    return p.Age
})
// countsByAge contains {25: 2, 30: 2}
```