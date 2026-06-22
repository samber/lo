---
name: ZipByErrX
slug: zipbyerrx
sourceRef: tuples.go#L444
category: core
subCategory: tuple
signatures:
  - "func ZipByErr2[A any, B any, Out any](a []A, b []B, iteratee func(a A, b B) (Out, error)) ([]Out, error)"
  - "func ZipByErr3[A any, B any, C any, Out any](a []A, b []B, c []C, iteratee func(a A, b B, c C) (Out, error)) ([]Out, error)"
  - "func ZipByErr4[A any, B any, C any, D any, Out any](a []A, b []B, c []C, d []D, iteratee func(a A, b B, c C, d D) (Out, error)) ([]Out, error)"
  - "func ZipByErr5[A any, B any, C any, D any, E any, Out any](a []A, b []B, c []C, d []D, e []E, iteratee func(a A, b B, c C, d D, e E) (Out, error)) ([]Out, error)"
  - "func ZipByErr6[A any, B any, C any, D any, E any, F any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, iteratee func(a A, b B, c C, d D, e E, f F) (Out, error)) ([]Out, error)"
  - "func ZipByErr7[A any, B any, C any, D any, E any, F any, G any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, iteratee func(a A, b B, c C, d D, e E, f F, g G) (Out, error)) ([]Out, error)"
  - "func ZipByErr8[A any, B any, C any, D any, E any, F any, G any, H any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, iteratee func(a A, b B, c C, d D, e E, f F, g G, h H) (Out, error)) ([]Out, error)"
  - "func ZipByErr9[A any, B any, C any, D any, E any, F any, G any, H any, I any, Out any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, iteratee func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (Out, error)) ([]Out, error)"
variantHelpers:
  - core#tuple#zipbyerrx
similarHelpers:
  - core#tuple#zipbyx
position: 30
---

Zips multiple slices and projects each grouped set through a function that can return an error. Stops iteration immediately when an error is encountered and returns the zero value (nil for slices).

Variants: `ZipByErr2..ZipByErr9`

```go
result, err := lo.ZipByErr2([]string{"a", "b"}, []int{1, 2}, func(a string, b int) (string, error) {
    if b == 2 {
        return "", fmt.Errorf("number 2 is not allowed")
    }
    return fmt.Sprintf("%s-%d", a, b), nil
})
// []string(nil), error("number 2 is not allowed")
```

```go
result, err := lo.ZipByErr2([]string{"a", "b"}, []int{1, 2}, func(a string, b int) (string, error) {
    return fmt.Sprintf("%s-%d", a, b), nil
})
// []string{"a-1", "b-2"}, nil
```

When collections are different sizes, the missing attributes are filled with zero value before calling the iteratee.
