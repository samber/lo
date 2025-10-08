---
name: Latest
slug: latest
sourceRef: it/find.go#L346
category: it
subCategory: find
signatures:
  - "func Latest(times iter.Seq[time.Time]) time.Time"
playUrl: "https://go.dev/play/p/3KeXB7-zj"
variantHelpers:
  - it#find#latest
similarHelpers:
  - core#slice#latest
position: 520
---

Searches for the latest (maximum) time.Time in a collection.

Returns zero value when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
import "time"

// Find the latest time from a collection
times := it.Slice([]time.Time{
    time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC),
    time.Date(2023, 3, 20, 14, 30, 0, 0, time.UTC),
    time.Date(2023, 8, 1, 9, 15, 0, 0, time.UTC),
})
latest := it.Latest(times)
// latest: 2023-08-01 09:15:00 +0000 UTC

// With empty collection
empty := it.Slice([]time.Time{})
latest := it.Latest(empty)
// latest: 0001-01-01 00:00:00 +0000 UTC (zero value)

// Find latest from parsed times
times := it.Slice([]time.Time{
    time.Parse(time.RFC3339, "2023-01-01T12:00:00Z"),
    time.Parse(time.RFC3339, "2023-01-01T10:00:00Z"),
    time.Parse(time.RFC3339, "2023-01-01T14:00:00Z"),
})
latest := it.Latest(times)
// latest: 2023-01-01 14:00:00 +0000 UTC

// Find latest log entry timestamp
logs := it.Slice([]time.Time{
    time.Now().Add(-2 * time.Hour),
    time.Now().Add(-1 * time.Hour),
    time.Now(),
})
latest := it.Latest(logs)
// latest: current time
```