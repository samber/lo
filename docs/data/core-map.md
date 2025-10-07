---
name: Map
slug: map
sourceRef: slice.go#L26
category: core
subCategory: slice
playUrl: https://go.dev/play/p/OkPcYAhBo0D
similarHelpers:
  - core#slice#filtermap
  - core#slice#flatmap
  - core#slice#uniqmap
  - core#slice#rejectmap
  - core#slice#mapkeys
  - core#slice#mapvalues
  - core#slice#mapentries
  - core#slice#maptoslice
  - core#slice#filtermaptoslice
  - parallel#slice#map
  - mutable#slice#map
variantHelpers:
  - core#slice#map
position: 10
signatures:
  - "func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R"
---

Transforms each element in a slice to a new type using a function. Takes both the element and its index, making it useful for transformations that need positional context.

```go
// Basic type transformation
transformed := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
    return strconv.FormatInt(x, 10)
})
// transformed: []string{"1", "2", "3", "4"}
```

```go
// Transforming structs
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

people := []Person{
    {FirstName: "John", LastName: "Doe", Age: 25},
    {FirstName: "Jane", LastName: "Smith", Age: 30},
}

fullNames := lo.Map(people, func(p Person, index int) string {
    return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
})
// fullNames: []string{"John Doe", "Jane Smith"}
```
