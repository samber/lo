---
name: Associate
slug: associate
sourceRef: it/seq.go#L412
category: it
subCategory: map
signatures:
  - "func Associate[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V"
  - "func AssociateI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V)) map[K]V"
  - "func SeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V"
  - "func SeqToMapI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V)) map[K]V"
  - "func FilterSeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V, bool)) map[K]V"
  - "func FilterSeqToMapI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V, bool)) map[K]V"
variantHelpers:
  - it#sequence#associate
  - it#sequence#associatei
  - it#sequence#seqtomap
  - it#sequence#seqtomapi
  - it#sequence#filterseqtomap
  - it#sequence#filterseqtomapi
similarHelpers:
  - core#slice#associate
  - core#map#keyby
  - it#map#keyby
position: 140
---

Associate returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
If any of two pairs have the same key the last one gets added to the map.

```go
collection := func(yield func(string) bool) {
    yield("apple")
    yield("banana")
    yield("cherry")
}

result := it.Associate(collection, func(s string) (string, int) {
    return s, len(s)
})
// result contains {"apple": 5, "banana": 6, "cherry": 6}
```

AssociateI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence, with index.

```go
collection := func(yield func(string) bool) {
    yield("a")
    yield("b")
    yield("c")
}

result := it.AssociateI(collection, func(item string, index int) (int, string) {
    return index, item
})
// result contains {0: "a", 1: "b", 2: "c"}
```

SeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
Alias of Associate().

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
}

result := it.SeqToMap(collection, func(i int) (int, string) {
    return i, fmt.Sprintf("item-%d", i)
})
// result contains {1: "item-1", 2: "item-2", 3: "item-3"}
```

FilterSeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
The third return value of the transform function is a boolean that indicates whether the key-value pair should be included in the map.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}

result := it.FilterSeqToMap(collection, func(i int) (int, string, bool) {
    return i, fmt.Sprintf("item-%d", i), i%2 == 0
})
// result contains {2: "item-2", 4: "item-4"}
```