---
name: ElementsMatch
slug: elementsmatch
sourceRef: it/intersect.go#L174
category: it
subCategory: intersect
signatures:
  - "func ElementsMatch[T comparable](list1, list2 iter.Seq[T]) bool"
playUrl: "https://go.dev/play/p/yGpdBGaWPCA"
variantHelpers:
  - it#intersect#elementsmatch
similarHelpers:
  - core#slice#elementsmatch
position: 720
---

Returns true if lists contain the same set of elements (including empty set).

If there are duplicate elements, the number of occurrences in each list should match.
The order of elements is not checked.
Will iterate through each sequence before returning and allocate a map large enough to hold all distinct elements.
Long heterogeneous input sequences can cause excessive memory usage.

Examples:

```go
// Lists with same elements in different order
list1 := it.Slice([]int{1, 2, 3, 4, 5})
list2 := it.Slice([]int{5, 4, 3, 2, 1})
match := it.ElementsMatch(list1, list2)
// match: true

// Lists with different elements
list1 = it.Slice([]int{1, 2, 3, 4, 5})
list2 = it.Slice([]int{1, 2, 3, 4, 6})
match = it.ElementsMatch(list1, list2)
// match: false (5 vs 6)

// Lists with duplicates
list1 = it.Slice([]int{1, 2, 2, 3, 4})
list2 = it.Slice([]int{4, 3, 2, 1, 2})
match = it.ElementsMatch(list1, list2)
// match: true (both have two 2's)

// Lists with different number of duplicates
list1 = it.Slice([]int{1, 2, 2, 3, 4})
list2 = it.Slice([]int{4, 3, 2, 1, 1})
match = it.ElementsMatch(list1, list2)
// match: false (list1 has two 2's, list2 has two 1's)

// Empty lists
empty1 := it.Slice([]int{})
empty2 := it.Slice([]int{})
match = it.ElementsMatch(empty1, empty2)
// match: true

// One empty, one not empty
empty := it.Slice([]int{})
nonEmpty := it.Slice([]int{1, 2, 3})
match = it.ElementsMatch(empty, nonEmpty)
// match: false

// String lists
words1 := it.Slice([]string{"hello", "world", "go"})
words2 := it.Slice([]string{"go", "hello", "world"})
match := it.ElementsMatch(words1, words2)
// match: true

words1 = it.Slice([]string{"hello", "world", "go"})
words2 = it.Slice([]string{"go", "hello", "golang"})
match = it.ElementsMatch(words1, words2)
// match: false

// Struct lists
type Person struct {
    Name string
    Age  int
}
people1 := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
people2 := it.Slice([]Person{
    {Name: "Charlie", Age: 35},
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
match := it.ElementsMatch(people1, people2)
// match: true

// Different lengths
list1 = it.Slice([]int{1, 2, 3})
list2 = it.Slice([]int{1, 2, 3, 4})
match = it.ElementsMatch(list1, list2)
// match: false

// Same elements but different counts
list1 = it.Slice([]int{1, 1, 2, 3})
list2 = it.Slice([]int{1, 2, 2, 3})
match = it.ElementsMatch(list1, list2)
// match: false (list1 has two 1's, list2 has two 2's)

// Boolean values
bools1 := it.Slice([]bool{true, false, true})
bools2 := it.Slice([]bool{true, true, false})
match := it.ElementsMatch(bools1, bools2)
// match: true

// Different boolean counts
bools1 = it.Slice([]bool{true, false, true})
bools2 = it.Slice([]bool{true, false, false})
match = it.ElementsMatch(bools1, bools2)
// match: false

// Lists with same single element
list1 = it.Slice([]int{42})
list2 = it.Slice([]int{42})
match = it.ElementsMatch(list1, list2)
// match: true

// Lists with different single elements
list1 = it.Slice([]int{42})
list2 = it.Slice([]int{43})
match = it.ElementsMatch(list1, list2)
// match: false
```