---
name: LastOr
slug: lastor
sourceRef: it/find.go#L407
category: it
subCategory: find
signatures:
  - "func LastOr[T any](collection iter.Seq[T], fallback T) T"
playUrl: "https://go.dev/play/p/0RlEI4-zq"
variantHelpers:
  - it#find#lastor
similarHelpers:
  - core#slice#lastor
position: 570
---

Returns the last element of a collection or the fallback value if empty.

Will iterate through the entire sequence.

Examples:

```go
// Get the last element or fallback value
numbers := it.Slice([]int{5, 2, 8, 1, 9})
last := it.LastOr(numbers, 42)
// last: 9

// With empty collection
empty := it.Slice([]int{})
last := it.LastOr(empty, 42)
// last: 42 (fallback value)

// With strings
words := it.Slice([]string{"hello", "world", "go"})
last := it.LastOr(words, "fallback")
// last: "go"

emptyWords := it.Slice([]string{})
last := it.LastOr(emptyWords, "fallback")
// last: "fallback"

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
last := it.LastOr(people, Person{Name: "Default", Age: 0})
// last: {Name: "Bob", Age: 25}

emptyPeople := it.Slice([]Person{})
last := it.LastOr(emptyPeople, Person{Name: "Default", Age: 0})
// last: {Name: "Default", Age: 0} (fallback value)

// With single element
single := it.Slice([]int{42})
last := it.LastOr(single, 99)
// last: 42

// Using with nil pointer fallback
values := it.Slice([]*string{ptr("hello"), ptr("world")})
last := it.LastOr(values, nil)
// last: pointer to "world"

emptyValues := it.Slice([]*string{})
last := it.LastOr(emptyValues, nil)
// last: nil (fallback value)
```