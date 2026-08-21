---
name: FirstOrEmpty
slug: firstorempty
sourceRef: it/find.go#L396
category: iter
subCategory: find
signatures:
  - "func FirstOrEmpty[T any](collection iter.Seq[T]) T"
playUrl: https://go.dev/play/p/NTUTgPCfevx
variantHelpers:
  - iter#find#firstorempty
similarHelpers:
  - core#find#firstorempty
position: 540
---

Returns the first element of a collection or zero value if empty.

Will iterate at most once.

Examples:

```go
// Get the first element or zero value
numbers := slices.Values([]int{5, 2, 8, 1, 9})
first := it.FirstOrEmpty(numbers)
// first: 5

// With empty collection
empty := slices.Values([]int{})
first = it.FirstOrEmpty(empty)
// first: 0 (zero value for int)

// With strings
words := slices.Values([]string{"hello", "world", "go"})
firstWord := it.FirstOrEmpty(words)
// firstWord: "hello"

emptyWords := slices.Values([]string{})
firstWord = it.FirstOrEmpty(emptyWords)
// firstWord: "" (zero value for string)

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
firstPerson := it.FirstOrEmpty(people)
// firstPerson: {Name: "Alice", Age: 30}

emptyPeople := slices.Values([]Person{})
firstPerson = it.FirstOrEmpty(emptyPeople)
// firstPerson: {Name: "", Age: 0} (zero value for Person)
```