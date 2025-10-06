---
name: FindDuplicatesBy
slug: findduplicatesby
sourceRef: it/find.go#L200
category: it
subCategory: find
signatures:
  - "func FindDuplicatesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I"
playUrl: ""
variantHelpers:
  - it#find#findduplicatesby
similarHelpers:
  - core#slice#findduplicatesby
position: 640
---

Returns a sequence with the first occurrence of each duplicated element in the collection, based on a transform function.

The order of result values is determined by the order duplicates occur in the sequence. A transform function is
invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
Will allocate a map large enough to hold all distinct transformed elements.
Long heterogeneous input sequences can cause excessive memory usage.

Examples:

```go
// Find duplicate people by age
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 30},  // Same age as Alice - Alice is returned
    {Name: "Diana", Age: 30},   // Same age as Alice - already marked as duplicate
    {Name: "Eve", Age: 25},     // Same age as Bob - Bob is returned
})
duplicates := it.FindDuplicatesBy(people, func(p Person) int { return p.Age })
// duplicates: sequence with Alice (age 30) and Bob (age 25)

// Find duplicate strings by length
words := it.Slice([]string{"hello", "world", "hi", "go", "bye", "yes"})
duplicates := it.FindDuplicatesBy(words, func(s string) int { return len(s) })
// duplicates: sequence with "hello" (length 5, also "world" has length 5)
// and "hi" (length 2, also "go", "yes" have length 2)

// Find duplicate items by first letter
items := it.Slice([]string{"apple", "apricot", "banana", "blueberry", "cherry", "cranberry"})
duplicates := it.FindDuplicatesBy(items, func(s string) byte { return s[0] })
// duplicates: sequence with "apple" (starts with 'a', also "apricot"),
// "banana" (starts with 'b', also "blueberry"),
// "cherry" (starts with 'c', also "cranberry")

// Find duplicate numbers by modulo
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
duplicates := it.FindDuplicatesBy(numbers, func(n int) int { return n % 3 })
// duplicates: sequence with 1, 2, 3 (remainders 1, 2, 0 appear multiple times)

// Find duplicate structs by composite key
type Order struct {
    CustomerID string
    ProductID  string
}
orders := it.Slice([]Order{
    {CustomerID: "A", ProductID: "1"},
    {CustomerID: "A", ProductID: "2"},
    {CustomerID: "B", ProductID: "1"},  // Same customer as first
    {CustomerID: "C", ProductID: "3"},
    {CustomerID: "A", ProductID: "3"},  // Same customer as first two
})
duplicates := it.FindDuplicatesBy(orders, func(o Order) string { return o.CustomerID })
// duplicates: sequence with first order {CustomerID: "A", ProductID: "1"}

// Find duplicate items by case-insensitive comparison
words := it.Slice([]string{"Hello", "hello", "WORLD", "world", "Go", "GO", "go"})
duplicates := it.FindDuplicatesBy(words, func(s string) string { return strings.ToLower(s) })
// duplicates: sequence with "Hello", "WORLD", "Go" (first occurrences of each case-insensitive duplicate)

// Find duplicate dates by year-month
import "time"
dates := it.Slice([]time.Time{
    time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 1, 20, 0, 0, 0, 0, time.UTC),  // Same month as first
    time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC),
    time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC),   // Different year
    time.Date(2023, 2, 25, 0, 0, 0, 0, time.UTC),  // Same month as third
})
duplicates := it.FindDuplicatesBy(dates, func(t time.Time) string {
    return fmt.Sprintf("%d-%02d", t.Year(), t.Month())
})
// duplicates: sequence with first January date and first February date

// Find duplicate emails by domain
type Email struct {
    Address string
}
emails := it.Slice([]Email{
    {Address: "user1@example.com"},
    {Address: "user2@example.com"},  // Same domain as first
    {Address: "user3@gmail.com"},
    {Address: "user4@example.com"},  // Same domain as first two
    {Address: "user5@yahoo.com"},
})
duplicates := it.FindDuplicatesBy(emails, func(e Email) string {
    parts := strings.Split(e.Address, "@")
    if len(parts) > 1 {
        return parts[1]
    }
    return ""
})
// duplicates: sequence with first email {Address: "user1@example.com"}
```