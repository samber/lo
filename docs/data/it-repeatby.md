---
name: RepeatBy
slug: repeatby
sourceRef: it/seq.go#L388
category: iter
subCategory: sequence
signatures:
  - "func RepeatBy[T any](count int, callback func(index int) T) iter.Seq[T]"
playUrl: https://go.dev/play/p/i7BuZQBcUzZ
variantHelpers:
  - iter#sequence#repeatby
similarHelpers:
  - core#slice#repeat
  - core#slice#times
  - iter#sequence#times
position: 130
---

Builds a sequence with values returned by N calls of callback.

```go
result := it.RepeatBy(3, func(index int) string {
    return fmt.Sprintf("item-%d", index+1)
})
var output []string
for item := range result {
    output = append(output, item)
}
// output contains ["item-1", "item-2", "item-3"]

result2 := it.RepeatBy(5, func(index int) int {
    return index * 2
})
var output2 []int
for item := range result2 {
    output2 = append(output2, item)
}
// output2 contains [0, 2, 4, 6, 8]
```