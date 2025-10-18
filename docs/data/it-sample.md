---
name: Sample
slug: sample
sourceRef: it/find.go#L455
category: it
subCategory: find
signatures:
  - "func Sample[T any](collection iter.Seq[T]) T"
playUrl: "https://go.dev/play/p/4VpIM8-zu"
variantHelpers:
  - it#find#sample
similarHelpers:
  - core#slice#sample
  - it#find#samples
  - it#find#sampleBy
  - it#find#samplesBy
position: 160
---

Returns a random item from collection.

Example:

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
item := it.Sample(seq)
// item is randomly one of: "apple", "banana", "cherry"

// Example with integers
numbers := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(40)
}
randomNum := it.Sample(numbers)
// randomNum is randomly one of: 10, 20, 30, 40

// Example with empty sequence - returns zero value
empty := func(yield func(string) bool) {
    // no yields
}
emptyResult := it.Sample(empty)
// emptyResult: "" (zero value for string)

// Example with single item
single := func(yield func(int) bool) {
    _ = yield(42)
}
singleResult := it.Sample(single)
// singleResult: 42 (always returns 42 since it's the only option)
```
