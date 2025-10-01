//go:build go1.23

package it

import (
	"iter"
	"slices"
	"sort"

	"github.com/samber/lo"
	"github.com/samber/lo/internal/constraints"
	"github.com/samber/lo/mutable"
)

// Length returns the length of collection.
func Length[T any](collection iter.Seq[T]) int {
	var count int

	for range collection {
		count++
	}

	return count
}

// Drain consumes an entire sequence.
func Drain[T any](collection iter.Seq[T]) {
	for range collection { //nolint:revive
	}
}

// Filter iterates over elements of collection, returning a sequence of all elements predicate returns true for.
func Filter[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return FilterI(collection, func(item T, _ int) bool { return predicate(item) })
}

// FilterI iterates over elements of collection, returning a sequence of all elements predicate returns true for.
func FilterI[T any, I ~func(func(T) bool)](collection I, predicate func(item T, index int) bool) I {
	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if predicate(item, i) && !yield(item) {
				return
			}
			i++
		}
	}
}

// Map manipulates a sequence and transforms it to a sequence of another type.
func Map[T, R any](collection iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	return MapI(collection, func(item T, _ int) R { return transform(item) })
}

// MapI manipulates a sequence and transforms it to a sequence of another type.
func MapI[T, R any](collection iter.Seq[T], transform func(item T, index int) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		var i int
		for item := range collection {
			if !yield(transform(item, i)) {
				return
			}
			i++
		}
	}
}

// UniqMap manipulates a sequence and transforms it to a sequence of another type with unique values.
func UniqMap[T any, R comparable](collection iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	return Uniq(Map(collection, transform))
}

// UniqMapI manipulates a sequence and transforms it to a sequence of another type with unique values.
func UniqMapI[T any, R comparable](collection iter.Seq[T], transform func(item T, index int) R) iter.Seq[R] {
	return Uniq(MapI(collection, transform))
}

// FilterMap returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return FilterMapI(collection, func(item T, _ int) (R, bool) { return callback(item) })
}

// FilterMapI returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMapI[T, R any](collection iter.Seq[T], callback func(item T, index int) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		var i int
		for item := range collection {
			if r, ok := callback(item, i); ok && !yield(r) {
				return
			}
			i++
		}
	}
}

// FlatMap manipulates a sequence and transforms and flattens it to a sequence of another type.
// The transform function can either return a sequence or a `nil`, and in the `nil` case
// no value is yielded.
func FlatMap[T, R any](collection iter.Seq[T], transform func(item T) iter.Seq[R]) iter.Seq[R] {
	return FlatMapI(collection, func(item T, _ int) iter.Seq[R] { return transform(item) })
}

// FlatMapI manipulates a sequence and transforms and flattens it to a sequence of another type.
// The transform function can either return a sequence or a `nil`, and in the `nil` case
// no value is yielded.
func FlatMapI[T, R any](collection iter.Seq[T], transform func(item T, index int) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		var i int
		for item := range collection {
			for r := range transform(item, i) {
				if !yield(r) {
					return
				}
			}
			i++
		}
	}
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	return ReduceI(collection, func(agg R, item T, _ int) R { return accumulator(agg, item) }, initial)
}

// ReduceI reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func ReduceI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	var i int
	for item := range collection {
		initial = accumulator(initial, item, i)
		i++
	}

	return initial
}

// ReduceLast is like Reduce except that it iterates over elements of collection in reverse.
func ReduceLast[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	return Reduce(Reverse(collection), accumulator, initial)
}

// ReduceLastI is like Reduce except that it iterates over elements of collection in reverse.
func ReduceLastI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	return ReduceI(Reverse(collection), accumulator, initial)
}

// ForEach iterates over elements of collection and invokes transform for each element.
func ForEach[T any](collection iter.Seq[T], transform func(item T)) {
	ForEachI(collection, func(item T, _ int) { transform(item) })
}

// ForEachI iterates over elements of collection and invokes transform for each element.
func ForEachI[T any](collection iter.Seq[T], transform func(item T, index int)) {
	var i int
	for item := range collection {
		transform(item, i)
		i++
	}
}

// ForEachWhile iterates over elements of collection and invokes predicate for each element
// collection return value decide to continue or break, like do while().
func ForEachWhile[T any](collection iter.Seq[T], predicate func(item T) bool) {
	ForEachWhileI(collection, func(item T, _ int) bool { return predicate(item) })
}

