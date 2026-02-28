---
name: OmitByErr
slug: omitbyerr
sourceRef: map.go#L171
category: core
subCategory: map
signatures:
  - "func OmitByErr[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) (bool, error)) (Map, error)"
variantHelpers:
  - core#map#omitbyerr
similarHelpers:
  - core#map#omitby
  - core#map#pickbyerr
  - core#map#pickby
position: 95
---

Returns a map of the same type excluding entries that match the predicate. Returns an error if the predicate function fails, stopping iteration immediately.

```go
m, err := lo.OmitByErr(
    map[string]int{"foo": 1, "bar": 2, "baz": 3},
    func(key string, value int) (bool, error) {
        if key == "bar" {
            return false, fmt.Errorf("bar not allowed")
        }
        return value%2 == 1, nil
    },
)
// map[string]int(nil), error("bar not allowed")
```

```go
m, err := lo.OmitByErr(
    map[string]int{"foo": 1, "bar": 2, "baz": 3},
    func(key string, value int) (bool, error) {
        return value%2 == 1, nil
    },
)
// map[string]int{"bar": 2}, nil
```
