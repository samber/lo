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

// Filter iterates over elements of collection, returning a sequence of all elements predicate returns true for.
func Filter[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return func(yield func(T) bool) {
		for item := range collection {
			if predicate(item) && !yield(item) {
				return
			}
		}
	}
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
func Map[T, R any](collection iter.Seq[T], iteratee func(item T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range collection {
			if !yield(iteratee(item)) {
				return
			}
		}
	}
}

// MapI manipulates a sequence and transforms it to a sequence of another type.
func MapI[T, R any](collection iter.Seq[T], iteratee func(item T, index int) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		var i int
		for item := range collection {
			if !yield(iteratee(item, i)) {
				return
			}
			i++
		}
	}
}

// UniqMap manipulates a sequence and transforms it to a sequence of another type with unique values.
func UniqMap[T any, R comparable](collection iter.Seq[T], iteratee func(item T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		seen := make(map[R]struct{})
		for item := range collection {
			r := iteratee(item)
			if _, ok := seen[r]; !ok {
				if !yield(r) {
					return
				}
				seen[r] = struct{}{}
			}
		}
	}
}

// UniqMapI manipulates a sequence and transforms it to a sequence of another type with unique values.
func UniqMapI[T any, R comparable](collection iter.Seq[T], iteratee func(item T, index int) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		seen := make(map[R]struct{})
		var i int
		for item := range collection {
			r := iteratee(item, i)
			if _, ok := seen[r]; !ok {
				if !yield(r) {
					return
				}
				seen[r] = struct{}{}
			}
			i++
		}
	}
}

// FilterMap returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range collection {
			if r, ok := callback(item); ok && !yield(r) {
				return
			}
		}
	}
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
func FlatMap[T, R any](collection iter.Seq[T], iteratee func(item T) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range collection {
			for r := range iteratee(item) {
				if !yield(r) {
					return
				}
			}
		}
	}
}

// FlatMapI manipulates a sequence and transforms and flattens it to a sequence of another type.
// The transform function can either return a sequence or a `nil`, and in the `nil` case
// no value is yielded.
func FlatMapI[T, R any](collection iter.Seq[T], iteratee func(item T, index int) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		var i int
		for item := range collection {
			for r := range iteratee(item, i) {
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
	for item := range collection {
		initial = accumulator(initial, item)
	}

	return initial
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

// ReduceRight is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	return Reduce(Reverse(collection), accumulator, initial)
}

// ReduceRightI is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRightI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	return ReduceI(Reverse(collection), accumulator, initial)
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection iter.Seq[T], iteratee func(item T)) {
	var i int
	for item := range collection {
		iteratee(item)
		i++
	}
}

// ForEachI iterates over elements of collection and invokes iteratee for each element.
func ForEachI[T any](collection iter.Seq[T], iteratee func(item T, index int)) {
	var i int
	for item := range collection {
		iteratee(item, i)
		i++
	}
}

// ForEachWhile iterates over elements of collection and invokes iteratee for each element
// collection return value decide to continue or break, like do while().
func ForEachWhile[T any](collection iter.Seq[T], iteratee func(item T) bool) {
	for item := range collection {
		if !iteratee(item) {
			return
		}
	}
}

// ForEachWhileI iterates over elements of collection and invokes iteratee for each element
// collection return value decide to continue or break, like do while().
func ForEachWhileI[T any](collection iter.Seq[T], iteratee func(item T, index int) bool) {
	var i int
	for item := range collection {
		if !iteratee(item, i) {
			return
		}
		i++
	}
}

// Times invokes the iteratee n times, returning a sequence of the results of each invocation.
// The iteratee is invoked with index as argument.
func Times[T any](count int, iteratee func(index int) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(iteratee(i)) {
				return
			}
		}
	}
}

// Uniq returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence.
func Uniq[T comparable, I ~func(func(T) bool)](collection I) I {
	return func(yield func(T) bool) {
		seen := make(map[T]struct{})

		for item := range collection {
			if _, ok := seen[item]; ok {
				continue
			}

			seen[item] = struct{}{}
			if !yield(item) {
				return
			}
		}
	}
}

// UniqBy returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence. It accepts `iteratee` which is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
func UniqBy[T any, U comparable, I ~func(func(T) bool)](collection I, iteratee func(item T) U) I {
	return func(yield func(T) bool) {
		seen := make(map[U]struct{})

		for item := range collection {
			key := iteratee(item)

			if _, ok := seen[key]; ok {
				continue
			}

			seen[key] = struct{}{}
			if !yield(item) {
				return
			}
		}
	}
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection iter.Seq[T], iteratee func(item T) U) map[U][]T {
	result := map[U][]T{}

	for item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}

	return result
}

