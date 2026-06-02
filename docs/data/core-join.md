---
name: Join
slug: join
sourceRed: slice.go#L1348
category: core
subCategory: slice
playUrl: https://go.dev/play/p/uEwlghJIHmN
signatures:
 - "func Join[T any, Slice []T](collection Slice, separator any) string"
---

Joins every collection item into a string, using the specified separator. If the separator is nil, it is considered as an empty string.

```go
result1 := lo.Join([]int{1, 2, 3, 4}, " - "))
// "1 - 2 - 3 - 4"

result1 := lo.Join([]string{"a", "b", "c", "d"}, 2))
// "a2b2c2d"
```
