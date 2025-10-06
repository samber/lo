---
name: Words
slug: words
sourceRef: string.go#L210
category: core
subCategory: string
playUrl: https://go.dev/play/p/-f3VIQqiaVw
variantHelpers:
  - core#string#words
similarHelpers:
  - core#string#pascalcase
  - core#string#camelcase
  - core#string#kebabcase
  - core#string#snakecase
  - core#string#capitalize
  - core#string#chunkstring
  - core#string#substring
position: 80
signatures:
  - "func Words(str string) []string"
---

Splits a string into a slice of its words, separating letters and digits and removing non-alphanumeric separators.

```go
lo.Words("helloWorld")
// []string{"hello", "world"}
```


