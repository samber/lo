---
name: LatestBy
slug: latestby
sourceRef: find.go#L530
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#latestby
similarHelpers:
  - core#find#latest
  - core#find#earliestby
  - core#find#earliest
  - core#find#maxby
  - core#find#minby
  - core#find#maxindexby
  - core#find#minindexby
  - core#find#findby
  - core#find#findkeyby
  - core#find#findduplicatesby
  - core#find#finduniquesby
position: 250
signatures:
  - "func LatestBy[T any](collection []T, iteratee func(item T) time.Time) T"
---

Searches a collection for the element with the maximum time extracted by the predicate. Returns zero value when the collection is empty.

```go
type Event struct{ At time.Time }
events := []Event{{At: time.Now()}, {At: time.Now().Add(2 * time.Hour)}}
last := lo.LatestBy(events, func(e Event) time.Time {
    return e.At
})
```