// GroupByMap returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupByMap[T any, K comparable, V any](collection iter.Seq[T], iteratee func(item T) (K, V)) map[K][]V {
	result := map[K][]V{}

	for item := range collection {
		k, v := iteratee(item)

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
// of running each element of collection through iteratee.
func PartitionBy[T any, K comparable](collection iter.Seq[T], iteratee func(item T) K) [][]T {
	result := [][]T{}
	seen := map[K]int{}

	for item := range collection {
		key := iteratee(item)

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
	return func(yield func(T) bool) {
		for range count {
			if !yield(initial.Clone()) {
				return
			}
		}
	}
}

// RepeatBy builds a sequence with values returned by N calls of callback.
func RepeatBy[T any](count int, predicate func(index int) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range count {
			if !yield(predicate(i)) {
				return
			}
		}
	}
}

// KeyBy transforms a sequence to a map based on a pivot callback.
func KeyBy[K comparable, V any](collection iter.Seq[V], iteratee func(item V) K) map[K]V {
	result := make(map[K]V)

	for item := range collection {
		k := iteratee(item)
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
		k, v, ok := transform(item)
		if ok {
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

	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if i >= n && !yield(item) {
				return
			}
			i++
		}
	}
}

// DropRight drops n elements from the end of a sequence.
func DropRight[T any, I ~func(func(T) bool)](collection I, n int) I {
	if n < 0 {
		panic("it.DropRight: n must not be negative")
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

// DropRightWhile drops elements from the end of a sequence while the predicate returns true.
func DropRightWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
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
	return func(yield func(T) bool) {
		for item := range collection {
			if !predicate(item) && !yield(item) {
				return
			}
		}
	}
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
	return func(yield func(R) bool) {
		for item := range collection {
			if r, ok := callback(item); !ok && !yield(r) {
				return
			}
		}
	}
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
	var count int

	for item := range collection {
		if item == value {
			count++
		}
	}

	return count
}

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection iter.Seq[T], predicate func(item T) bool) int {
	var count int

	for item := range collection {
		if predicate(item) {
			count++
		}
	}

	return count
}

// CountValues counts the number of each element in the collection.
func CountValues[T comparable](collection iter.Seq[T]) map[T]int {
	result := make(map[T]int)

	for item := range collection {
		result[item]++
	}

	return result
}

// CountValuesBy counts the number of each element returned from mapper function.
// Is equivalent to chaining Map and CountValues.
func CountValuesBy[T any, U comparable](collection iter.Seq[T], mapper func(item T) U) map[U]int {
	result := make(map[U]int)

	for item := range collection {
		result[mapper(item)]++
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

	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if i >= offset && i < offset+length && !yield(item) {
				return
			}
			i++
		}
	}
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
			if i >= start && i < end && !yield(item) {
				return
			}
			i++
		}
	}
}

// Replace returns a sequence with the first n non-overlapping instances of old replaced by new.
func Replace[T comparable, I ~func(func(T) bool)](collection I, old, nEw T, n int) I {
	return func(yield func(T) bool) {
		for item := range collection {
			if n == 0 || item != old {
				if !yield(item) {
					return
				}
			} else {
				if !yield(nEw) {
					return
				}
				n--
			}
		}
	}
}

// ReplaceAll returns a sequence with all non-overlapping instances of old replaced by new.
func ReplaceAll[T comparable, I ~func(func(T) bool)](collection I, old, nEw T) I {
	return Replace(collection, old, nEw, -1)
}

// Compact returns a sequence of all non-zero elements.
func Compact[T comparable, I ~func(func(T) bool)](collection I) I {
	return func(yield func(T) bool) {
		var zero T

		for item := range collection {
			if item != zero && !yield(item) {
				return
			}
		}
	}
}

// IsSorted checks if a sequence is sorted.
func IsSorted[T constraints.Ordered](collection iter.Seq[T]) bool {
	first := true
	var prev T
	for item := range collection {
		if first {
			first = false
		} else if prev > item {
			return false
		}
		prev = item
	}
	return true
}

// IsSortedByKey checks if a sequence is sorted by iteratee.
func IsSortedByKey[T any, K constraints.Ordered](collection iter.Seq[T], iteratee func(item T) K) bool {
	first := true
	var prev K
	for item := range collection {
		key := iteratee(item)
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
