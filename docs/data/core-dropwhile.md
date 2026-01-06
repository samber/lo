---
name: DropWhile
slug: dropwhile
sourceRef: slice.go#L472
category: core
subCategory: slice
playUrl: https://go.dev/play/p/7gBPYw2IK16
variantHelpers:
  - core#slice#dropwhile
similarHelpers:
  - core#slice#droprightwhile
  - core#slice#drop
  - core#slice#dropright
  - core#slice#dropbyindex
  - core#slice#filterreject
  - core#slice#partitionby
  - core#slice#takewhile
position: 190
signatures:
  - "func DropWhile[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice"
---

Drops elements from the beginning while the predicate returns true.

```go
lo.DropWhile([]string{"a", "aa", "aaa", "aa", "aa"}, func(val string) bool {
    return len(val) <= 2
})
// []string{"aaa", "aa", "aa"}
```


