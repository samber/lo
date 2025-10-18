---
name: FirstOrEmpty
slug: firstorempty
sourceRef: it/find.go#L370
category: it
subCategory: find
signatures:
  - "func FirstOrEmpty[T any](collection iter.Seq[T]) T"
playUrl: "https://go.dev/play/p/6NhAE0-zm"
variantHelpers:
  - it#find#firstorempty
similarHelpers:
  - core#slice#firstorempty
position: 540
---

Returns the first element of a collection or zero value if empty.

Will iterate at most once.

Examples:

```go
// Get the first element or zero value
numbers := it.Slice([]int{5, 2, 8, 1, 9})
first := it.FirstOrEmpty(numbers)
// first: 5

// With empty collection
empty := it.Slice([]int{})
first := it.FirstOrEmpty(empty)
// first: 0 (zero value for int)

// With strings
words := it.Slice([]string{"hello", "world", "go"})
first := it.FirstOrEmpty(words)
// first: "hello"

emptyWords := it.Slice([]string{})
first := it.FirstOrEmpty(emptyWords)
// first: "" (zero value for string)

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
first := it.FirstOrEmpty(people)
// first: {Name: "Alice", Age: 30}

emptyPeople := it.Slice([]Person{})
first := it.FirstOrEmpty(emptyPeople)
// first: {Name: "", Age: 0} (zero value for Person)
```