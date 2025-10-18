//go:build go1.23

package it

import (
	"iter"
	"slices"

	"github.com/samber/lo"
	"github.com/samber/lo/internal/constraints"
	"github.com/samber/lo/mutable"
)

// Length returns the length of collection.
// Will iterate through the entire sequence.
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
// Play: https://go.dev/play/p/psenko2KKsX
func Filter[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return FilterI(collection, func(item T, _ int) bool { return predicate(item) })
}

// FilterI iterates over elements of collection, returning a sequence of all elements predicate returns true for.
// Play: https://go.dev/play/p/5fpdlQvdL-q
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
// Play: https://go.dev/play/p/rWZiPB-RZOo
func Map[T, R any](collection iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	return MapI(collection, func(item T, _ int) R { return transform(item) })
}

// MapI manipulates a sequence and transforms it to a sequence of another type.
// Play: https://go.dev/play/p/6gqemRweL-r
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
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/7hrfnSxfL-s
func UniqMap[T any, R comparable](collection iter.Seq[T], transform func(item T) R) iter.Seq[R] {
	return Uniq(Map(collection, transform))
}

// UniqMapI manipulates a sequence and transforms it to a sequence of another type with unique values.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/8isgTsyfL-t
func UniqMapI[T any, R comparable](collection iter.Seq[T], transform func(item T, index int) R) iter.Seq[R] {
	return Uniq(MapI(collection, transform))
}

// FilterMap returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/9jthUzgF-u
func FilterMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return FilterMapI(collection, func(item T, _ int) (R, bool) { return callback(item) })
}

// FilterMapI returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/0XrQKOk-vw
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
// Play: https://go.dev/play/p/1YsRLPl-wx
func FlatMap[T, R any](collection iter.Seq[T], transform func(item T) iter.Seq[R]) iter.Seq[R] {
	return FlatMapI(collection, func(item T, _ int) iter.Seq[R] { return transform(item) })
}

// FlatMapI manipulates a sequence and transforms and flattens it to a sequence of another type.
// The transform function can either return a sequence or a `nil`, and in the `nil` case
// no value is yielded.
// Play: https://go.dev/play/p/2ZtSMQm-xy
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
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/FmkVUf39ZP_Y
func Reduce[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	return ReduceI(collection, func(agg R, item T, _ int) R { return accumulator(agg, item) }, initial)
}

// ReduceI reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/3AuTNRn-yz
func ReduceI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	var i int
	for item := range collection {
		initial = accumulator(initial, item, i)
		i++
	}

	return initial
}

// ReduceLast is like Reduce except that it iterates over elements of collection in reverse.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/4BvOSo-yza
func ReduceLast[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	return Reduce(Reverse(collection), accumulator, initial)
}

// ReduceLastI is like Reduce except that it iterates over elements of collection in reverse.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/5CwPTPz-zb
func ReduceLastI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R {
	return ReduceI(Reverse(collection), accumulator, initial)
}

// ForEach iterates over elements of collection and invokes transform for each element.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/agIsKpG-S-P
func ForEach[T any](collection iter.Seq[T], transform func(item T)) {
	ForEachI(collection, func(item T, _ int) { transform(item) })
}

// ForEachI iterates over elements of collection and invokes transform for each element.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/6NhAE0-zm
func ForEachI[T any](collection iter.Seq[T], transform func(item T, index int)) {
	var i int
	for item := range collection {
		transform(item, i)
		i++
	}
}

// ForEachWhile iterates over elements of collection and invokes predicate for each element
// collection return value decide to continue or break, like do while().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/7OiBF1-zn
func ForEachWhile[T any](collection iter.Seq[T], predicate func(item T) bool) {
	ForEachWhileI(collection, func(item T, _ int) bool { return predicate(item) })
}

// ForEachWhileI iterates over elements of collection and invokes predicate for each element
// collection return value decide to continue or break, like do while().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/8PjCG2-zo
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
// Play: https://go.dev/play/p/9QkDH3-zp
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
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/0RlEI4-zq
func Uniq[T comparable, I ~func(func(T) bool)](collection I) I {
	return UniqBy(collection, func(item T) T { return item })
}

// UniqBy returns a duplicate-free version of a sequence, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/1SmFJ5-zr
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
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/2TnGK6-zs
func GroupBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U][]T {
	return GroupByMap(collection, func(item T) (U, T) { return transform(item), item })
}

// GroupByMap returns an object composed of keys generated from the results of running each element of collection through transform.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/3UoHL7-zt
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
// Play: https://go.dev/play/p/4VpIM8-zu
func Chunk[T any](collection iter.Seq[T], size int) iter.Seq[[]T] {
	if size <= 0 {
		panic("it.Chunk: size must be greater than 0")
	}

	return func(yield func([]T) bool) {
		var newSlice []T
		for item := range collection {
			if newSlice == nil {
				newSlice = make([]T, 0, size)
			}
			newSlice = append(newSlice, item)
			if len(newSlice) == size {
				if !yield(newSlice) {
					return
				}
				newSlice = nil
			}
		}
		if newSlice != nil {
			yield(newSlice)
		}
	}
}