// ForEachWhileI iterates over elements of collection and invokes predicate for each element
// collection return value decide to continue or break, like do while().
func ForEachWhileI[T any](collection iter.Seq[T], predicate func(item T, index int) bool) {
	var i int
	for item := range collection {
		if !predicate(item, i) {
			return
		}
		i++
	}
}

// Times invokes transform n times and returns a sequence of results.
// The transform is invoked with index as argument.
func Times[T any](count int, transform func(index int) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(transform(i)) {
				return
			}
		}
	}
}

// Uniq returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence.
func Uniq[T comparable, I ~func(func(T) bool)](collection I) I {
	return UniqBy(collection, func(item T) T { return item })
}

// UniqBy returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
func UniqBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	return func(yield func(T) bool) {
		seen := make(map[U]struct{})

		for item := range collection {
			key := transform(item)

			if _, ok := seen[key]; !ok {
				if !yield(item) {
					return
				}
				seen[key] = struct{}{}
			}
		}
	}
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through transform.
func GroupBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U][]T {
	return GroupByMap(collection, func(item T) (U, T) { return transform(item), item })
}

// GroupByMap returns an object composed of keys generated from the results of running each element of collection through transform.
func GroupByMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K][]V {
	result := make(map[K][]V)

	for item := range collection {
		k, v := transform(item)

		result[k] = append(result[k], v)
	}

	return result
}

// Chunk returns a sequence of elements split into groups of length size. If the sequence can't be split evenly,
// the final chunk will be the remaining elements.
func Chunk[T any](collection iter.Seq[T], size int) iter.Seq[[]T] {
	if size <= 0 {
		panic("it.Chunk: size must be greater than 0")
	}

	return func(yield func([]T) bool) {
		newSlice := make([]T, 0, size)
		for item := range collection {
			newSlice = append(newSlice, item)
			if len(newSlice) == size {
				if !yield(newSlice) {
					return
				}
				newSlice = make([]T, 0, size)
			}
		}
		if len(newSlice) > 0 {
			yield(newSlice)
		}
	}
}

// PartitionBy returns a sequence of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through transform.
func PartitionBy[T any, K comparable](collection iter.Seq[T], transform func(item T) K) [][]T {
	var result [][]T
	seen := map[K]int{}

	for item := range collection {
		key := transform(item)

		resultIndex, ok := seen[key]
		if !ok {
			resultIndex = len(result)
			seen[key] = resultIndex
			result = append(result, []T{})
		}

		result[resultIndex] = append(result[resultIndex], item)
	}

	return result
}

