---
name: CamelCase
slug: camelcase
sourceRef: string.go#L310
category: core
subCategory: string
playUrl: https://go.dev/play/p/4JNDzaMwXkm
variantHelpers:
  - core#string#camelcase
  - core#string#camelcasewithlanguage
similarHelpers:
  - core#string#pascalcase
  - core#string#snakecase
  - core#string#kebabcase
  - core#string#capitalize
  - core#string#words
position: 50
signatures:
  - "func CamelCase(str string) string"
  - "func CamelCaseWithLanguage(str string, tag language.Tag) string"
---

Converts a string to camelCase.

```go
lo.CamelCase("hello_world")
// "helloWorld"
```

`CamelCaseWithLanguage` uses locale-aware casing. The first word is fully lowercased, subsequent words are title-cased, both according to the given language tag.

```go
lo.CamelCaseWithLanguage("ISTANBUL_CITY", language.Turkish)
// "ıstanbulCıty"  (capital I → ı, title I → İ in Turkish)
```


