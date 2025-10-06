---
name: FromSeqPtrOr
slug: fromseqptror
sourceRef: it/type_manipulation.go#L11
category: it
subCategory: type
signatures:
  - "func FromSeqPtrOr[T any](collection iter.Seq[*T], fallback T) iter.Seq[T]"
variantHelpers:
  - it#type#fromseqptr
similarHelpers: []
position: 242
---

FromSeqPtrOr returns a sequence with the pointer values or the fallback value.

```go
one := 1
var two *int = nil

collection := func(yield func(*int) bool) {
    yield(&one)
    yield(two)
}

values := it.FromSeqPtrOr(collection, 99)
var result []int
for val := range values {
    result = append(result, val)
}
// result contains [1, 99]
```