---
name: Empty
slug: empty
sourceRef: it/type_manipulation.go#L44
category: it
subCategory: type
signatures:
  - "func Empty[T any]() iter.Seq[T]"
playUrl: "https://go.dev/play/p/E5fF1hH8Bc3"
variantHelpers:
  - it#type#empty
similarHelpers:
  - it#type#isempty
  - it#type#isnotempty
position: 0
---

Returns an empty sequence of the specified type.

Examples:

```go
emptySeq := it.Empty[int]()
count := 0
for range emptySeq {
    count++
}
// count == 0
```

```go
emptySeq := it.Empty[string]()
var result []string
for v := range emptySeq {
    result = append(result, v)
}
// result is empty slice
```