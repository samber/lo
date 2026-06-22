---
name: EarliestByErr
slug: earliestbyerr
sourceRef: find.go#L484
category: core
subCategory: find
playUrl: https://go.dev/play/p/zJUBUj7ANvq
variantHelpers:
  - core#find#earliestbyerr
similarHelpers:
  - core#find#earliestby
  - core#find#latestby
  - core#find#earliest
  - core#find#latest
  - core#find#minby
  - core#find#maxby
position: 191
signatures:
  - "func EarliestByErr[T any](collection []T, iteratee func(item T) (time.Time, error)) (T, error)"
---

Searches a collection for the element with the minimum time extracted by the predicate. Returns zero value and nil error when the collection is empty.

If the iteratee returns an error, iteration stops and the error is returned.

```go
type Event struct{ At time.Time }
events := []Event{{At: time.Now().Add(2 * time.Hour)}, {At: time.Now()}}
first, err := lo.EarliestByErr(events, func(e Event) (time.Time, error) {
    return e.At, nil
})
// Event, <nil>
```

Example with error:

```go
type Event struct{ At time.Time }
events := []Event{{At: time.Now()}, {At: time.Now().Add(time.Hour)}}
first, err := lo.EarliestByErr(events, func(e Event) (time.Time, error) {
    if e.At.After(time.Now().Add(30 * time.Minute)) {
        return time.Time{}, fmt.Errorf("event too far in the future")
    }
    return e.At, nil
})
// Event, error("event too far in the future")
```