// PartitionBy returns a sequence of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through transform.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/5WqJN9-zv
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
// Play: https://go.dev/play/p/6XrKO0-zw
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
// Will allocate a slice the size of collections.
// Play: https://go.dev/play/p/7YsLP1-zx
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
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/8isgTsyfL-t
func Shuffle[T any, I ~func(func(T) bool)](collection I) I {
	slice := slices.Collect(iter.Seq[T](collection))
	mutable.Shuffle(slice)
	return I(slices.Values(slice))
}

// Reverse reverses a sequence so that the first element becomes the last, the second element becomes the second to last, and so on.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/9jthUzgF-u
func Reverse[T any, I ~func(func(T) bool)](collection I) I {
	slice := slices.Collect(iter.Seq[T](collection))
	mutable.Reverse(slice)
	return I(slices.Values(slice))
}

// Fill replaces elements of a sequence with `initial` value.
// Play: https://go.dev/play/p/0XrQKOk-vw
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
// Play: https://go.dev/play/p/1YsRLPl-wx
func Repeat[T lo.Clonable[T]](count int, initial T) iter.Seq[T] {
	return RepeatBy(count, func(int) T { return initial.Clone() })
}

// RepeatBy builds a sequence with values returned by N calls of transform.
// Play: https://go.dev/play/p/2ZtSMQm-xy
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
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/3AuTNRn-yz
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
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/4BvOSo-yza
func Associate[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V {
	return AssociateI(collection, func(item T, _ int) (K, V) { return transform(item) })
}

// AssociateI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/5CwPTPz-zb
func AssociateI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V)) map[K]V {
	result := make(map[K]V)

	var i int
	for item := range collection {
		k, v := transform(item, i)
		result[k] = v
		i++
	}

	return result
}

// SeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Alias of Associate().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/6NhAE0-zm
func SeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V)) map[K]V {
	return Associate(collection, transform)
}

// SeqToMapI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// Alias of AssociateI().
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/7OiBF1-zn
func SeqToMapI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V)) map[K]V {
	return AssociateI(collection, transform)
}

// FilterSeqToMap returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// The third return value of the transform function is a boolean that indicates whether the key-value pair should be included in the map.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/8PjCG2-zo
func FilterSeqToMap[T any, K comparable, V any](collection iter.Seq[T], transform func(item T) (K, V, bool)) map[K]V {
	return FilterSeqToMapI(collection, func(item T, _ int) (K, V, bool) {
		return transform(item)
	})
}

// FilterSeqToMapI returns a map containing key-value pairs provided by transform function applied to elements of the given sequence.
// If any of two pairs have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original sequence.
// The third return value of the transform function is a boolean that indicates whether the key-value pair should be included in the map.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/9QkDH3-zp
func FilterSeqToMapI[T any, K comparable, V any](collection iter.Seq[T], transform func(item T, index int) (K, V, bool)) map[K]V {
	result := make(map[K]V)

	var i int
	for item := range collection {
		if k, v, ok := transform(item, i); ok {
			result[k] = v
		}
		i++
	}

	return result
}

// Keyify returns a map with each unique element of the sequence as a key.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/0RlEI4-zq
func Keyify[T comparable](collection iter.Seq[T]) map[T]struct{} {
	result := make(map[T]struct{})

	for item := range collection {
		result[item] = struct{}{}
	}

	return result
}

