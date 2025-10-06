---
name: EarliestBy
slug: earliestby
sourceRef: find.go#L385
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#earliestby
similarHelpers:
  - core#find#latestby
  - core#find#earliest
  - core#find#latest
  - core#find#minby
  - core#find#maxby
  - core#find#minindexby
  - core#find#maxindexby
  - core#find#findby
  - core#find#findkeyby
  - core#find#findduplicatesby
  - core#find#finduniquesby
position: 190
signatures:
  - "func EarliestBy[T any](collection []T, iteratee func(item T) time.Time) T"
---

Searches a collection for the element with the minimum time extracted by the predicate. Returns zero value when the collection is empty.

```go
type Event struct{ At time.Time }
events := []Event{{At: time.Now().Add(2 * time.Hour)}, {At: time.Now()}}
first := lo.EarliestBy(events, func(e Event) time.Time {
    return e.At
})
```


