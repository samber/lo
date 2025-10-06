---
name: FindUniquesBy
slug: finduniquesby
sourceRef: it/find.go#L163
category: it
subCategory: find
signatures:
  - "func FindUniquesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I"
playUrl: ""
variantHelpers:
  - it#find#finduniquesby
similarHelpers:
  - core#slice#finduniquesby
position: 630
---

Returns a sequence with all the elements that appear in the collection only once, based on a transform function.

The order of result values is determined by the order they occur in the collection. A transform function is
invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
Will iterate through the entire sequence before yielding and allocate a map large enough to hold all distinct transformed elements.
Long heterogeneous input sequences can cause excessive memory usage.

Examples:

```go
// Find unique people by age
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 30},  // Same age as Alice, so Alice is not unique
    {Name: "Diana", Age: 28},
})
uniques := it.FindUniquesBy(people, func(p Person) int { return p.Age })
// uniques: sequence with Bob (age 25) and Diana (age 28)

// Find unique strings by length
words := it.Slice([]string{"hello", "world", "hi", "go", "bye"})
uniques := it.FindUniquesBy(words, func(s string) int { return len(s) })
// uniques: sequence with words of unique lengths
// "hello" (5), "world" (5) - not unique due to same length
// "hi" (2), "go" (2), "bye" (3) - "bye" is unique (length 3)

// Find unique items by first letter
items := it.Slice([]string{"apple", "apricot", "banana", "blueberry", "cherry"})
uniques := it.FindUniquesBy(items, func(s string) byte { return s[0] })
// uniques: sequence with "cherry" (only word starting with 'c')

// Find unique numbers by modulo
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
uniques := it.FindUniquesBy(numbers, func(n int) int { return n % 3 })
// uniques: sequence with numbers that have unique remainders when divided by 3

// Find unique structs by composite key
type Order struct {
    CustomerID string
    ProductID  string
}
orders := it.Slice([]Order{
    {CustomerID: "A", ProductID: "1"},
    {CustomerID: "A", ProductID: "2"},
    {CustomerID: "B", ProductID: "1"},  // Same customer as first orders
    {CustomerID: "C", ProductID: "3"},
})
uniques := it.FindUniquesBy(orders, func(o Order) string { return o.CustomerID })
// uniques: sequence with Order{CustomerID: "C", ProductID: "3"} only

// Find unique items by case-insensitive comparison
words := it.Slice([]string{"Hello", "hello", "WORLD", "world", "Go"})
uniques := it.FindUniquesBy(words, func(s string) string { return strings.ToLower(s) })
// uniques: sequence with "Go" only (others have case-insensitive duplicates)

// Find unique dates by year-month
import "time"
dates := it.Slice([]time.Time{
    time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 1, 20, 0, 0, 0, 0, time.UTC),  // Same month as first
    time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC),
    time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),   // Different year
})
uniques := it.FindUniquesBy(dates, func(t time.Time) string {
    return fmt.Sprintf("%d-%02d", t.Year(), t.Month())
})
// uniques: sequence with dates from unique year-month combinations
```