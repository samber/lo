---
name: MinIndexBy
slug: minindexby
sourceRef: it/find.go#L258
category: it
subCategory: find
signatures:
  - "func MinIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int)"
playUrl: "https://go.dev/play/p/6DxQUQ0-zc"
variantHelpers:
  - it#find#minindexby
similarHelpers:
  - core#slice#minindexby
position: 470
---

Searches the minimum value of a collection using a comparison function and returns both the value and its index.

If several values are equal to the smallest value, returns the first such value.
Returns (zero value, -1) when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
// Find the minimum string by length and its index
words := it.Slice([]string{"apple", "hi", "banana", "ok"})
value, index := it.MinIndexBy(words, func(a, b string) bool {
    return len(a) < len(b)
})
// value: "hi", index: 1

// Find the minimum person by age and its index
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
value, index := it.MinIndexBy(people, func(a, b Person) bool {
    return a.Age < b.Age
})
// value: {Name: "Bob", Age: 25}, index: 1
```