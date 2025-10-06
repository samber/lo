---
name: MaxIndexBy
slug: maxindexby
sourceRef: it/find.go#L326
category: it
subCategory: find
signatures:
  - "func MaxIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int)"
playUrl: ""
variantHelpers:
  - it#find#maxindexby
similarHelpers:
  - core#slice#maxindexby
position: 490
---

Searches the maximum value of a collection using a comparison function and returns both the value and its index.

If several values are equal to the greatest value, returns the first such value.
Returns (zero value, -1) when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
// Find the maximum string by length and its index
words := it.Slice([]string{"apple", "hi", "banana", "xylophone"})
value, index := it.MaxIndexBy(words, func(a, b string) bool {
    return len(a) > len(b)
})
// value: "xylophone", index: 3

// Find the maximum person by age and its index
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
value, index := it.MaxIndexBy(people, func(a, b Person) bool {
    return a.Age > b.Age
})
// value: {Name: "Charlie", Age: 35}, index: 2

// Find the maximum number by absolute value and its index
numbers := it.Slice([]int{-5, 2, -8, 1})
value, index := it.MaxIndexBy(numbers, func(a, b int) bool {
    return abs(a) > abs(b)
})
// value: -8, index: 2
```