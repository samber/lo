---
name: GroupBy
slug: groupby
sourceRef: it/seq.go#L244
category: it
subCategory: sequence
signatures:
  - "func GroupBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U][]T"
playUrl: https://go.dev/play/p/2TnGK6-zs
variantHelpers:
  - it#sequence#groupby
similarHelpers:
  - core#slice#groupby
  - it#sequence#partitionby
  - it#sequence#groupbymap
position: 80
---

Returns an object composed of keys generated from running each element of collection through a transform function. The value of each key is an array of elements responsible for generating the key.

Examples:

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("apricot")
    _ = yield("blueberry")
}
grouped := it.GroupBy(seq, func(s string) string {
    return string(s[0]) // group by first letter
})
// grouped contains map with keys: "a": ["apple", "apricot"], "b": ["banana", "blueberry"]
```