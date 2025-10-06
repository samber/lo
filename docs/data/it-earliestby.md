---
name: EarliestBy
slug: earliestby
sourceRef: it/find.go#L285
category: it
subCategory: find
signatures:
  - "func EarliestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T"
playUrl: ""
variantHelpers:
  - it#find#earliestby
similarHelpers:
  - core#slice#earliestby
position: 510
---

Searches for the element with the earliest time using a transform function.

Returns zero value when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
import "time"

type Event struct {
    Name string
    Time time.Time
}

// Find the earliest event by time
events := it.Slice([]Event{
    {"Meeting", time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC)},
    {"Lunch", time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)},
    {"Breakfast", time.Date(2023, 5, 15, 8, 0, 0, 0, time.UTC)},
})
earliest := it.EarliestBy(events, func(e Event) time.Time {
    return e.Time
})
// earliest: {Name: "Breakfast", Time: 2023-05-15 08:00:00 +0000 UTC}

// Find the earliest task by deadline
type Task struct {
    ID       int
    Deadline time.Time
}
tasks := it.Slice([]Task{
    {1, time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)},
    {2, time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)},
    {3, time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)},
})
earliest := it.EarliestBy(tasks, func(t Task) time.Time {
    return t.Deadline
})
// earliest: {ID: 2, Deadline: 2023-05-15 00:00:00 +0000 UTC}
```