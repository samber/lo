---
name: MapValuesErr
slug: mapvalueserr
sourceRef: map.go#L322
category: core
subCategory: map
signatures:
  - "func MapValuesErr[K comparable, V any, R any](in map[K]V, iteratee func(value V, key K) (R, error)) (map[K]R, error)"
variantHelpers:
  - core#map#mapvalueserr
similarHelpers:
  - core#map#mapvalues
  - core#map#mapkeyserr
  - core#map#mapentrieserr
position: 195
---

Transforms map values using a predicate while keeping keys. Returns an error if the iteratee function fails, stopping iteration immediately.

```go
in := map[int]int64{1: 1, 2: 2, 3: 3}
out, err := lo.MapValuesErr(in, func(v int64, _ int) (string, error) {
    if v == 2 {
        return "", fmt.Errorf("even number not allowed")
    }
    return strconv.FormatInt(v, 10), nil
})
// map[int]string(nil), error("even number not allowed")
```

```go
in := map[int]int64{1: 1, 2: 2, 3: 3}
out, err := lo.MapValuesErr(in, func(v int64, _ int) (string, error) {
    return strconv.FormatInt(v, 10), nil
})
// map[int]string{1:"1", 2:"2", 3:"3"}, nil
```
