---
name: GroupByMap
slug: groupbymap
sourceRef: slice.go#L194
category: core
subCategory: slice
playUrl: https://go.dev/play/p/iMeruQ3_W80
variantHelpers:
  - core#slice#groupbymap
similarHelpers:
  - core#slice#groupby
  - core#slice#partitionby
  - core#slice#keyby
  - core#map#associate
  - parallel#slice#groupby
position: 130
signatures:
  - "func GroupByMap[T any, K comparable, V any](collection []T, iteratee func(item T) (K, V)) map[K][]V"
---

Groups items by a key computed from each element and maps each element to a value.

```go
groups := lo.GroupByMap(
    []int{0, 1, 2, 3, 4, 5},
    func(i int) (int, int) {
        return i % 3, i * 2
    },
)
// map[int][]int{0:{0,6}, 1:{2,8}, 2:{4,10}}
```


