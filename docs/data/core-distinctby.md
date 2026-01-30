---
name: DistinctBy
slug: distinctby
sourceRef: channel.go#L350
category: core
subCategory: channel
signatures:
  - "func DistinctBy[T any, K comparable](upstream <-chan T, key func(item T) K) <-chan T"
similarHelpers:
  - core#slice#uniq
  - core#slice#uniqby
position: 259
---

Removes duplicate values from a channel based on a key function while preserving order.

```go
type user struct {
    ID   int
    Name string
}

ch := lo.SliceToChannel(10, []user{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 1, Name: "Alicia"},
})

distinct := lo.DistinctBy(ch, func(u user) int { return u.ID })
var result []user
for v := range distinct {
    result = append(result, v)
}
// result contains [{1 Alice} {2 Bob}]
```