// Drop drops n elements from the beginning of a sequence.
// Play: https://go.dev/play/p/1SmFJ5-zr
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
// Will allocate a slice of length n.
// Play: https://go.dev/play/p/2TnGK6-zs
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
// Play: https://go.dev/play/p/3UoHL7-zt
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
// Will allocate a slice large enough to hold the longest sequence of matching elements.
// Long input sequences of consecutive matches can cause excessive memory usage.
// Play: https://go.dev/play/p/4VpIM8-zu
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
// Will allocate a map large enough to hold all distinct indexes.
// Play: https://go.dev/play/p/5WqJN9-zv
func DropByIndex[T any, I ~func(func(T) bool)](collection I, indexes ...int) I {
	set := lo.Keyify(indexes)
	return RejectI(collection, func(_ T, index int) bool { return lo.HasKey(set, index) })
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.
// Play: https://go.dev/play/p/6XrKO0-zw
func Reject[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I {
	return RejectI(collection, func(item T, _ int) bool { return predicate(item) })
}

// RejectI is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.
// Play: https://go.dev/play/p/7YsLP1-zx
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
//
// Play: https://go.dev/play/p/8isgTsyfL-t
func RejectMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return RejectMapI(collection, func(item T, _ int) (R, bool) { return callback(item) })
}

// RejectMapI is the opposite of FilterMap, this method returns a sequence obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/9jthUzgF-u
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
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/0XrQKOk-vw
func Count[T comparable](collection iter.Seq[T], value T) int {
	return CountBy(collection, func(item T) bool { return item == value })
}

// CountBy counts the number of elements in the collection for which predicate is true.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/1SmFJ5-zr
func CountBy[T any](collection iter.Seq[T], predicate func(item T) bool) int {
	var count int

	for range Filter(collection, predicate) {
		count++
	}

	return count
}

// CountValues counts the number of each element in the collection.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/2TnGK6-zs
func CountValues[T comparable](collection iter.Seq[T]) map[T]int {
	return CountValuesBy(collection, func(item T) T { return item })
}

// CountValuesBy counts the number of each element returned from transform function.
// Is equivalent to chaining Map and CountValues.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/3UoHL7-zt
func CountValuesBy[T any, U comparable](collection iter.Seq[T], transform func(item T) U) map[U]int {
	result := make(map[U]int)

	for item := range collection {
		result[transform(item)]++
	}

	return result
}

// Subset returns a subset of a sequence from `offset` up to `length` elements.
// Will iterate at most offset+length times.
// Play: https://go.dev/play/p/4VpIM8-zu
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
// Will iterate at most end times.
// Play: https://go.dev/play/p/5WqJN9-zv
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
// Play: https://go.dev/play/p/6XrKO0-zw
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
// Play: https://go.dev/play/p/7YsLP1-zx
func ReplaceAll[T comparable, I ~func(func(T) bool)](collection I, old, nEw T) I {
	return Replace(collection, old, nEw, -1)
}

// Compact returns a sequence of all non-zero elements.
// Play: https://go.dev/play/p/8isgTsyfL-t
func Compact[T comparable, I ~func(func(T) bool)](collection I) I {
	return Filter(collection, lo.IsNotEmpty)
}

// IsSorted checks if a sequence is sorted.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/9jthUzgF-u
func IsSorted[T constraints.Ordered](collection iter.Seq[T]) bool {
	return IsSortedBy(collection, func(item T) T { return item })
}

// IsSortedBy checks if a sequence is sorted by transform.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/0XrQKOk-vw
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
// Play: https://go.dev/play/p/1SmFJ5-zr
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
// Will iterate at most the size of separator before returning.
// Play: https://go.dev/play/p/2TnGK6-zs
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
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/3UoHL7-zt
func CutSuffix[T comparable, I ~func(func(T) bool)](collection I, separator []T) (before I, found bool) {
	slice := slices.Collect(iter.Seq[T](collection))
	result, ok := lo.CutSuffix(slice, separator)
	return I(slices.Values(result)), ok
}

// Trim removes all the leading and trailing cutset from the collection.
// Will allocate a map large enough to hold all distinct cutset elements.
// Play: https://go.dev/play/p/4VpIM8-zu
func Trim[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I {
	predicate := lo.Partial(lo.HasKey, lo.Keyify(cutset))
	return DropWhile(DropLastWhile(collection, predicate), predicate)
}

// TrimFirst removes all the leading cutset from the collection.
// Will allocate a map large enough to hold all distinct cutset elements.
// Play: https://go.dev/play/p/5WqJN9-zv
func TrimFirst[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I {
	return DropWhile(collection, lo.Partial(lo.HasKey, lo.Keyify(cutset)))
}

// TrimPrefix removes all the leading prefix from the collection.
// Play: https://go.dev/play/p/6XrKO0-zw
func TrimPrefix[T comparable, I ~func(func(T) bool)](collection I, prefix []T) I {
	n := len(prefix)
	if n == 0 {
		return collection
	}

	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if i < 0 {
				if !yield(item) {
					return
				}
				continue
			}

			if item == prefix[i] {
				i = (i + 1) % n
				continue
			}

			for j := 0; j < i; j++ {
				if !yield(prefix[j]) {
					return
				}
			}
			if !yield(item) {
				return
			}
			i = -1
		}
		for j := 0; j < i; j++ {
			if !yield(prefix[j]) {
				return
			}
		}
	}
}

// TrimLast removes all the trailing cutset from the collection.
// Will allocate a map large enough to hold all distinct cutset elements.
// Play: https://go.dev/play/p/7YsLP1-zx
func TrimLast[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I {
	return DropLastWhile(collection, lo.Partial(lo.HasKey, lo.Keyify(cutset)))
}

// TrimSuffix removes all the trailing suffix from the collection.
// Play: https://go.dev/play/p/8isgTsyfL-t
func TrimSuffix[T comparable, I ~func(func(T) bool)](collection I, suffix []T) I {
	n := len(suffix)
	if n == 0 {
		return collection
	}

	return func(yield func(T) bool) {
		var i int
		for item := range collection {
			if item == suffix[i%n] {
				i++
			} else {
				for j := 0; j < i; j++ {
					if !yield(suffix[j%n]) {
						return
					}
				}
				i = 0
				if item == suffix[i] {
					i++
				} else if !yield(item) {
					return
				}
			}
		}
		if i%n != 0 {
			for j := 0; j < i; j++ {
				if !yield(suffix[j%n]) {
					return
				}
			}
		}
	}
}
