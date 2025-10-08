---
name: IsNotEmpty
slug: isnotempty
sourceRef: it/type_manipulation.go#L59
category: it
subCategory: condition
signatures:
  - "func IsNotEmpty[T any](collection iter.Seq[T]) bool"
playUrl: 
variantHelpers:
  - it#condition#isempty
similarHelpers:
  - core#slice#isnotempty
  - core#slice#isempty
position: 10
---

Returns true if the collection is not empty, false otherwise.

```go
result1 := it.IsNotEmpty(it.Range(1, 5))
// true

result2 := it.IsNotEmpty(it.Empty[int]())
// false
```