---
name: RuneLength
slug: runelength
sourceRef: string.go#L160
category: core
subCategory: string
playUrl: https://go.dev/play/p/tuhgW_lWY8l
variantHelpers:
  - core#string#runelength
similarHelpers:
  - core#string#substring
  - core#string#chunkstring
  - core#string#words
  - core#string#capitalize
  - core#string#ellipsis
position: 30
signatures:
  - "func RuneLength(str string) int"
---

Returns the number of runes (Unicode code points) in a string.

```go
lo.RuneLength("hellô")
// 5
```


