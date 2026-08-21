---
name: LastOrEmpty
slug: lastorempty
sourceRef: it/find.go#L429
category: iter
subCategory: find
signatures:
  - "func LastOrEmpty[T any](collection iter.Seq[T]) T"
playUrl: https://go.dev/play/p/teODFK4YqM4
variantHelpers:
  - iter#find#lastorempty
similarHelpers:
  - core#find#lastorempty
position: 560
---

Returns the last element of a collection or zero value if empty.

Will iterate through the entire sequence.

Examples:

```go
// Get the last element or zero value
numbers := slices.Values([]int{5, 2, 8, 1, 9})
last := it.LastOrEmpty(numbers)
// last: 9

// With empty collection
empty := slices.Values([]int{})
last = it.LastOrEmpty(empty)
// last: 0 (zero value for int)

// With strings
words := slices.Values([]string{"hello", "world", "go"})
lastWord := it.LastOrEmpty(words)
// lastWord: "go"

emptyWords := slices.Values([]string{})
lastWord = it.LastOrEmpty(emptyWords)
// lastWord: "" (zero value for string)

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
lastPerson := it.LastOrEmpty(people)
// lastPerson: {Name: "Bob", Age: 25}

emptyPeople := slices.Values([]Person{})
lastPerson = it.LastOrEmpty(emptyPeople)
// lastPerson: {Name: "", Age: 0} (zero value for Person)

// With single element
single := slices.Values([]int{42})
last = it.LastOrEmpty(single)
// last: 42
```