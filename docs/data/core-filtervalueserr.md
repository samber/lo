---
name: FilterValuesErr
slug: filtervalueserr
sourceRef: map.go#L519
category: core
subCategory: map
signatures:
  - "func FilterValuesErr[K comparable, V any](in map[K]V, predicate func(key K, value V) (bool, error)) ([]V, error)"
playUrl: https://go.dev/play/p/hKvHlqLzbdE
variantHelpers:
  - core#map#filtervalueserr
similarHelpers:
  - core#map#filtervalues
  - core#map#filterkeyserr
  - core#slice#filter
position: 245
---

Transforms a map into a slice of values based on a predicate that can return an error. It is a mix of Filter() and Values() with error handling. If the predicate returns true, the value is included. If the predicate returns an error, iteration stops immediately and returns the error.

```go
kv := map[int]string{1:"foo", 2:"bar", 3:"baz"}
result, err := lo.FilterValuesErr(kv, func(k int, v string) (bool, error) {
    if k == 3 {
        return false, errors.New("key 3 not allowed")
    }
    return v == "foo", nil
})
// []string(nil), error("key 3 not allowed")
```

```go
kv := map[int]string{1:"foo", 2:"bar", 3:"baz"}
result, err := lo.FilterValuesErr(kv, func(k int, v string) (bool, error) {
    return v == "bar", nil
})
// []string{"bar"}, nil
```
