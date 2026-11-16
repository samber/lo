---
name: KebabCase
slug: kebabcase
sourceRef: string.go#L346
category: core
subCategory: string
playUrl: https://go.dev/play/p/ZBeMB4-pq45
variantHelpers:
  - core#string#kebabcase
  - core#string#kebabcasewithlanguage
similarHelpers:
  - core#string#pascalcase
  - core#string#camelcase
  - core#string#snakecase
  - core#string#capitalize
  - core#string#words
position: 60
signatures:
  - "func KebabCase(str string) string"
  - "func KebabCaseWithLanguage(str string, tag language.Tag) string"
---

Converts a string to kebab-case.

```go
lo.KebabCase("helloWorld")
// "hello-world"
```

`KebabCaseWithLanguage` uses a locale-aware lowercase caser instead of `strings.ToLower`. This matters for languages such as Turkish, where capital `I` lowercases to `ı` (dotless i, U+0131), not `i`.

```go
lo.KebabCaseWithLanguage("ISTANBUL_CITY", language.Turkish)
// "ıstanbul-cıty"

lo.KebabCaseWithLanguage("ISTANBUL_CITY", language.English)
// "istanbul-city"
```


