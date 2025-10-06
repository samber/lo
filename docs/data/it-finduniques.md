---
name: FindUniques
slug: finduniques
sourceRef: it/find.go#L159
category: it
subCategory: find
signatures:
  - "func FindUniques[T comparable, I ~func(func(T) bool)](collection I) I"
playUrl: ""
variantHelpers:
  - it#find#finduniques
similarHelpers:
  - core#slice#finduniques
  - it#find#findduplicates
  - it#find#finduniquesby
position: 80
---

Returns a sequence with elements that appear only once in the original collection (duplicates are removed).

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(4)
}
uniqueSeq := it.FindUniques(seq)
var result []int
for v := range uniqueSeq {
    result = append(result, v)
}
// result contains 1, 3 (elements that appear only once)
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("apple")
    _ = yield("cherry")
}
uniqueSeq := it.FindUniques(seq)
var result []string
for v := range uniqueSeq {
    result = append(result, v)
}
// result contains "banana", "cherry" (unique elements)
```