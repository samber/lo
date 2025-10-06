---
name: UniqValues
slug: uniqvalues
sourceRef: it/map.go#L59
category: it
subCategory: map
signatures:
  - "func UniqValues[K, V comparable](in ...map[K]V) iter.Seq[V]"
playUrl: ""
variantHelpers:
  - it#map#uniqvalues
similarHelpers:
  - core#slice#uniqvalues
position: 810
---

Creates a sequence of unique values from multiple maps.

Will allocate a map large enough to hold all distinct input values.
Long input sequences with heterogeneous values can cause excessive memory usage.

Examples:

```go
// Single map
m1 := map[string]int{
    "apple":  1,
    "banana": 2,
    "cherry": 3,
}
uniqueValues := it.UniqValues(m1)
// uniqueValues: sequence with 1, 2, 3

// Multiple maps with duplicate values
m1 := map[string]int{
    "apple":  1,
    "banana": 2,
}
m2 := map[string]int{
    "orange": 1,  // Same value as "apple"
    "grape":  3,
}
uniqueValues = it.UniqValues(m1, m2)
// uniqueValues: sequence with 1, 2, 3 (no duplicates)

// Maps with string values
scores1 := map[int]string{
    1: "Alice",
    2: "Bob",
    3: "Charlie",
}
scores2 := map[int]string{
    4: "Alice",   // Same value
    5: "David",
    6: "Bob",     // Same value
}
uniqueValues = it.UniqValues(scores1, scores2)
// uniqueValues: sequence with "Alice", "Bob", "Charlie", "David"

// Maps with boolean values
boolMaps := []map[string]bool{
    {"enabled": true, "debug": false},
    {"test": true, "prod": false},  // Same values
    {"dev": true, "staging": false}, // Same values
}
uniqueValues = it.UniqValues(boolMaps...)
// uniqueValues: sequence with true, false

// Maps with float values
prices1 := map[string]float64{
    "apple":  1.99,
    "banana": 2.99,
}
prices2 := map[string]float64{
    "orange": 1.99,  // Same price as apple
    "grape":  3.99,
}
uniqueValues = it.UniqValues(prices1, prices2)
// uniqueValues: sequence with 1.99, 2.99, 3.99

// Maps with struct values
type Product struct {
    Name  string
    Price float64
}
products1 := map[int]Product{
    1: {Name: "Book", Price: 19.99},
    2: {Name: "Pen", Price: 1.99},
}
products2 := map[int]Product{
    3: {Name: "Notebook", Price: 19.99},  // Same price as book
    4: {Name: "Book", Price: 19.99},      // Same struct as products1[1]
}
uniqueValues = it.UniqValues(products1, products2)
// uniqueValues: sequence with {Book 19.99}, {Pen 1.99}, {Notebook 19.99}

// Maps with pointer values
type Person struct {
    Name string
}
alice := &Person{Name: "Alice"}
bob := &Person{Name: "Bob"}
people1 := map[string]*Person{
    "user1": alice,
    "user2": bob,
}
people2 := map[string]*Person{
    "user3": alice,  // Same pointer
    "user4": &Person{Name: "Charlie"},
}
uniqueValues = it.UniqValues(people1, people2)
// uniqueValues: sequence with pointers to Alice, Bob, Charlie

// Empty maps
empty1 := map[string]int{}
empty2 := map[string]int{}
uniqueValues = it.UniqValues(empty1, empty2)
// uniqueValues: empty sequence

// Mix of empty and non-empty maps
m1 := map[string]int{"a": 10}
empty := map[string]int{}
m2 := map[string]int{"b": 20, "c": 10}  // 10 is duplicate
uniqueValues = it.UniqValues(m1, empty, m2)
// uniqueValues: sequence with 10, 20

// Maps with same values from different keys
m1 := map[string]int{
    "key1": 100,
    "key2": 200,
}
m2 := map[string]int{
    "key3": 100,  // Same value as key1
    "key4": 200,  // Same value as key2
    "key5": 300,
}
uniqueValues = it.UniqValues(m1, m2)
// uniqueValues: sequence with 100, 200, 300

// Large number of duplicate values
maps := make([]map[int]string, 10)
for i := range maps {
    maps[i] = map[int]string{
        i: "common",  // All maps have the same value
        i + 100: fmt.Sprintf("unique_%d", i),
    }
}
uniqueValues = it.UniqValues(maps...)
// uniqueValues: sequence with "common" and 10 unique values
```