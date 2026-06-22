---
name: PickByErr
slug: pickbyerr
sourceRef: map.go#L117
category: core
subCategory: map
signatures:
  - "func PickByErr[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) (bool, error)) (Map, error)"
variantHelpers:
  - core#map#pickbyerr
similarHelpers:
  - core#map#pickby
  - core#map#omitby
  - core#map#omitbyerr
position: 65
---

Returns a map of the same type filtered by a key/value predicate. Returns an error if the predicate function fails, stopping iteration immediately.

```go
m, err := lo.PickByErr(
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
m, err := lo.PickByErr(
    map[string]int{"foo": 1, "bar": 2, "baz": 3},
    func(key string, value int) (bool, error) {
        return value%2 == 1, nil
    },
)
// map[string]int{"foo": 1, "baz": 3}, nil
```
