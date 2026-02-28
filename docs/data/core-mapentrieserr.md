---
name: MapEntriesErr
slug: mapentrieserr
sourceRef: map.go#L350
category: core
subCategory: map
signatures:
  - "func MapEntriesErr[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2, error)) (map[K2]V2, error)"
variantHelpers:
  - core#map#mapentrieserr
similarHelpers:
  - core#map#mapentries
  - core#map#mapkeyserr
  - core#map#mapvalueserr
position: 205
---

Transforms both keys and values using an predicate function. Returns an error if the iteratee function fails, stopping iteration immediately.

```go
in := map[string]int{"foo": 1, "bar": 2, "baz": 3}
out, err := lo.MapEntriesErr(in, func(k string, v int) (int, string, error) {
    if k == "bar" {
        return 0, "", fmt.Errorf("bar not allowed")
    }
    return v, k, nil
})
// map[int]string(nil), error("bar not allowed")
```

```go
in := map[string]int{"foo": 1, "bar": 2}
out, err := lo.MapEntriesErr(in, func(k string, v int) (int, string, error) {
    return v, k, nil
})
// map[int]string{1:"foo", 2:"bar"}, nil
```
