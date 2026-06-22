---
name: SnakeCase
slug: snakecase
sourceRef: string.go#L200
category: core
subCategory: string
playUrl: https://go.dev/play/p/ziB0V89IeVH
variantHelpers:
  - core#string#snakecase
similarHelpers:
  - core#string#pascalcase
  - core#string#camelcase
  - core#string#kebabcase
  - core#string#capitalize
  - core#string#words
position: 70
signatures:
  - "func SnakeCase(str string) string"
---

Converts a string to snake_case.

```go
lo.SnakeCase("HelloWorld")
// "hello_world"
```


