# Adding New Helper Documentation

When adding a new helper to the `lo` library, follow these guidelines to create or update documentation files in the `docs/data/` directory.

## File Structure

Documentation files follow the naming pattern `core-{helper-name}.md` for core helpers.

### Frontmatter Format

Each documentation file must start with a frontmatter section:

```yaml
---
name: HelperName
slug: helpername
sourceRef: file.go#L123
category: core
subCategory: slice
signatures:
  - "func HelperName(params) returnType"
  - "func (receiver *Type) MethodName(params) returnType"
  - "func HelperNameI(params) returnType"
  - "func HelperNameWithContext(params) returnType"
playUrl: https://go.dev/play/p/EXAMPLE
variantHelpers:
  - core#slice#helpername
  - core#slice#helpername2
  - core#slice#helpername3
  - core#slice#helpername4
  - core#slice#helpername5
  - core#slice#helpernamei
  - core#slice#helpernamewithcontext
similarHelpers:
  - mutable#slice#helpername
  - core#slice#filterhelpername
  - core#slice#zipx
position: 0
---
```

### Frontmatter Fields

- **name**: The display name of the helper (PascalCase)
- **slug**: URL-friendly name (kebab-case, matches filename without `core-` prefix)
- **sourceRef**: Source file reference with line number (format: `file.go#L123`)
- **category**: `core`, `mutable`, `parallel`, `it`... The category must match the file name.
- **subCategory**: The functional category (e.g., `condition`, `map`, `find`, `slice`...)
- **signatures**: Array of function signatures as strings. Do not list signatures from other sub-packages/category.
- **playUrl**: Go Playground URL with working example
- **variantHelpers**: Array of variant helper names. Must contain at least the default helper named above. This field is for:
  - Variations of the same helper with different signatures (e.g., `Map`, `MapI`, `MapWithContext`, `MapIWithContext`)
  - Helper variants that add functionality like indexes, context, or different parameter types
  - All variants must be in the same category and subcategory as the main helper
  - Examples: `core#slice#map`, `core#slice#mapi`, `core#slice#mapwithcontext`
- **similarHelpers**: Array of related helper names (leave empty if none). This field is for:
  - Equivalent helpers in other packages/categories (e.g., `parallel#slice#Map`, `mutable#slice#Filter`)
  - Helper compositions or related functionality (e.g., `FilterMap` is similar to both `Map` and `Filter`)
  - Helpers with different names but similar purposes (e.g., `FindBy` variants vs base `Find`)
  - Cross-references to helpers that users might want to consider as alternatives
- **position**: Position in the list (0, 10, 20, 30...). Order must follow the order in source code. Helpers are grouped by category+sub-category and displayed on a page. Position number is reset for each page.

## Content Structure

After the frontmatter, include:

1. **Brief description**: One sentence explaining what the helper does
2. **Code example**: Working Go code demonstrating usage
3. **Expected output**: Comment showing the result

```markdown
Brief description of what this helper does. Be concise and not too long.

```go
result := lo.HelperName(example)
// expected result
```
```

Multiple examples can be used for demonstration the helper, such as edge cases. If multiple signatures are grouped under this documentation, it could be useful to describe some (all?) of them.

## Understanding variantHelpers vs similarHelpers

### variantHelpers
Use `variantHelpers` for different versions of the **same helper function**:

**Example**: Map helper variants (all in `core#slice#map`):
```yaml
variantHelpers:
  - core#slice#map        # func Map[T, R]([]T, func(T, int) R) []R
  - core#slice#maperr     # func MapErr[T, R]([]T, func(T, int) (R, error)) ([]R, error)
  - core#slice#mapi       # func MapI[T, R]([]T, func(T, int) R) []R (with index)
  - core#slice#mapwithcontext  # func MapWithContext[T, R]([]T, func(T, int, context.Context) R, context.Context) []R
```

### similarHelpers
Use `similarHelpers` for **related but different helpers**:

