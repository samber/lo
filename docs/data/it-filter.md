---
name: Filter
slug: filter
sourceRef: it/seq.go#L33
category: it
subCategory: sequence
signatures:
  - "func Filter[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
  - "func FilterI[T any, I ~func(func(T) bool)](collection I, predicate func(item T, index int) bool) I"
playUrl: ""
variantHelpers:
  - it#sequence#filter
  - it#sequence#filteri
similarHelpers:
  - core#slice#filter
  - core#slice#filteri
  - it#sequence#reject
position: 10
---

Returns a sequence of all elements for which the predicate function returns true.

### Filter

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
filtered := it.Filter(seq, func(x int) bool {
    return x%2 == 0
})
var result []int
for v := range filtered {
    result = append(result, v)
}
// result contains 2, 4 (even numbers)
```

### FilterI

FilterI iterates over elements of collection, returning a sequence of all elements predicate returns true for. The predicate function includes the index.

```go
result := it.FilterI(it.Range(1, 6), func(item int, index int) bool {
    return item%2 == 0 && index > 1
})
var filtered []int
for v := range result {
    filtered = append(filtered, v)
}
// filtered contains [4, 6]
```