---
name: CutPrefix
slug: cutprefix
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func CutPrefix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (after I, found bool)"
variantHelpers: []
similarHelpers:
  - core#string#cutprefix
position: 260
---

CutPrefix returns collection without the provided leading prefix and reports whether it found the prefix.
If collection doesn't start with prefix, CutPrefix returns collection, false.
If prefix is empty, CutPrefix returns collection, true.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}

after, found := it.CutPrefix(collection, []int{1, 2})
var result []int
for item := range after {
    result = append(result, item)
}
// result contains [3, 4], found is true

after2, found2 := it.CutPrefix(collection, []int{9, 10})
var result2 []int
for item := range after2 {
    result2 = append(result2, item)
}
// result2 contains [1, 2, 3, 4], found2 is false
```