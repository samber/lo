---
name: FindErr
slug: finderr
sourceRef: find.go#L83
category: core
subCategory: find
signatures:
  - "func FindErr[T any](collection []T, predicate func(item T) (bool, error)) (T, error)"
playUrl:
variantHelpers: []
similarHelpers:
  - core#find#find
  - core#find#findorelse
  - core#find#findkey
  - core#find#findindexof
  - core#slice#filter
position: 41
---

Searches for an element in a slice based on a predicate that can return an error. Returns the element and nil error if found. Returns zero value and nil error if not found. If the predicate returns an error, iteration stops immediately and returns zero value and the error.

```go
result, err := lo.FindErr([]string{"a", "b", "c", "d"}, func(i string) (bool, error) {
    return i == "b", nil
})
// "b", nil

result, err = lo.FindErr([]string{"foobar"}, func(i string) (bool, error) {
    return i == "b", nil
})
// "", nil

result, err = lo.FindErr([]string{"a", "b", "c"}, func(i string) (bool, error) {
    if i == "b" {
        return false, fmt.Errorf("b is not allowed")
    }
    return i == "b", nil
})
// "", error("b is not allowed")
```
