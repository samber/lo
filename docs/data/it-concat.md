---
name: Concat
slug: concat
sourceRef: it/seq.go#L358
category: it
subCategory: sequence
playUrl:
variantHelpers:
  - it#sequence#concat
similarHelpers:
  - it#sequence#flatten
  - core#slice#concat
position: 160
signatures:
  - "func Concat[T any, I ~func(func(T) bool)](collection ...I) I"
---

Returns a sequence of all the elements in iterators. Concat conserves the order of the elements.

```go
list1 := slices.Values([]int{0, 1, 2})
list2 := slices.Values([]int{3, 4, 5})
list3 := slices.Values([]int{6, 7, 8})

result := Concat(list1, list2, list3)

// result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
```
