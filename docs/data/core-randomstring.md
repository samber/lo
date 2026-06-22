---
name: RandomString
slug: randomstring
sourceRef: string.go#L35
category: core
subCategory: string
playUrl: https://go.dev/play/p/rRseOQVVum4
variantHelpers:
  - core#string#randomstring
similarHelpers:
  - core#string#substring
  - core#string#chunkstring
  - core#string#words
  - core#string#capitalize
  - core#string#camelcase
  - core#string#pascalcase
  - core#string#kebabcase
  - core#string#snakecase
position: 0
signatures:
  - "func RandomString(size int, charset []rune) string"
---

Returns a random string of the specified length from the given charset.

```go
str := lo.RandomString(5, lo.LettersCharset)
// e.g., "eIGbt"
```


