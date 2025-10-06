---
name: Samples
slug: samples
sourceRef: it/find.go#L467
category: it
subCategory: find
signatures:
  - "func Samples[T any, I ~func(func(T) bool)](collection I, count int) I"
playUrl: ""
variantHelpers:
  - it#find#samples
similarHelpers:
  - core#slice#samples
  - it#find#sample
  - it#find#sampleby
  - it#find#samplesby
position: 610
---

Returns N random unique items from collection.

Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
Long input sequences can cause excessive memory usage.

Examples:

```go
// Get 3 random unique items from collection
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
samples := it.Samples(numbers, 3)
// samples: sequence of 3 random unique numbers from 1-10

// Get all items if count equals collection size
numbers := it.Slice([]int{1, 2, 3, 4, 5})
samples := it.Samples(numbers, 5)
// samples: sequence containing all 5 numbers in random order

// Get fewer items than collection size
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
samples := it.Samples(numbers, 3)
// samples: sequence of 3 random unique numbers

// With strings
words := it.Slice([]string{"apple", "banana", "cherry", "date", "elderberry"})
samples := it.Samples(words, 2)
// samples: sequence of 2 random unique words

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
    {Name: "Diana", Age: 28},
    {Name: "Eve", Age: 32},
})
samples := it.Samples(people, 3)
// samples: sequence of 3 random unique people

// Count larger than collection size - returns all items in random order
numbers := it.Slice([]int{1, 2, 3})
samples := it.Samples(numbers, 10)
// samples: sequence of all 3 numbers in random order

// Zero count - returns empty sequence
numbers := it.Slice([]int{1, 2, 3, 4, 5})
samples := it.Samples(numbers, 0)
// samples: empty sequence

// Negative count - returns empty sequence
numbers := it.Slice([]int{1, 2, 3, 4, 5})
samples := it.Samples(numbers, -1)
// samples: empty sequence
```