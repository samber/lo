---
name: DropByIndex
slug: dropbyindex
sourceRef: it/seq.go#L745
category: iter
subCategory: sequence
signatures:
  - "func DropByIndex[T any, I ~func(func(T) bool)](collection I, indexes ...int) I"
playUrl: https://go.dev/play/p/vPbrZYgiU4q
variantHelpers:
  - iter#sequence#drop
similarHelpers:
  - core#slice#dropbyindex
  - core#intersect#withoutnth
position: 55
---

Removes elements from a collection at the specified indexes.

```go
result := it.DropByIndex(it.RangeFrom(1, 6), 1, 3)
// [1, 3, 5]
```