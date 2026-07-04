---
name: SnakeCase
slug: snakecase
sourceRef: string.go#L376
category: core
subCategory: string
playUrl: https://go.dev/play/p/ziB0V89IeVH
variantHelpers:
  - core#string#snakecase
  - core#string#snakecasewithlanguage
similarHelpers:
  - core#string#pascalcase
  - core#string#camelcase
  - core#string#kebabcase
  - core#string#capitalize
  - core#string#words
position: 70
signatures:
  - "func SnakeCase(str string) string"
  - "func SnakeCaseWithLanguage(str string, tag language.Tag) string"
---

Converts a string to snake_case.

```go
lo.SnakeCase("HelloWorld")
// "hello_world"
```

`SnakeCaseWithLanguage` uses a locale-aware lowercase caser instead of `strings.ToLower`. This matters for languages such as Turkish, where capital `I` lowercases to `ı` (dotless i, U+0131), not `i`.

```go
lo.SnakeCaseWithLanguage("ISTANBUL_CITY", language.Turkish)
// "ıstanbul_cıty"

lo.SnakeCaseWithLanguage("ISTANBUL_CITY", language.English)
// "istanbul_city"
```


