---
name: AttemptWhile
slug: attemptwhile
sourceRef: retry.go#L198
category: core
subCategory: retry
playUrl: https://go.dev/play/p/1VS7HxlYMOG
variantHelpers:
  - core#retry#attemptwhile
similarHelpers:
  - core#retry#attempt
  - core#retry#attemptwithdelay
  - core#retry#attemptwhilewithdelay
position: 20
signatures:
  - "func AttemptWhile(maxIteration int, f func(int) (error, bool)) (int, error)"
---

Invokes a function up to N times until it returns nil. The second return value controls whether to continue attempting.

```go
count, err := lo.AttemptWhile(5, func(i int) (error, bool) {
    if i == 2 {
        return nil, false
    }
    return errors.New("fail"), true
})
// count == 3, err == nil
```


