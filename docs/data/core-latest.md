---
name: Latest
slug: latest
sourceRef: find.go#L508
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#latest
similarHelpers:
  - core#find#earliest
  - core#find#latestby
  - core#find#earliestby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
  - core#find#minindex
  - core#find#maxindex
  - core#find#minindexby
  - core#find#maxindexby
position: 240
signatures:
  - "func Latest(times ...time.Time) time.Time"
---

Searches the maximum time.Time in the provided arguments. Returns zero value when the input is empty.

```go
t1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
max := lo.Latest(t1, t2)
// 2024-01-01 00:00:00 +0000 UTC
```