// Flatten returns a sequence a single level deep.
func Flatten[T any, I ~func(func(T) bool)](collection []I) I {
	return func(yield func(T) bool) {
		for _, item := range collection {
			for item := range item {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Interleave round-robin alternating input sequences and sequentially appending value at index into result.
func Interleave[T any](collections ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next := make([]func() (T, bool), len(collections))
		for i, c := range collections {
			var stop func()
			next[i], stop = iter.Pull(c)
			defer stop()
		}
		var done int
		for done < len(next) {
			done = 0
			for i, n := range next {
				if n == nil {
					done++
				} else if t, ok := n(); !ok {
					next[i] = nil
					done++
				} else if !yield(t) {
					return
				}
			}
		}
	}
}

// Shuffle returns a sequence of shuffled values. Uses the Fisher-Yates shuffle algorithm.
func Shuffle[T any, I ~func(func(T) bool)](collection I) I {
	slice := slices.Collect(iter.Seq[T](collection))
	mutable.Shuffle(slice)
	return I(slices.Values(slice))
}

// Reverse reverses a sequence so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any, I ~func(func(T) bool)](collection I) I {
	slice := slices.Collect(iter.Seq[T](collection))
	mutable.Reverse(slice)
	return I(slices.Values(slice))
}

// Fill replaces elements of a sequence with `initial` value.
func Fill[T lo.Clonable[T], I ~func(func(T) bool)](collection I, initial T) I {
	return func(yield func(T) bool) {
		for range collection {
			if !yield(initial.Clone()) {
				return
			}
		}
	}
}

// Repeat builds a sequence with N copies of initial value.
func Repeat[T lo.Clonable[T]](count int, initial T) iter.Seq[T] {
	return RepeatBy(count, func(int) T { return initial.Clone() })
}

// RepeatBy builds a sequence with values returned by N calls of transform.
func RepeatBy[T any](count int, transform func(index int) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range count {
			if !yield(transform(i)) {
				return
			}
		}
	}
}

// KeyBy transforms a sequence to a map based on a pivot transform function.
func KeyBy[K comparable, V any](collection iter.Seq[V], transform func(item V) K) map[K]V {
	result := make(map[K]V)

	for item := range collection {
		k := transform(item)
		result[k] = item
	}

	return result
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
func Associate[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V {
	result := make(map[K]V)

	for item := range collection {
		k, v := transform(item)
		result[k] = v
	}

	return result
}

// SeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Alias of Associate().
func SeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V {
	return Associate(collection, transform)
}

// FilterSeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// The third return value of the transform function is a boolean that indicates whether the key-value pair should be included in the map.
func FilterSeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V, bool)) map[K]V {
	result := make(map[K]V)

	for item := range collection {
		if k, v, ok := transform(item); ok {
			result[k] = v
		}
	}

	return result
}

// Keyify returns a map with each unique element of the sequence as a key.
func Keyify[T comparable](collection iter.Seq[T]) map[T]struct{} {
	result := make(map[T]struct{})

	for item := range collection {
		result[item] = struct{}{}
	}

	return result
}

// Drop drops n elements from the beginning of a sequence.
func Drop[T any, I ~func(func(T) bool)](collection I, n int) I {
	if n < 0 {
		panic("it.Drop: n must not be negative")
	}

	if n == 0 {
		return collection
	}

	return FilterI(collection, func(item T, index int) bool { return index >= n })
}

// DropLast drops n elements from the end of a sequence.
func DropLast[T any, I ~func(func(T) bool)](collection I, n int) I {
	if n < 0 {
		panic("it.DropLast: n must not be negative")
	}

	if n == 0 {
		return collection
	}

	return func(yield func(T) bool) {
		buf := make([]T, 0, n)
		var i int
		for item := range collection {
			if len(buf) < n {
				buf = append(buf, item)
			} else {
				if !yield(buf[i]) {
					return
				}
				buf[i] = item
				i = (i + 1) % n
			}
		}
	}
}

// DropWhile drops elements from the beginning of a sequence while the predicate returns true.
func DropWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return func(yield func(T) bool) {
		dropping := true
		for item := range collection {
			if dropping && !predicate(item) {
				dropping = false
			}
			if !dropping && !yield(item) {
				return
			}
		}
	}
}

// DropLastWhile drops elements from the end of a sequence while the predicate returns true.
func DropLastWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return func(yield func(T) bool) {
		var buf []T
		for item := range collection {
			if predicate(item) {
				buf = append(buf, item)
				continue
			}
			if len(buf) > 0 {
				for _, item := range buf {
					if !yield(item) {
						return
					}
				}
				buf = buf[:0]
			}
			if !yield(item) {
				return
			}
		}
	}
}

// DropByIndex drops elements from a sequence by the index.
func DropByIndex[T any, I ~func(func(T) bool)](collection I, indexes ...int) I {
	return func(yield func(T) bool) {
		indexes = lo.Filter(lo.Uniq(indexes), func(item, _ int) bool { return item >= 0 })
		sort.Ints(indexes)

		var i int
		for item := range collection {
			if len(indexes) > 0 && indexes[0] == i {
				indexes = indexes[1:]
			} else if !yield(item) {
				return
			}
			i++
		}
	}
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.
func Reject[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return RejectI(collection, func(item T, _ int) bool { return predicate(item) })
}

// RejectI is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.
func RejectI[T any, I ~func(func(T) bool)](collection I, predicate func(item T, index int) bool) I {
	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if !predicate(item, i) && !yield(item) {
				return
			}
			i++
		}
	}
}

// RejectMap is the opposite of FilterMap, this method returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func RejectMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return RejectMapI(collection, func(item T, _ int) (R, bool) { return callback(item) })
}

// RejectMapI is the opposite of FilterMap, this method returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func RejectMapI[T, R any](collection iter.Seq[T], callback func(item T, index int) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		var i int
		for item := range collection {
			if r, ok := callback(item, i); !ok && !yield(r) {
				return
			}
			i++
		}
	}
}

// Count counts the number of elements in the collection that equal value.
func Count[T comparable](collection iter.Seq[T], value T) int {
	return CountBy(collection, func(item T) bool { return item == value })
}

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection iter.Seq[T], predicate func(item T) bool) int {
	var count int

	for range Filter(collection, predicate) {
		count++
	}

	return count
}

