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

Trims and truncates a string to the specified byte length and appends an ellipsis if truncated.

```go
lo.Ellipsis("  Lorem Ipsum  ", 5)
// "Lo..."

str = lo.Ellipsis("Lorem Ipsum", 100)
// "Lorem Ipsum"

str = lo.Ellipsis("Lorem Ipsum", 3)
// "..."
```


