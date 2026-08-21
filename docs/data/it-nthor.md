---
name: NthOr
slug: nthor
sourceRef: it/find.go#L472
category: iter
subCategory: find
signatures:
  - "func NthOr[T any, N constraints.Integer](collection iter.Seq[T], nth N, fallback T) T"
playUrl: https://go.dev/play/p/MNweuhpy4Ym
variantHelpers:
  - iter#find#nthor
similarHelpers:
  - core#find#nthor
position: 590
---

Returns the element at index `nth` of collection. If `nth` is out of bounds, returns the fallback value instead of an error.

Will iterate n times through the sequence.

Examples:

```go
// Get element at specific index
numbers := slices.Values([]int{5, 2, 8, 1, 9})
element := it.NthOr(numbers, 2, 42)
// element: 8

// Get first element (index 0)
first := it.NthOr(numbers, 0, 42)
// first: 5

// Get last element
last := it.NthOr(numbers, 4, 42)
// last: 9

// Out of bounds - negative, returns fallback
element = it.NthOr(numbers, -1, 42)
// element: 42 (fallback)

// Out of bounds - too large, returns fallback
element = it.NthOr(numbers, 10, 42)
// element: 42 (fallback)

// With strings
words := slices.Values([]string{"hello", "world", "go", "lang"})
wordElement := it.NthOr(words, 1, "fallback")
// wordElement: "world"

// Out of bounds with string fallback
wordElement = it.NthOr(words, 10, "fallback")
// wordElement: "fallback"

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
fallback := Person{Name: "Default", Age: 0}
personElement := it.NthOr(people, 1, fallback)
// personElement: {Name: "Bob", Age: 25}

// Out of bounds with struct fallback
personElement = it.NthOr(people, 5, fallback)
// personElement: {Name: "Default", Age: 0}

// With different integer types
numbers = slices.Values([]int{1, 2, 3, 4, 5})
element = it.NthOr(numbers, int8(3), 99)
// element: 4
```