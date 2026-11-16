---
name: Capitalize
slug: capitalize
sourceRef: string.go#L424
category: core
subCategory: string
playUrl: https://go.dev/play/p/uLTZZQXqnsa
variantHelpers:
  - core#string#capitalize
  - core#string#capitalizewithlanguage
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
  - "func CapitalizeWithLanguage(str string, tag language.Tag) string"
---

Converts the first character to uppercase and the remaining to lowercase.

```go
lo.Capitalize("heLLO")
// "Hello"
```

`CapitalizeWithLanguage` uses a locale-aware title caser. This matters for languages such as Turkish, where the uppercase of `i` is `İ` (dotted I), not `I`.

```go
lo.CapitalizeWithLanguage("istanbul", language.Turkish)
// "İstanbul"

lo.CapitalizeWithLanguage("istanbul", language.English)
// "Istanbul"
```


