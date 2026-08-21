---
name: LastOr
slug: lastor
sourceRef: it/find.go#L437
category: iter
subCategory: find
signatures:
  - "func LastOr[T any](collection iter.Seq[T], fallback T) T"
playUrl: https://go.dev/play/p/HNubjW2Mrxs
variantHelpers:
  - iter#find#lastor
similarHelpers:
  - core#find#lastor
position: 570
---

Returns the last element of a collection or the fallback value if empty.

Will iterate through the entire sequence.

Examples:

```go
// Get the last element or fallback value
numbers := slices.Values([]int{5, 2, 8, 1, 9})
last := it.LastOr(numbers, 42)
// last: 9

// With empty collection
empty := slices.Values([]int{})
last = it.LastOr(empty, 42)
// last: 42 (fallback value)

// With strings
words := slices.Values([]string{"hello", "world", "go"})
lastStr := it.LastOr(words, "fallback")
// lastStr: "go"

emptyWords := slices.Values([]string{})
lastStr = it.LastOr(emptyWords, "fallback")
// lastStr: "fallback"

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
lastPerson := it.LastOr(people, Person{Name: "Default", Age: 0})
// lastPerson: {Name: "Bob", Age: 25}

emptyPeople := slices.Values([]Person{})
lastPerson = it.LastOr(emptyPeople, Person{Name: "Default", Age: 0})
// lastPerson: {Name: "Default", Age: 0} (fallback value)

// With single element
single := slices.Values([]int{42})
last = it.LastOr(single, 99)
// last: 42

// Using with nil pointer fallback
values := slices.Values([]*string{lo.ToPtr("hello"), lo.ToPtr("world")})
lastPtr := it.LastOr(values, nil)
// lastPtr: pointer to "world"

emptyValues := slices.Values([]*string{})
lastPtr = it.LastOr(emptyValues, nil)
// lastPtr: nil (fallback value)
```