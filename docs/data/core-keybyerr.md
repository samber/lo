---
name: KeyByErr
slug: keybyerr
sourceRef: slice.go#L576
category: core
subCategory: slice
signatures:
  - "func KeyByErr[K comparable, V any](collection []V, iteratee func(item V) (K, error)) (map[K]V, error)"
variantHelpers:
  - core#slice#keybyerr
similarHelpers:
  - core#slice#keyby
  - core#slice#groupby
  - core#slice#groupbyerr
  - core#slice#partitionby
  - core#map#associate
  - core#slice#keyify
position: 231
---

Transforms a slice to a map using a pivot callback to compute keys. Stops iteration immediately when an error is encountered.

```go
// Error case - stops on first error
result, err := lo.KeyByErr([]string{"a", "aa", "aaa", ""}, func(str string) (int, error) {
    if str == "" {
        return 0, fmt.Errorf("empty string not allowed")
    }
    return len(str), nil
})
// map[int]string(nil), error("empty string not allowed")
```

```go
// Success case
result, err := lo.KeyByErr([]string{"a", "aa", "aaa"}, func(str string) (int, error) {
    if str == "" {
        return 0, fmt.Errorf("empty string not allowed")
    }
    return len(str), nil
})
// map[int]string{1: "a", 2: "aa", 3: "aaa"}, nil
```

