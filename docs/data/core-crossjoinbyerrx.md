---
name: CrossJoinByErrX
slug: crossjoinbyerrx
sourceRef: tuples.go#L1320
category: core
subCategory: tuple
signatures:
  - "func CrossJoinByErr2[A any, B any, Out any](listA []A, listB []B, transform func(a A, b B) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr3[A any, B any, C any, Out any](listA []A, listB []B, listC []C, transform func(a A, b B, c C) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr4[A any, B any, C any, D any, Out any](listA []A, listB []B, listC []C, listD []D, transform func(a A, b B, c C, d D) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr5[A any, B any, C any, D any, E any, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, transform func(a A, b B, c C, d D, e E) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr6[A any, B any, C any, D any, E any, F any, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, transform func(a A, b B, c C, d D, e E, f F) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr7[A any, B any, C any, D any, E any, F any, G any, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, transform func(a A, b B, c C, d D, e E, f F, g G) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr8[A any, B any, C any, D any, E any, F any, G any, H any, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, transform func(a A, b B, c C, d D, e E, f F, g G, h H) (Out, error)) ([]Out, error)"
  - "func CrossJoinByErr9[A any, B any, C any, D any, E any, F any, G any, H any, I any, Out any](listA []A, listB []B, listC []C, listD []D, listE []E, listF []F, listG []G, listH []H, listI []I, transform func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (Out, error)) ([]Out, error)"
variantHelpers:
  - core#tuple#crossjoinbyerrx
similarHelpers:
  - core#tuple#crossjoinbyx
position: 61
---

Computes a cartesian product and projects each combination through a function that can return an error. Stops iteration immediately when an error is encountered and returns the zero value (nil for slices).

Variants: `CrossJoinByErr2..CrossJoinByErr9`

```go
result, err := lo.CrossJoinByErr2([]string{"a", "b"}, []int{1, 2}, func(a string, b int) (string, error) {
    if a == "b" {
        return "", fmt.Errorf("b not allowed")
    }
    return fmt.Sprintf("%s-%d", a, b), nil
})
// []string(nil), error("b not allowed")
```

```go
result, err := lo.CrossJoinByErr2([]string{"a", "b"}, []int{1, 2}, func(a string, b int) (string, error) {
    return fmt.Sprintf("%s-%d", a, b), nil
})
// []string{"a-1", "a-2", "b-1", "b-2"}, nil
```

Returns an empty list if any input list is empty.
