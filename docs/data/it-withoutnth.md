---
name: WithoutNth
slug: withoutnth
sourceRef: it/intersect.go#L253
category: iter
subCategory: intersect
signatures:
  - "func WithoutNth[T comparable, I ~func(func(T) bool)](collection I, nths ...int) I"
playUrl: "https://go.dev/play/p/KGE7Lpsk18P"
variantHelpers:
  - iter#intersect#withoutnth
similarHelpers:
  - core#intersect#withoutnth
position: 710
---

Returns a sequence excluding the elements at the specified indices.

Will allocate a map large enough to hold all distinct indices.

Examples:

```go
// Exclude elements at specific indices
numbers := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
filtered := it.WithoutNth(numbers, 2, 5, 8)
// filtered: sequence with 1, 2, 4, 5, 7, 8, 10
// excludes elements at indices 2 (3), 5 (6), and 8 (9)

// Exclude single element
words := slices.Values([]string{"hello", "world", "go", "lang", "awesome"})
filteredStr := it.WithoutNth(words, 1)
// filteredStr: sequence with "hello", "go", "lang", "awesome"
// excludes "world" at index 1

// Exclude first element
numbers = slices.Values([]int{10, 20, 30, 40, 50})
filtered = it.WithoutNth(numbers, 0)
// filtered: sequence with 20, 30, 40, 50
// excludes 10 at index 0

// Exclude last element
numbers = slices.Values([]int{10, 20, 30, 40, 50})
filtered = it.WithoutNth(numbers, 4)
// filtered: sequence with 10, 20, 30, 40
// excludes 50 at index 4

// Exclude multiple elements including duplicates
words = slices.Values([]string{"a", "b", "c", "d", "e", "f", "g"})
filteredStr = it.WithoutNth(words, 1, 3, 1, 5)
// filteredStr: sequence with "a", "c", "e", "g"
// excludes elements at indices 1 (b), 3 (d), and 5 (f)
// index 1 appears twice but element at index 1 is only excluded once

// Exclude with negative indices (out of bounds, no effect)
numbers = slices.Values([]int{1, 2, 3, 4, 5})
filtered = it.WithoutNth(numbers, -1, 2)
// filtered: sequence with 1, 2, 4, 5
// excludes element at index 2 (3), ignores -1

// Exclude with indices larger than collection (out of bounds, no effect)
numbers = slices.Values([]int{1, 2, 3, 4, 5})
filtered = it.WithoutNth(numbers, 10, 2)
// filtered: sequence with 1, 2, 4, 5
// excludes element at index 2 (3), ignores 10

// Exclude all elements
numbers = slices.Values([]int{1, 2, 3, 4, 5})
filtered = it.WithoutNth(numbers, 0, 1, 2, 3, 4)
// filtered: empty sequence

// Exclude no indices (returns original)
numbers = slices.Values([]int{1, 2, 3, 4, 5})
filtered = it.WithoutNth(numbers)
// filtered: sequence with 1, 2, 3, 4, 5 (unchanged)

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
    {Name: "Diana", Age: 28},
    {Name: "Eve", Age: 32},
})
filteredPerson := it.WithoutNth(people, 1, 3)
// filteredPerson: sequence with Alice, Charlie, Eve
// excludes Bob at index 1 and Diana at index 3

// With mixed valid and invalid indices
words = slices.Values([]string{"first", "second", "third", "fourth"})
filteredStr = it.WithoutNth(words, -1, 1, 10, 2)
// filteredStr: sequence with "first", "fourth"
// excludes "second" at index 1 and "third" at index 2
// ignores -1 and 10 as they are out of bounds

// Exclude from empty collection
empty := slices.Values([]int{})
filtered = it.WithoutNth(empty, 0, 1, 2)
// filtered: empty sequence
```