---
name: ReplaceAll
slug: replaceall
sourceRef: it/seq.go#L699
category: it
subCategory: slice
signatures:
  - "func ReplaceAll[T comparable, I ~func(func(T) bool)](collection I, old, nEw T) I"
variantHelpers:
  - it#slice#replaceall
similarHelpers:
  - core#slice#replaceall
position: 191
---

Returns a sequence with all non-overlapping instances of old replaced by new.

```go
// Basic replacement
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(2)
    yield(3)
}

replaced := it.ReplaceAll(collection, 2, 99)
var result []int
for item := range replaced {
    result = append(result, item)
}
// result: [1, 99, 99, 3]

// With strings - replacing multiple occurrences
strings := func(yield func(string) bool) {
    yield("apple")
    yield("banana")
    yield("apple")
    yield("cherry")
    yield("apple")
}

replacedStrings := it.ReplaceAll(strings, "apple", "orange")
var fruitResult []string
for item := range replacedStrings {
    fruitResult = append(fruitResult, item)
}
// fruitResult: ["orange", "banana", "orange", "cherry", "orange"]

// No matches found
noMatch := it.ReplaceAll(strings, "grape", "kiwi")
var noMatchResult []string
for item := range noMatch {
    noMatchResult = append(noMatchResult, item)
}
// noMatchResult: ["apple", "banana", "apple", "cherry", "apple"] (unchanged)

// Empty collection
empty := func(yield func(int) bool) {
    // no yields
}
emptyReplaced := it.ReplaceAll(empty, 1, 99)
var emptyResult []int
for item := range emptyReplaced {
    emptyResult = append(emptyResult, item)
}
// emptyResult: [] (empty sequence)
```