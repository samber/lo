---
name: Keyify
slug: keyify
sourceRef: it/seq.go#L720
category: it
subCategory: slice
signatures:
  - "func Keyify[T comparable](collection iter.Seq[T]) map[T]struct{}"
playUrl: https://go.dev/play/p/aHOD29_l-rF
variantHelpers: []
similarHelpers:
  - core#slice#keyby
position: 202
---

Keyify returns a map with each unique element of the sequence as a key.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(1)
    yield(3)
    yield(2)
}

keyMap := it.Keyify(collection)
// keyMap contains {1: {}, 2: {}, 3: {}}
```