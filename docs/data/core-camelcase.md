---
name: CamelCase
slug: camelcase
sourceRef: string.go#L176
category: core
subCategory: string
playUrl: https://go.dev/play/p/Go6aKwUiq59
variantHelpers:
  - core#string#camelcase
similarHelpers:
  - core#string#pascalcase
  - core#string#snakecase
  - core#string#kebabcase
  - core#string#capitalize
  - core#string#words
position: 50
signatures:
  - "func CamelCase(str string) string"
---

Converts a string to camelCase.

```go
lo.CamelCase("hello_world")
// "helloWorld"
```


