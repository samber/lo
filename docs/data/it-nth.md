---
name: Nth
slug: nth
sourceRef: it/find.go#L417
category: it
subCategory: find
signatures:
  - "func Nth[T any, N constraints.Integer](collection iter.Seq[T], nth N) (T, error)"
playUrl: ""
variantHelpers:
  - it#find#nth
similarHelpers:
  - core#slice#nth
position: 580
---

Returns the element at index `nth` of collection. Returns an error when nth is out of bounds.

Will iterate n times through the sequence.

Examples:

```go
// Get element at specific index
numbers := it.Slice([]int{5, 2, 8, 1, 9})
element, err := it.Nth(numbers, 2)
// element: 8, err: nil

// Get first element (index 0)
first, err := it.Nth(numbers, 0)
// first: 5, err: nil

// Get last element
last, err := it.Nth(numbers, 4)
// last: 9, err: nil

// Out of bounds - negative
_, err := it.Nth(numbers, -1)
// err: nth: -1 out of bounds

// Out of bounds - too large
_, err := it.Nth(numbers, 10)
// err: nth: 10 out of bounds

// With strings
words := it.Slice([]string{"hello", "world", "go", "lang"})
element, err := it.Nth(words, 1)
// element: "world", err: nil

// With different integer types
numbers := it.Slice([]int{1, 2, 3, 4, 5})
element, err := it.Nth(numbers, int8(3))
// element: 4, err: nil

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
element, err := it.Nth(people, 1)
// element: {Name: "Bob", Age: 25}, err: nil
```