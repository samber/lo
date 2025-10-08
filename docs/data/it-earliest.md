---
name: Earliest
slug: earliest
sourceRef: it/find.go#L278
category: it
subCategory: find
signatures:
  - "func Earliest(times iter.Seq[time.Time]) time.Time"
playUrl: "https://go.dev/play/p/7EyYRV1-zd"
variantHelpers:
  - it#find#earliest
similarHelpers:
  - core#slice#earliest
position: 500
---

Searches for the earliest (minimum) time.Time in a collection.

Returns zero value when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
import "time"

// Find the earliest time from a collection
times := it.Slice([]time.Time{
    time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC),
    time.Date(2023, 3, 20, 14, 30, 0, 0, time.UTC),
    time.Date(2023, 8, 1, 9, 15, 0, 0, time.UTC),
})
earliest := it.Earliest(times)
// earliest: 2023-03-20 14:30:00 +0000 UTC

// With empty collection
empty := it.Slice([]time.Time{})
earliest := it.Earliest(empty)
// earliest: 0001-01-01 00:00:00 +0000 UTC (zero value)

// Find earliest from parsed times
times := it.Slice([]time.Time{
    time.Parse(time.RFC3339, "2023-01-01T12:00:00Z"),
    time.Parse(time.RFC3339, "2023-01-01T10:00:00Z"),
    time.Parse(time.RFC3339, "2023-01-01T14:00:00Z"),
})
earliest := it.Earliest(times)
// earliest: 2023-01-01 10:00:00 +0000 UTC
```