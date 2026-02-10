---
name: Ellipsis
slug: ellipsis
sourceRef: string.go#L235
category: core
subCategory: string
playUrl: https://go.dev/play/p/qE93rgqe1TW
variantHelpers:
  - core#string#ellipsis
  # typo / deprecated
  - core#string#elipse
similarHelpers:
  - core#string#substring
  - core#string#runelength
  - core#string#trim
  - core#string#capitalize
position: 100
signatures:
  - "func Ellipsis(str string, length int) string"
---

Trims and truncates a string to the specified length in runes (Unicode code points) and appends an ellipsis if truncated. Multi-byte characters such as emoji or CJK ideographs are never split in the middle.

```go
lo.Ellipsis("  Lorem Ipsum  ", 5)
// "Lo..."

str = lo.Ellipsis("Lorem Ipsum", 100)
// "Lorem Ipsum"

str = lo.Ellipsis("Lorem Ipsum", 3)
// "..."

str = lo.Ellipsis("hello 世界! 你好", 8)
// "hello..."

str = lo.Ellipsis("🏠🐶🐱🌟", 4)
// "🏠🐶🐱🌟"
```


