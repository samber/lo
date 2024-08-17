
## Describe your changes

...

## Checklist before requesting a review

- [ ] 👓 I have performed a self-review of my code
- [ ] 👶 This helper does not already exist
- [ ] 🧪 This helper is tested
- [ ] 🏎️ My code limits memory allocation and is fast
- [ ] 🧞‍♂️ This helper is immutable and my tests prove it
- [ ] ✍️ I implemented the parallel and mutable variants
- [ ] 📖 My helper has been added to README
- [ ] 🔬 An example has been added to xxxxx_example_test.go
- [ ] ⛹️ An example has been created on https://go.dev/play

## Conventions

- Returning `(ok bool)` is often better than `(err error)`
- `panic(...)` must be limited
- Helpers should allow batching (eg: receive variadic arguments)
- Use an index at the end of the helper name to declare variants (eg: `lo.Must0`, `lo.Must1`, `lo.Must2`...)
