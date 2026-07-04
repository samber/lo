---
name: PascalCase
slug: pascalcase
sourceRef: string.go#L280
category: core
subCategory: string
playUrl: https://go.dev/play/p/uxER7XpRHLB
variantHelpers:
  - core#string#pascalcase
  - core#string#pascalcasewithlanguage
similarHelpers:
  - core#string#camelcase
  - core#string#snakecase
  - core#string#kebabcase
  - core#string#capitalize
  - core#string#words
position: 40
signatures:
  - "func PascalCase(str string) string"
  - "func PascalCaseWithLanguage(str string, tag language.Tag) string"
---

Converts a string to PascalCase.

```go
lo.PascalCase("hello_world")
// "HelloWorld"
```

`PascalCaseWithLanguage` uses a locale-aware title caser. This matters for languages such as Turkish, where the uppercase of `i` is `İ` (dotted I), not `I`.

```go
lo.PascalCaseWithLanguage("istanbul_city", language.Turkish)
// "İstanbulCity"
```