**Example**: FilterMap combines Map and Filter functionality:
```yaml
similarHelpers:
  - core#slice#map        # Related transformation helper
  - core#slice#filter     # Related filtering helper
```

**Example**: Cross-package equivalents:
```yaml
similarHelpers:
  - parallel#slice#map    # Parallel version in different package
  - mutable#slice#map     # Mutable version in different package
  - mutable#slice#zipx     # Mutable version in different package
```

**Example**: Related helpers with similar functionality:
```yaml
similarHelpers:
  - core#slice#find       # Single result search
  - core#slice#filter     # Multiple result filtering
  - core#slice#findby     # Key-based search
  - core#slice#findorelse # Search with default value
```

**Example**: Helper families with similar patterns:
```yaml
similarHelpers:
  - core#slice#min        # Find minimum value
  - core#slice#max        # Find maximum value
  - core#slice#minby      # Find minimum by key function
  - core#slice#maxby      # Find maximum by key function
  - core#slice#minindex   # Find minimum value index
  - core#slice#maxindex   # Find maximum value index
```

**Example**: Composition helpers that combine multiple operations:
```yaml
similarHelpers:
  - core#slice#filtermap  # Filter + Map combination
  - core#slice#mapkeys    # Map to extract keys
  - core#slice#mapvalues # Map to extract values
```

When you add similarHelpers to a new helper, please update the linked similar helpers documentation, and add the one you're adding.

Don't link helpers having numeric declensions (eg: use core#slice#zipx instead of core#slice#zip2).

### Key Differences
- **variantHelpers**: Same helper, different signatures/parameters (same package)
- **similarHelpers**: Different helpers, related functionality (can be cross-package)

## Grouping Related Helpers

When multiple helpers operate on the same struct or serve similar purposes, consolidate them into a single file:

**Example**: Map helpers:
- `Map()` base helper
- `MapI()` add index to predicate callback
- `MapWithContext()` add context to predicate callback
- `MapIWithContext()` add index and context to predicate callback

**Example**: Switch helpers all operate on `switchCase[T, R]`:
- `Switch()` - constructor
- `Case()` - method for adding cases
- `CaseF()` - method for adding cases with functions
- `Default()` - method for default values
- `DefaultF()` - method for default values with functions

In such cases:
1. Use the primary helper name in the filename (e.g., `core-switch.md`)
2. Include all related signatures in the `signatures` array
3. List all related helpers in `variantHelpers` array
4. Document each helper in its own section with `### HelperName` headers

## Naming Conventions

### Categories
Use these established subCategories:
- `condition` - conditional logic (if/else, switch)
- `map` - transformation functions
- `find` - search and lookup functions
- `slice` - array manipulation
- `math` - mathematical operations
- `string` - string operations
- `type` - type utilities
- `error-handling` - error management
- `retry` - retry mechanisms
- `time` - time operations
- `function` - function utilities
- `channel` - channel operations
- `tuple` - tuple operations
- `intersect` - set intersections

### Helper Names
- Follow Go naming conventions (PascalCase for exported)
- Use descriptive names that clearly indicate purpose
- For function variants, use consistent suffixes:
  - `F` suffix for function-based versions (lazy evaluation)
  - `I` suffix for variants having `index int` argument in predicate callback
  - `Err` suffix for variants returning an error in predicate callback
  - `WithContext` suffix when context.Context is provided
  - `X` suffix for helpers with varying arguments (eg: MustX: Must2, Must3, Must4...)

## Go Playground Examples

Every helper must have a working Go Playground example linked in two places:

1. **Source code**: `// Play: <url>` comment on the last line of the doc block, right before the `func` keyword
2. **Doc file**: `playUrl: <url>` field in the YAML frontmatter of `docs/data/<category>-<slug>.md`

### Creating a New Playground Example

#### Step 1: Write the Example Code

