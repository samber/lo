---
name: LatestBy
slug: latestby
sourceRef: it/find.go#L353
category: it
subCategory: find
signatures:
  - "func LatestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T"
playUrl: "https://go.dev/play/p/4LfYC8-zk"
variantHelpers:
  - it#find#latestby
similarHelpers:
  - core#slice#latestby
position: 530
---

Searches for the element with the latest time using a transform function.

Returns zero value when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
import "time"

type Event struct {
    Name string
    Time time.Time
}

// Find the latest event by time
events := it.Slice([]Event{
    {"Meeting", time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC)},
    {"Lunch", time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)},
    {"Breakfast", time.Date(2023, 5, 15, 8, 0, 0, 0, time.UTC)},
})
latest := it.LatestBy(events, func(e Event) time.Time {
    return e.Time
})
// latest: {Name: "Lunch", Time: 2023-05-15 12:00:00 +0000 UTC}

// Find the latest task by deadline
type Task struct {
    ID       int
    Deadline time.Time
}
tasks := it.Slice([]Task{
    {1, time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)},
    {2, time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC)},
    {3, time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)},
})
latest := it.LatestBy(tasks, func(t Task) time.Time {
    return t.Deadline
})
// latest: {ID: 3, Deadline: 2023-07-01 00:00:00 +0000 UTC}

// Find the most recent activity
type Activity struct {
    User    string
    Action  string
    Time    time.Time
}
activities := it.Slice([]Activity{
    {"alice", "login", time.Now().Add(-24 * time.Hour)},
    {"bob", "logout", time.Now().Add(-12 * time.Hour)},
    {"alice", "post", time.Now().Add(-1 * time.Hour)},
})
latest := it.LatestBy(activities, func(a Activity) time.Time {
    return a.Time
})
// latest: {User: "alice", Action: "post", Time: 1 hour ago}
```