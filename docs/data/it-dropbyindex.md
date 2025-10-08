---
name: DropByIndex
slug: dropbyindex
sourceRef: it/seq.go#L581
category: it
subCategory: sequence
signatures:
  - "func DropByIndex[T any, I ~func(func(T) bool)](collection I, indexes ...int) I"
playUrl:
variantHelpers:
  - it#slice#drop
similarHelpers:
  - core#slice#dropbyindex
  - core#slice#withoutnth
position: 55
---

Removes elements from a collection at the specified indexes.

```go
result := it.DropByIndex(it.Range(1, 6), 1, 3)
// [1, 3, 5]
```