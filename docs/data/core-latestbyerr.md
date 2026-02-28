---
name: LatestByErr
slug: latestbyerr
sourceRef: find.go#L737
category: core
subCategory: find
signatures:
  - "func LatestByErr[T any](collection []T, iteratee func(item T) (time.Time, error)) (T, error)"
variantHelpers:
  - core#find#latestbyerr
similarHelpers:
  - core#find#latestby
  - core#find#latest
  - core#find#earliestby
  - core#find#earliestbyerr
  - core#find#earliest
  - core#find#maxby
  - core#find#minby
  - core#find#maxindexby
  - core#find#minindexby
  - core#find#findby
  - core#find#findkeyby
  - core#find#findduplicatesby
  - core#find#finduniquesby
position: 251
---

Searches a collection for the element with the maximum time extracted by the predicate. Returns zero value when the collection is empty. Stops iteration immediately when an error is encountered.

```go
type Event struct{ At time.Time }
events := []Event{{At: time.Now()}, {At: time.Now().Add(2 * time.Hour)}}
last, err := lo.LatestByErr(events, func(e Event) (time.Time, error) {
    return e.At, nil
})
// Event{At: ...}, nil
```

```go
// Error case - stops on first error
type Event struct{ At time.Time }
events := []Event{{At: time.Now()}, {At: time.Time{}}, {At: time.Now().Add(2 * time.Hour)}}
_, err := lo.LatestByErr(events, func(e Event) (time.Time, error) {
    if e.At.IsZero() {
        return time.Time{}, fmt.Errorf("zero time not allowed")
    }
    return e.At, nil
})
// error("zero time not allowed")
```

