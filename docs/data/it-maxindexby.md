---
name: MaxIndexBy
slug: maxindexby
sourceRef: it/find.go#L350
category: iter
subCategory: find
signatures:
  - "func MaxIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int)"
playUrl: https://go.dev/play/p/MXyE6BTILjx
variantHelpers:
  - iter#find#maxindexby
similarHelpers:
  - core#find#maxindexby
position: 490
---

Searches the maximum value of a collection using a comparison function and returns both the value and its index.

If several values are equal to the greatest value, returns the first such value.
Returns (zero value, -1) when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
// Find the maximum string by length and its index
words := slices.Values([]string{"apple", "hi", "banana", "xylophone"})
value, index := it.MaxIndexBy(words, func(a, b string) bool {
    return len(a) > len(b)
})
// value: "xylophone", index: 3

// Find the maximum person by age and its index
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
personValue, personIndex := it.MaxIndexBy(people, func(a, b Person) bool {
    return a.Age > b.Age
})
// personValue: {Name: "Charlie", Age: 35}, personIndex: 2

// Find the maximum number by absolute value and its index
numbers := slices.Values([]int{-5, 2, -8, 1})
absValue, absIndex := it.MaxIndexBy(numbers, func(a, b int) bool {
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -b
    }
    return a > b
})
// absValue: -8, absIndex: 2
```