---
name: UnzipByErrX
slug: unzipbyerrx
sourceRef: tuples.go#L1024
category: core
subCategory: tuple
signatures:
  - "func UnzipByErr2[In any, A any, B any](items []In, predicate func(In) (a A, b B, err error)) ([]A, []B, error)"
  - "func UnzipByErr3[In any, A any, B any, C any](items []In, predicate func(In) (a A, b B, c C, err error)) ([]A, []B, []C, error)"
  - "func UnzipByErr4[In any, A any, B any, C any, D any](items []In, predicate func(In) (a A, b B, c C, d D, err error)) ([]A, []B, []C, []D, error)"
  - "func UnzipByErr5[In any, A any, B any, C any, D any, E any](items []In, predicate func(In) (a A, b B, c C, d D, e E, err error)) ([]A, []B, []C, []D, []E, error)"
  - "func UnzipByErr6[In any, A any, B any, C any, D any, E any, F any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, err error)) ([]A, []B, []C, []D, []E, []F, error)"
  - "func UnzipByErr7[In any, A any, B any, C any, D any, E any, F any, G any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, g G, err error)) ([]A, []B, []C, []D, []E, []F, []G, error)"
  - "func UnzipByErr8[In any, A any, B any, C any, D any, E any, F any, G any, H any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, g G, h H, err error)) ([]A, []B, []C, []D, []E, []F, []G, []H, error)"
  - "func UnzipByErr9[In any, A any, B any, C any, D any, E any, F any, G any, H any, I any](items []In, predicate func(In) (a A, b B, c C, d D, e E, f F, g G, h H, i I, err error)) ([]A, []B, []C, []D, []E, []F, []G, []H, []I, error)"
variantHelpers:
  - core#tuple#unzipbyerrx
similarHelpers:
  - core#tuple#unzipbyx
  - core#tuple#tuplex
  - core#slice#map
position: 41
---

Transforms each input element into a tuple and splits results into parallel slices. The iteratee can return an error to stop iteration immediately. Variants support arities from 2 to 9.

Variants: `UnzipByErr2..UnzipByErr9`

```go
a, b, err := lo.UnzipByErr2([]string{"hello", "error", "world"}, func(str string) (string, int, error) {
    if str == "error" {
        return "", 0, fmt.Errorf("error string not allowed")
    }
    return str, len(str), nil
})
// []string{}
// []int{}
// error string not allowed
```

On error, all result slices are `nil` and iteration stops immediately.
