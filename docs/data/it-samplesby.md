---
name: SamplesBy
slug: samplesby
sourceRef: it/find.go#L474
category: it
subCategory: find
signatures:
  - "func SamplesBy[T any, I ~func(func(T) bool)](collection I, count int, randomIntGenerator func(int) int) I"
playUrl: "https://go.dev/play/p/7YsLP1-zx"
variantHelpers:
  - it#find#samplesby
similarHelpers:
  - core#slice#samplesby
  - it#find#sample
  - it#find#samples
  - it#find#sampleby
position: 620
---

Returns N random unique items from collection, using randomIntGenerator as the random index generator.

Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
Long input sequences can cause excessive memory usage.

Examples:

```go
import (
    "math/rand"
    "time"
)

// Use default random generator with seed
rng := rand.New(rand.NewSource(time.Now().UnixNano()))

// Get 3 random unique items with custom random generator
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
samples := it.SamplesBy(numbers, 3, rng.Intn)
// samples: sequence of 3 random unique numbers

// Use deterministic random generator for testing
deterministicRng := rand.New(rand.NewSource(42))
samples := it.SamplesBy(numbers, 3, deterministicRng.Intn)
// samples: predictable sequence of 3 unique numbers

// With strings
words := it.Slice([]string{"apple", "banana", "cherry", "date", "elderberry"})
samples := it.SamplesBy(words, 2, rng.Intn)
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
samples := it.SamplesBy(people, 3, rng.Intn)
// samples: sequence of 3 random unique people

// Custom random function that biases towards certain items
biasedRandom := func(max int) int {
    // Bias towards first half of collection
    return rng.Intn(max/2 + 1)
}
samples := it.SamplesBy(numbers, 3, biasedRandom)
// samples: sequence of 3 random unique numbers, biased towards lower indices

// With zero-based modulo function (wraps around)
moduloRandom := func(max int) int {
    return rng.Intn(max*3) % max
}
samples := it.SamplesBy(numbers, 3, moduloRandom)
// samples: sequence of 3 random unique numbers

// Test with deterministic function
deterministicFunc := func(max int) int {
    return (max - 1) / 2 // Always return middle index
}
samples := it.SamplesBy(numbers, 1, deterministicFunc)
// samples: sequence with single element from middle
```