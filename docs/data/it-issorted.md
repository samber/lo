---
name: IsSorted
slug: issorted
sourceRef: it/seq.go#L720
category: it
subCategory: slice
signatures:
  - "func IsSorted[T constraints.Ordered](collection iter.Seq[T]) bool"
variantHelpers: []
similarHelpers:
  - core#slice#issorted
position: 200
---

IsSorted checks if a sequence is sorted.

```go
sorted := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}
unsorted := func(yield func(int) bool) {
    yield(1)
    yield(3)
    yield(2)
    yield(4)
}

fmt.Println(it.IsSorted(sorted))    // true
fmt.Println(it.IsSorted(unsorted))  // false
```