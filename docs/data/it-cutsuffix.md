---
name: CutSuffix
slug: cutsuffix
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func CutSuffix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (before I, found bool)"
variantHelpers: []
similarHelpers:
  - core#string#cutsuffix
position: 261
---

CutSuffix returns collection without the provided ending suffix and reports whether it found the suffix.
If collection doesn't end with suffix, CutSuffix returns collection, false.
If suffix is empty, CutSuffix returns collection, true.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}

before, found := it.CutSuffix(collection, []int{3, 4})
var result []int
for item := range before {
    result = append(result, item)
}
// result contains [1, 2], found is true
```