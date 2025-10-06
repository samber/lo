---
name: HasSuffix
slug: hassuffix
sourceRef: it/find.go#L71
category: it
subCategory: find
signatures:
  - "func HasSuffix[T comparable](collection iter.Seq[T], suffix ...T) bool"
playUrl: ""
variantHelpers:
  - it#find#hassuffix
similarHelpers:
  - core#slice#hassuffix
  - it#find#hasprefix
position: 30
---

Returns true if the collection has the specified suffix. The suffix can be specified as multiple arguments.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
hasSuffix := it.HasSuffix(seq, 3, 4)
// hasSuffix == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("hello")
    _ = yield("world")
}
hasSuffix := it.HasSuffix(seq, "world")
// hasSuffix == true
```

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
hasSuffix := it.HasSuffix(seq, 1, 2)
// hasSuffix == false
```