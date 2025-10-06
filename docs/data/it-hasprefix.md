---
name: HasPrefix
slug: hasprefix
sourceRef: it/find.go#L49
category: it
subCategory: find
signatures:
  - "func HasPrefix[T comparable](collection iter.Seq[T], prefix ...T) bool"
playUrl: ""
variantHelpers:
  - it#find#hasprefix
similarHelpers:
  - core#slice#hasprefix
  - it#find#hassuffix
position: 20
---

Returns true if the collection has the specified prefix. The prefix can be specified as multiple arguments.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
hasPrefix := it.HasPrefix(seq, 1, 2)
// hasPrefix == true
```

```go
seq := func(yield func(string) bool) {
    _ = yield("hello")
    _ = yield("world")
}
hasPrefix := it.HasPrefix(seq, "hello")
// hasPrefix == true
```

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
hasPrefix := it.HasPrefix(seq, 2, 3)
// hasPrefix == false
```