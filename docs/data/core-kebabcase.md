---
name: KebabCase
slug: kebabcase
sourceRef: string.go#L190
category: core
subCategory: string
playUrl: https://go.dev/play/p/96gT_WZnTVP
variantHelpers:
  - core#string#kebabcase
similarHelpers:
  - core#string#pascalcase
  - core#string#camelcase
  - core#string#snakecase
  - core#string#capitalize
  - core#string#words
position: 60
signatures:
  - "func KebabCase(str string) string"
---

Converts a string to kebab-case.

```go
lo.KebabCase("helloWorld")
// "hello-world"
```


