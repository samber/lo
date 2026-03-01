---
name: FilterKeysErr
slug: filterkeyserr
sourceRef: map.go#L498
category: core
subCategory: map
signatures:
  - "func FilterKeysErr[K comparable, V any](in map[K]V, predicate func(key K, value V) (bool, error)) ([]K, error)"
playUrl:
variantHelpers:
  - core#map#filterkeyserr
similarHelpers:
  - core#map#filterkeys
  - core#map#filtervalueserr
  - core#slice#filter
position: 235
---

Transforms a map into a slice of keys based on a predicate that can return an error. It is a mix of Filter() and Keys() with error handling. If the predicate returns true, the key is included. If the predicate returns an error, iteration stops immediately and returns the error.

```go
kv := map[int]string{1:"foo", 2:"bar", 3:"baz"}
result, err := lo.FilterKeysErr(kv, func(k int, v string) (bool, error) {
    if k == 3 {
        return false, errors.New("key 3 not allowed")
    }
    return v == "foo", nil
})
// []int(nil), error("key 3 not allowed")
```

```go
kv := map[int]string{1:"foo", 2:"bar", 3:"baz"}
result, err := lo.FilterKeysErr(kv, func(k int, v string) (bool, error) {
    return v == "bar", nil
})
// []int{2}, nil
```
