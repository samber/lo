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
  - `WithContext` suffix when context.Context is provided
  - `X` suffix for helpers with varying arguments (eg: MustX: Must2, Must3, Must4...)

## Go Playground Examples

Every helper must have a working Go Playground example:
1. Create a minimal, self-contained example
2. Use realistic but simple data
3. Include the expected result as a comment
4. Test the example to ensure it works

When creating the go playground examples, please run it to be sure it compiles and returns the expected output. If invalid, loop until it works.

Add these examples in the source code comments, on top of helpers, with a syntax like `// Play: <url>`.

If the documentation is created at the same time of the helper source code, then the Go playground execution might fail, since we need to merge+release the source code first to make this new helper available to Go playground compiler. In that case, skip the creation of the example and set no URL.

## Validation Scripts

Run these scripts to validate your documentation:

```bash
# Check cross-references
node scripts/check-cross-references.js

# Check for duplicates in categories
node scripts/check-duplicates-in-category.js

# Check filename matches frontmatter
node scripts/check-filename-matches-frontmatter.js

# Check for similar existing helpers
node scripts/check-similar-exists.js

# Check for similar keys in directory
node scripts/check-similar-keys-exist-in-directory.js
```

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
