---
name: Earliest
slug: earliest
sourceRef: find.go#L363
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#earliest
similarHelpers:
  - core#find#latest
  - core#find#earliestby
  - core#find#latestby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
  - core#find#minindex
  - core#find#maxindex
  - core#find#minindexby
  - core#find#maxindexby
position: 180
signatures:
  - "func Earliest(times ...time.Time) time.Time"
---

Searches the minimum time.Time in the provided arguments. Returns zero value when the input is empty.

```go
t1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
min := lo.Earliest(t2, t1)
// 2023-01-01 00:00:00 +0000 UTC
```


