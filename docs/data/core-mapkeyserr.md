---
name: MapKeysErr
slug: mapkeyserr
sourceRef: map.go#L293
category: core
subCategory: map
signatures:
  - "func MapKeysErr[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) (R, error)) (map[R]V, error)"
variantHelpers:
  - core#map#mapkeyserr
similarHelpers:
  - core#map#mapkeys
  - core#map#mapvalueserr
position: 185
---

Transforms map keys using a predicate while keeping values. Returns an error if the iteratee function fails, stopping iteration immediately.

```go
in := map[int]int{1: 1, 2: 2, 3: 3}
out, err := lo.MapKeysErr(in, func(v int, _ int) (string, error) {
    if v == 2 {
        return "", fmt.Errorf("even number not allowed")
    }
    return strconv.Itoa(v), nil
})
// map[string]int(nil), error("even number not allowed")
```

```go
in := map[int]int{1: 1, 2: 2, 3: 3}
out, err := lo.MapKeysErr(in, func(v int, _ int) (string, error) {
    return strconv.Itoa(v), nil
})
// map[string]int{"1":1, "2":2, "3":3}, nil
```
