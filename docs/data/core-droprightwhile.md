---
name: DropRightWhile
slug: droprightwhile
sourceRef: slice.go#L486
category: core
subCategory: slice
playUrl: https://go.dev/play/p/3-n71oEC0Hz
variantHelpers:
  - core#slice#droprightwhile
similarHelpers:
  - core#slice#dropwhile
  - core#slice#drop
  - core#slice#dropright
  - core#slice#dropbyindex
  - core#slice#filterreject
  - core#slice#partitionby
position: 200
signatures:
  - "func DropRightWhile[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice"
---

Drops elements from the end while the predicate returns true.

```go
lo.DropRightWhile([]string{"a", "aa", "aaa", "aa", "aa"}, func(val string) bool {
    return len(val) <= 2
})
// []string{"a", "aa", "aaa"}
```


