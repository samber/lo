---
name: NewTransaction
slug: newtransaction
sourceRef: retry.go#L253
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/Qxrd7MGQGh1
variantHelpers:
  - core#concurrency#newtransaction
similarHelpers:
  - core#concurrency#synchronize
  - core#concurrency#asyncx
  - core#concurrency#waitfor
position: 30
signatures:
  - "func NewTransaction[T any]() *Transaction[T]"
---

Creates a new Saga transaction that chains steps with rollback functions.

Use Then to add steps with exec and rollback functions. Call Process with an initial state to execute the pipeline; if a step returns an error, previously executed steps are rolled back in reverse order using their rollback functions.

```go
type Acc struct{ Sum int }

tx := lo.NewTransaction[Acc]().
    Then(
        func(a Acc) (Acc, error) {
            a.Sum += 10
            return a, nil
        },
        func(a Acc) Acc {
            a.Sum -= 10
            return a
        },
    ).
    Then(
        func(a Acc) (Acc, error) {
            a.Sum *= 3
            return a, nil
        },
        func(a Acc) Acc {
            a.Sum /= 3
            return a
        },
    )

res, err := tx.Process(Acc{Sum: 1})
// res.Sum == 33, err == nil
```