Write a minimal, self-contained `main.go` that demonstrates the helper. Guidelines:

- Use realistic but simple data
- Print the result with `fmt.Println` so the output is visible
- Include edge cases when useful (e.g., empty input, error case)
- For `Err` variants, show both a success and an error case
- For time-based helpers, use `time.Date()` for deterministic output
- For random helpers (`SampleBy`, `SamplesBy`), use `rand.New(rand.NewSource(42))` for reproducible output

#### Step 2: Import Conventions

Use the correct import path depending on the package:

```go
// Core helpers
import "github.com/samber/lo"
// Usage: lo.Map(...)

// Iterator helpers (it/ package, requires Go 1.23+)
import (
    "slices"
    "github.com/samber/lo/it"
)
// Usage: slices.Collect(it.Map(...))
// Convert slices to iterators: slices.Values([]int{1, 2, 3})

// Parallel helpers
import lop "github.com/samber/lo/parallel"
// Usage: lop.Map(...)
```

#### Step 3: Run and Share via Go Playground

Use the `go-playground` MCP tool to execute the example and get a shareable URL:

```
mcp__go-playground__run_and_share_go_code(code: "<your code>")
```

This compiles the code on go.dev/play, runs it, and returns:
- The program output (to verify correctness)
- A shareable URL like `https://go.dev/play/p/XXXXXXX`

If the output doesn't match expectations, fix the code and re-run until it produces the correct result.

#### Step 4: Add the URL to Source Code

Add a `// Play:` comment as the **last line** of the function's doc comment block:

```go
// Map manipulates a slice and transforms it to a slice of another type.
// Play: https://go.dev/play/p/refNB9ZTIGo
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
```

#### Step 5: Add the URL to Documentation

Set the `playUrl` field in the corresponding `docs/data/*.md` file:

```yaml
---
name: Map
slug: map
playUrl: https://go.dev/play/p/refNB9ZTIGo
...
---
```

### Troubleshooting

**First-run timeouts**: The Go Playground may timeout on the first execution if `github.com/samber/lo` hasn't been cached yet. Simply retry — subsequent runs succeed because the module is cached.

**New helpers not yet released**: If documentation is created at the same time as the helper source code, the Go Playground cannot compile it because the module version hasn't been published yet. In that case, skip the playground example and leave `playUrl` empty. Create the example after the next release.

**SIMD helpers**: Helpers in `exp/simd/` require `go1.26+goexperiment.simd+amd64` build tags, which the Go Playground does not support. These helpers cannot have playground examples.

### Bulk Verification

To verify all playground URLs compile, you can use `mcp__go-playground__execute_go_playground_url` to re-run an existing URL and check the output. To read the source code of an existing playground, use `mcp__go-playground__read_go_playground_url`.

## Example: Complete File

```yaml
---
name: Map
slug: map
sourceRef: map.go#L123
category: core
subCategory: map
signatures:
  - "func Map[T any, R any](collection []T, transform func(item T, index int) R) []R"
playUrl: https://go.dev/play/p/EXAMPLE
similarHelpers: []
position: 0
---

Applies a function to each element of a collection and returns a new collection with the results.

```go
result := lo.Map([]int{1, 2, 3}, func(item int, index int) string {
    return fmt.Sprintf("%d", item)
})
// []string{"1", "2", "3"}
```
```

## Checklist

Before submitting:

- [ ] Frontmatter is complete and correctly formatted
- [ ] Filename matches slug (with `core-` or `mutable-` or `parallel-` or `it-` prefix)
- [ ] Source reference points to correct line number
- [ ] Category and subCategory are appropriate
- [ ] All signatures are included and properly formatted
- [ ] Go Playground example works and demonstrates usage
- [ ] Expected output is shown as a comment
- [ ] Similar helpers are listed if applicable
- [ ] Related helpers are consolidated into single file when appropriate
- [ ] All validation scripts pass without errors
- [ ] Helper is added to llms.txt
