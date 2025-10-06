---
name: Fill
slug: fill
sourceRef: slice.go#L338
category: core
subCategory: slice
playUrl: https://go.dev/play/p/VwR34GzqEub
variantHelpers:
  - core#slice#fill
similarHelpers:
  - core#slice#repeat
  - core#slice#slice
  - core#slice#flatten
  - core#slice#chunk
  - core#slice#chunkentries
  - core#slice#interleave
  - core#slice#reverse
  - core#slice#shuffle
  - core#slice#sample
  - core#slice#drop
  - core#slice#dropwhile
  - core#slice#dropright
  - core#slice#droprightwhile
position: 200
signatures:
  - "func Fill[T Clonable[T], Slice ~[]T](collection Slice, initial T) Slice"
---

Fills a slice with clones of the provided initial value.

```go
type foo struct{ bar string }
func (f foo) Clone() foo {
    return foo{f.bar}
}

lo.Fill([]foo{{"a"}, {"a"}}, foo{"b"})
// []foo{{"b"}, {"b"}}
```


