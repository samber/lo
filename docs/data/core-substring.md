---
name: Substring
slug: substring
sourceRef: string.go#L105
category: core
subCategory: string
playUrl: https://go.dev/play/p/TQlxQi82Lu1
variantHelpers:
  - core#string#substring
similarHelpers:
  - core#string#chunkstring
  - core#string#runelength
  - core#string#ellipsis
  - core#string#words
  - core#slice#replace
  - core#slice#replaceall
position: 10
signatures:
  - "func Substring[T ~string](str T, offset int, length uint) T"
---

Returns a substring starting at the given offset with the specified length. Supports negative offsets; out-of-bounds are clamped.

```go
// Basic usage
result := lo.Substring("hello", 2, 3)
// result: "llo"

// Negative offset - counts from end
result = lo.Substring("hello", -4, 3)
// result: "ell"

// Length longer than string - clamped to available characters
result = lo.Substring("hello", 1, 10)
// result: "ello" (only 4 characters available from position 1)

// Zero length - returns empty string
result = lo.Substring("hello", 1, 0)
// result: ""

// Offset beyond string length - returns empty string
result = lo.Substring("hello", 10, 3)
// result: ""

// With Unicode strings (byte-based)
result = lo.Substring("héllo", 1, 3)
// result: "él" (note: works with bytes, not runes)

// Negative offset with negative values clamped
result = lo.Substring("hello", -10, 3)
// result: "hel" (offset clamped to 0)
```