// CountValues counts the number of each element in the collection.
func CountValues[T comparable](collection iter.Seq[T]) map[T]int {
	return CountValuesBy(collection, func(item T) T { return item })
}

// CountValuesBy counts the number of each element returned from transform function.
// Is equivalent to chaining Map and CountValues.
func CountValuesBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U]int {
	result := make(map[U]int)

	for item := range collection {
		result[transform(item)]++
	}

	return result
}

// Subset returns a subset of a sequence from `offset` up to `length` elements.
func Subset[T any, I ~func(func(T) bool)](collection I, offset, length int) I {
	if offset < 0 {
		panic("it.Subset: offset must not be negative")
	}
	if length < 0 {
		panic("it.Subset: length must not be negative")
	}

	return Slice(collection, offset, offset+length)
}

// Slice returns a subset of a sequence from `start` up to, but not including `end`.
func Slice[T any, I ~func(func(T) bool)](collection I, start, end int) I {
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}

	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if i >= start && (i >= end || !yield(item)) {
				return
			}
			i++
		}
	}
}

// Replace returns a sequence with the first n non-overlapping instances of old replaced by new.
func Replace[T comparable, I ~func(func(T) bool)](collection I, old, nEw T, n int) I {
	return I(Map(iter.Seq[T](collection), func(item T) T {
		if n != 0 && item == old {
			n--
			return nEw
		}
		return item
	}))
}

// ReplaceAll returns a sequence with all non-overlapping instances of old replaced by new.
func ReplaceAll[T comparable, I ~func(func(T) bool)](collection I, old, nEw T) I {
	return Replace(collection, old, nEw, -1)
}

// Compact returns a sequence of all non-zero elements.
func Compact[T comparable, I ~func(func(T) bool)](collection I) I {
	return Filter(collection, lo.IsNotEmpty)
}

// IsSorted checks if a sequence is sorted.
func IsSorted[T constraints.Ordered](collection iter.Seq[T]) bool {
	return IsSortedBy(collection, func(item T) T { return item })
}

// IsSortedBy checks if a sequence is sorted by transform.
func IsSortedBy[T any, K constraints.Ordered](collection iter.Seq[T], transform func(item T) K) bool {
	first := true
	var prev K
	for item := range collection {
		key := transform(item)
		if first {
			first = false
		} else if prev > key {
			return false
		}
		prev = key
	}
	return true
}

// Splice inserts multiple elements at index i. The helper is protected against overflow errors.
func Splice[T any, I ~func(func(T) bool)](collection I, index int, elements ...T) I {
	if index < 0 {
		panic("it.Splice: index must not be negative")
	}

	if len(elements) == 0 {
		return collection
	}

	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if i == index {
				for _, element := range elements {
					if !yield(element) {
						return
					}
				}
			}
			if !yield(item) {
				return
			}
			i++
		}
		if i <= index {
			for _, element := range elements {
				if !yield(element) {
					return
				}
			}
		}
	}
}

// CutPrefix returns collection without the provided leading prefix
// and reports whether it found the prefix.
// If collection doesn't start with prefix, CutPrefix returns collection, false.
// If prefix is empty, CutPrefix returns collection, true.
func CutPrefix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (after I, found bool) { //nolint:gocyclo
	if len(separator) == 0 {
		return collection, true
	}

	next, stop := iter.Pull(iter.Seq[T](collection))
	for i := range separator {
		item, ok := next()
		if !ok {
			return func(yield func(T) bool) {
				defer stop()
				for j := 0; j < i; j++ {
					if !yield(separator[j]) {
						return
					}
				}
			}, false
		}

		if item != separator[i] {
			return func(yield func(T) bool) {
				defer stop()
				for j := 0; j < i; j++ {
					if !yield(separator[j]) {
						return
					}
				}
				if ok && !yield(item) {
					return
				}
				for {
					if item, ok := next(); !ok || !yield(item) {
						return
					}
				}
			}, false
		}
	}

	return func(yield func(T) bool) {
		defer stop()
		for {
			if item, ok := next(); !ok || !yield(item) {
				return
			}
		}
	}, true
}

// CutSuffix returns collection without the provided ending suffix and reports
// whether it found the suffix. If collection doesn't end with suffix, CutSuffix returns collection, false.
// If suffix is empty, CutSuffix returns collection, true.
func CutSuffix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (before I, found bool) {
	slice := slices.Collect(iter.Seq[T](collection))
	result, ok := lo.CutSuffix(slice, separator)
	return I(slices.Values(result)), ok
}
