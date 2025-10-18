---
name: FindDuplicates
slug: findduplicates
sourceRef: it/find.go#L196
category: it
subCategory: find
signatures:
  - "func FindDuplicates[T comparable, I ~func(func(T) bool)](collection I) I"
playUrl: "https://go.dev/play/p/1YsRLPl-wx"
variantHelpers:
  - it#find#findduplicates
similarHelpers:
  - core#slice#findduplicates
  - it#find#finduniques
  - it#find#findduplicatesby
position: 90
---

Returns the first occurrence of each duplicated element in the collection (elements that appear more than once).

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(4)
    _ = yield(4)
}
dupSeq := it.FindDuplicates(seq)
var result []int
for v := range dupSeq {
    result = append(result, v)
}
// result contains 2, 4 (first occurrence of each duplicated element)
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("apple")
    _ = yield("cherry")
    _ = yield("banana")
}
dupSeq := it.FindDuplicates(seq)
var result []string
for v := range dupSeq {
    result = append(result, v)
}
// result contains "apple", "banana" (duplicated elements)
```