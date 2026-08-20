# AGENTS.md — samber/lo

Guidance for AI coding agents (Claude Code, Cursor, Copilot, Codex...) using or
extending `samber/lo`, a Lodash-style utility library for Go 1.18+ generics.

## What this module provides

| Package | Import path | What it's for |
|---|---|---|
| Core | `github.com/samber/lo` | 274 immutable, type-safe helpers for slices, maps, strings, channels, math, tuples, retries and error handling. |
| Iterators (Go 1.23+) | `github.com/samber/lo/it` | Lazy, allocation-free helpers over `iter.Seq`/`iter.Seq2` — 157 helpers. Requires Go 1.23+. |
| Mutable | `github.com/samber/lo/mutable` | In-place, zero-allocation variants of core slice helpers, for performance-critical code. |
| Parallel | `github.com/samber/lo/parallel` | Goroutine-per-item variants of core slice helpers, with ordered results. |
| Experimental SIMD | `github.com/samber/lo/exp/simd` | Unstable, high-performance slice ops using AVX/AVX2/AVX512. Requires `go1.26+goexperiment.simd+amd64`. Not covered by SemVer. |

449 helpers are documented in full at https://lo.samber.dev with runnable
Go Playground examples — read a helper's doc page or its GoDoc comment
before assuming a signature; do not guess.

## Which package to reach for

- Transforming a slice/map you already hold in memory, synchronously → `lo`.
- The same, but you want to avoid intermediate allocations, or the input
  is `iter.Seq`/`iter.Seq2` (Go 1.23+ `range-over-func`) → `lo/it`.
- Mutating a slice in place instead of allocating a new one → `lo/mutable`.
- The transform function is slow/IO-bound and items are independent →
  `lo/parallel` (one goroutine per item — do not use this for cheap,
  CPU-bound transforms; goroutine overhead will dominate).

## When NOT to use `lo`

- If the task is covered by the standard library's `slices` or `maps`
  packages (`slices.Contains`, `slices.Sort`, `slices.Index`, `maps.Keys`,
  simple `slices.Chunk`...), prefer the stdlib. `lo` fills gaps the stdlib
  doesn't cover (`GroupBy`, `PartitionBy`, error-aware `...Err` variants,
  parallel execution, etc.) — it is not a blanket replacement for it.
- Don't import `lo` for a single one-off transform that a 3-line `for` loop
  already expresses clearly — the value of `lo` is consistency across many
  call sites, not brevity on one.
- `lo/exp/simd` is explicitly unstable; do not introduce it into
  production code paths without the caller being aware of the build-tag
  and platform constraints.

## Conventions when writing code that uses `lo`

- Function-suffix conventions (see `docs/CLAUDE.md` for the full list):
  `F` = lazy/function-based variant, `I` = index-aware predicate,
  `Err` = predicate/transform can return an error, `WithContext` =
  accepts a `context.Context`, `X` = variable-arity (`Must2`, `Zip3`...).
- Prefer the `Err`-suffixed variant over `Must`/panicking helpers in
  library code; reserve `Must*` for `main()`, `init()`, or tests.
- Zero dependencies outside the Go standard library for the `lo` core
  package (verify in `go.mod` before adding one).

## Documentation contribution conventions

If you are asked to add or update a helper, read `docs/CLAUDE.md` first —
it defines the required frontmatter for `docs/data/*.md`, the
`variantHelpers`/`similarHelpers` cross-reference format, and the
Go Playground example workflow. After editing any `.go` file, run
`cd docs && npm run check-docs` — it validates cross-references, filenames,
duplicate slugs and that every documented helper has a rendering page.
