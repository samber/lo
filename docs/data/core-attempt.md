---
name: Attempt
slug: attempt
sourceRef: retry.go#L153
category: core
subCategory: retry
playUrl: https://go.dev/play/p/3ggJZ2ZKcMj
variantHelpers:
  - core#retry#attempt
similarHelpers:
  - core#retry#attemptwithdelay
  - core#retry#attemptwhile
  - core#retry#attemptwhilewithdelay
position: 0
signatures:
  - "func Attempt(maxIteration int, f func(index int) error) (int, error)"
---

Invokes a function up to N times until it returns nil. Returns the number of iterations and the last error.

```go
iter, err := lo.Attempt(5, func(i int) error {
    if i == 2 {
        return nil
    }
    return errors.New("failed")
})
// 3, nil
```


