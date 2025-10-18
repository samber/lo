---
name: None
slug: none
sourceRef: it/intersect.go#L63
category: it
subCategory: intersect
signatures:
  - "func None[T comparable](collection iter.Seq[T], subset ...T) bool"
playUrl: "https://go.dev/play/p/KmX-fXictQl"
variantHelpers:
  - it#intersect#none
similarHelpers:
  - core#slice#none
position: 680
---

Returns true if no element of a subset is contained in a collection or if the subset is empty.

Will iterate through the entire sequence if subset elements never match.

Examples:

```go
// Check if collection contains none of the forbidden values
numbers := it.Slice([]int{1, 3, 5, 7, 9})
forbidden := []int{2, 4, 6, 8}
hasNone := it.None(numbers, forbidden...)
// hasNone: true

numbers = it.Slice([]int{1, 3, 5, 8, 9})
hasNone = it.None(numbers, forbidden...)
// hasNone: false (8 is in both collection and forbidden)

// Check if collection contains none of unwanted words
words := it.Slice([]string{"hello", "world", "go", "lang"})
unwanted := []string{"bad", "evil", "wrong"}
hasNone := it.None(words, unwanted...)
// hasNone: true

words = it.Slice([]string{"hello", "bad", "go", "lang"})
hasNone = it.None(words, unwanted...)
// hasNone: false ("bad" is in both)

// Check if collection contains none of specific IDs
ids := it.Slice([]int{101, 102, 103, 104})
restrictedIds := []int{201, 202, 203}
hasNone := it.None(ids, restrictedIds...)
// hasNone: true

ids = it.Slice([]int{101, 102, 203, 104})
hasNone = it.None(ids, restrictedIds...)
// hasNone: false (203 is restricted)

// Check with empty subset (always returns true)
numbers = it.Slice([]int{1, 3, 5, 7, 9})
hasNone = it.None(numbers)
// hasNone: true

// Check with strings containing specific characters
words := it.Slice([]string{"hello", "world", "go", "lang"})
forbiddenChars := []string{"@", "#", "$"}
hasNone := it.None(words, forbiddenChars...)
// hasNone: true

words = it.Slice([]string{"hello", "world", "go@"})
hasNone = it.None(words, forbiddenChars...)
// hasNone: false (contains "@")

// Check if collection has none of problematic status codes
statusCodes := it.Slice([]int{200, 201, 204})
errorCodes := []int{400, 401, 403, 404, 500}
hasNone := it.None(statusCodes, errorCodes...)
// hasNone: true

statusCodes = it.Slice([]int{200, 404, 204})
hasNone = it.None(statusCodes, errorCodes...)
// hasNone: false (contains 404)

// Check with empty collection (always returns true)
empty := it.Slice([]int{})
hasNone = it.None(empty, 1, 2, 3)
// hasNone: true

// Check for none of forbidden usernames
usernames := it.Slice([]string{"alice", "bob", "charlie"})
forbiddenUsers := []string{"admin", "root", "system"}
hasNone := it.None(usernames, forbiddenUsers...)
// hasNone: true

usernames = it.Slice([]string{"alice", "admin", "charlie"})
hasNone = it.None(usernames, forbiddenUsers...)
// hasNone: false ("admin" is forbidden)
```