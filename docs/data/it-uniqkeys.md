---
name: UniqKeys
slug: uniqkeys
sourceRef: it/map.go#L26
category: it
subCategory: map
signatures:
  - "func UniqKeys[K comparable, V any](in ...map[K]V) iter.Seq[K]"
playUrl: "https://go.dev/play/p/_NicwfgAHbO"
variantHelpers:
  - it#map#uniqkeys
similarHelpers:
  - core#slice#uniqkeys
position: 800
---

Creates a sequence of unique keys from multiple maps.

Will allocate a map large enough to hold all distinct input keys.
Long input sequences with heterogeneous keys can cause excessive memory usage.

Examples:

```go
// Single map
m1 := map[string]int{
    "apple":  1,
    "banana": 2,
    "cherry": 3,
}
uniqueKeys := it.UniqKeys(m1)
// uniqueKeys: sequence with "apple", "banana", "cherry"

// Multiple maps with duplicate keys
m1 := map[string]int{
    "apple":  1,
    "banana": 2,
}
m2 := map[string]int{
    "banana": 3,
    "cherry": 4,
    "apple":  5,
}
uniqueKeys = it.UniqKeys(m1, m2)
// uniqueKeys: sequence with "apple", "banana", "cherry" (no duplicates)

// Maps with integer keys
scores1 := map[int]string{
    1: "Alice",
    2: "Bob",
    3: "Charlie",
}
scores2 := map[int]string{
    3: "David",
    4: "Eve",
    1: "Frank",
}
uniqueKeys = it.UniqKeys(scores1, scores2)
// uniqueKeys: sequence with 1, 2, 3, 4

// Maps with struct keys
type Person struct {
    Name string
    Age  int
}
people1 := map[Person]bool{
    {Name: "Alice", Age: 30}: true,
    {Name: "Bob", Age: 25}:   true,
}
people2 := map[Person]bool{
    {Name: "Bob", Age: 25}:     false,  // Same struct
    {Name: "Charlie", Age: 35}: true,
}
uniqueKeys = it.UniqKeys(people1, people2)
// uniqueKeys: sequence with {Alice 30}, {Bob 25}, {Charlie 35}

// Empty maps
empty1 := map[string]int{}
empty2 := map[string]int{}
uniqueKeys = it.UniqKeys(empty1, empty2)
// uniqueKeys: empty sequence

// Mix of empty and non-empty maps
m1 := map[string]int{"a": 1}
empty := map[string]int{}
m2 := map[string]int{"b": 2}
uniqueKeys = it.UniqKeys(m1, empty, m2)
// uniqueKeys: sequence with "a", "b"

// Maps with same keys but different values
m1 := map[string]int{
    "key1": 10,
    "key2": 20,
}
m2 := map[string]int{
    "key1": 100,  // Same key, different value
    "key3": 30,
}
uniqueKeys = it.UniqKeys(m1, m2)
// uniqueKeys: sequence with "key1", "key2", "key3" (key1 appears once)

// Large number of maps
maps := make([]map[int]string, 100)
for i := range maps {
    maps[i] = map[int]string{i: fmt.Sprintf("value%d", i)}
}
uniqueKeys = it.UniqKeys(maps...)
// uniqueKeys: sequence with keys 0, 1, 2, ..., 99
```