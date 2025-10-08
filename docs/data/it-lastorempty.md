---
name: LastOrEmpty
slug: lastorempty
sourceRef: it/find.go#L400
category: it
subCategory: find
signatures:
  - "func LastOrEmpty[T any](collection iter.Seq[T]) T"
playUrl: "https://go.dev/play/p/9QkDH3-zp"
variantHelpers:
  - it#find#lastorempty
similarHelpers:
  - core#slice#lastorempty
position: 560
---

Returns the last element of a collection or zero value if empty.

Will iterate through the entire sequence.

Examples:

```go
// Get the last element or zero value
numbers := it.Slice([]int{5, 2, 8, 1, 9})
last := it.LastOrEmpty(numbers)
// last: 9

// With empty collection
empty := it.Slice([]int{})
last := it.LastOrEmpty(empty)
// last: 0 (zero value for int)

// With strings
words := it.Slice([]string{"hello", "world", "go"})
last := it.LastOrEmpty(words)
// last: "go"

emptyWords := it.Slice([]string{})
last := it.LastOrEmpty(emptyWords)
// last: "" (zero value for string)

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
last := it.LastOrEmpty(people)
// last: {Name: "Bob", Age: 25}

emptyPeople := it.Slice([]Person{})
last := it.LastOrEmpty(emptyPeople)
// last: {Name: "", Age: 0} (zero value for Person)

// With single element
single := it.Slice([]int{42})
last := it.LastOrEmpty(single)
// last: 42
```