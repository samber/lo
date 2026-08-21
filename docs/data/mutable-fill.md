---
name: Fill
slug: fill
sourceRef: mutable/slice_example_test.go#L72
category: mutable
subCategory: slice
signatures:
  - "func Fill[T any, Slice ~[]T](collection Slice, initial T)"
playUrl: https://go.dev/play/p/VwR34GzqEub
variantHelpers:
  - mutable#slice#fill
similarHelpers:
  - core#slice#fill
  - core#slice#repeat
  - core#slice#times
position: 60
---

Fills a slice with clones of the specified initial value, returning a new slice.

```go
type foo struct{ bar string }
func (f foo) Clone() foo {
    return foo{f.bar}
}

slice := lo.Fill(make([]foo, 5), foo{"a"})
// []foo{{"a"}, {"a"}, {"a"}, {"a"}, {"a"}}

slice = lo.Fill(make([]foo, 3), foo{"b"})
// []foo{{"b"}, {"b"}, {"b"}}

slice = lo.Fill([]foo{{"x"}, {"y"}, {"z"}}, foo{"c"})
// []foo{{"c"}, {"c"}, {"c"}}
```