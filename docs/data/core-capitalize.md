---
name: Capitalize
slug: capitalize
sourceRef: string.go#L227
category: core
subCategory: string
playUrl: https://go.dev/play/p/uLTZZQXqnsa
variantHelpers:
  - core#string#capitalize
similarHelpers:
  - core#string#pascalcase
  - core#string#camelcase
  - core#string#snakecase
  - core#string#kebabcase
  - core#string#words
  - core#string#runelength
position: 90
signatures:
  - "func Capitalize(str string) string"
---

Converts the first character to uppercase and the remaining to lowercase.

```go
lo.Capitalize("heLLO")
// "Hello"
```


