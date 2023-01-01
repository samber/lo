package lo

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
// Play: https://go.dev/play/p/Apjg3WeSi7K
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	result, _ := FilterErr(collection, func(item V, index int) (bool, error) {
		return predicate(item, index), nil
	})
	return result
}

// FilterErr is similar to Filter, with error handling for the predicate function
func FilterErr[V any](collection []V, predicate func(item V, index int) (bool, error)) ([]V, error) {
	result := []V{}

	for i, item := range collection {
		if res, err := predicate(item, i); err != nil {
			return nil, err
		} else if res {
			result = append(result, item)
		}
	}

	return result, nil
}

// Map manipulates a slice and transforms it to a slice of another type.
// Play: https://go.dev/play/p/OkPcYAhBo0D
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result, _ := MapErr(collection, func(t T, i int) (R, error) {
		return iteratee(t, i), nil
	})
	return result
}

// MapErr is similar to Map, with error handling for the iteratee function
func MapErr[T any, R any](collection []T, iteratee func(T, int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	for i, item := range collection {
		if res, err := iteratee(item, i); err != nil {
			return nil, err
		} else {
			result[i] = res
		}
	}

	return result, nil
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// Play: https://go.dev/play/p/-AuYXfy7opz
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result, _ := FilterMapErr(collection, func(item T, index int) (R, bool, error) {
		r, p := callback(item, index)
		return r, p, nil
	})
	return result
}

// FilterMapErr is similar to FilterMap, with error handling for the callback function
func FilterMapErr[T any, R any](collection []T, callback func(item T, index int) (R, bool, error)) ([]R, error) {
	result := []R{}

	for i, item := range collection {
		if r, ok, err := callback(item, i); err != nil {
			return nil, err
		} else if ok {
			result = append(result, r)
		}
	}

	return result, nil
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
// Play: https://go.dev/play/p/YSoYmQTA8-U
func FlatMap[T any, R any](collection []T, iteratee func(item T, index int) []R) []R {
	result, _ := FlatMapErr(collection, func(item T, index int) ([]R, error) {
		return iteratee(item, index), nil
	})
	return result
}

// FlatMapErr is similar to FlatMap, with error handling for the iteratee function
func FlatMapErr[T any, R any](collection []T, iteratee func(item T, index int) ([]R, error)) ([]R, error) {
	result := []R{}

	for i, item := range collection {
		if res, err := iteratee(item, i); err != nil {
			return nil, err
		} else {
			result = append(result, res...)
		}
	}

	return result, nil
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
// Play: https://go.dev/play/p/R4UHXZNaaUG
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	result, _ := ReduceErr(collection, func(agg R, item T, index int) (R, error) {
		return accumulator(agg, item, index), nil
	}, initial)
	return result
}

// ReduceErr is similar to Reduce, with error handling for the accumulator function
func ReduceErr[T any, R any](collection []T, accumulator func(agg R, item T, index int) (R, error), initial R) (R, error) {
	for i, item := range collection {
		if res, err := accumulator(initial, item, i); err != nil {
			return initial, err
		} else {
			initial = res
		}
	}

	return initial, nil
}

// ReduceRight helper is like Reduce except that it iterates over elements of collection from right to left.
// Play: https://go.dev/play/p/Fq3W70l7wXF
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	result, _ := ReduceRightErr(collection, func(agg R, item T, index int) (R, error) {
		return accumulator(agg, item, index), nil
	}, initial)
	return result
}

// ReduceRightErr is similar to ReduceRight, with error handling for the accumulator function
func ReduceRightErr[T any, R any](collection []T, accumulator func(agg R, item T, index int) (R, error), initial R) (R, error) {
	for i := len(collection) - 1; i >= 0; i-- {
		if res, err := accumulator(initial, collection[i], i); err != nil {
			return initial, err
		} else {
			initial = res
		}
	}

	return initial, nil
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// Play: https://go.dev/play/p/oofyiUPRf8t
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	_ = ForEachErr(collection, func(item T, index int) error {
		iteratee(item, index)
		return nil
	})
}

// ForEachErr is similar to ForEach, with error handling for the iteratee function
func ForEachErr[T any](collection []T, iteratee func(item T, index int) error) error {
	for i, item := range collection {
		if err := iteratee(item, i); err != nil {
			return err
		}
	}

	return nil
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// Play: https://go.dev/play/p/vgQj3Glr6lT
func Times[T any](count int, iteratee func(index int) T) []T {
	result, _ := TimesErr(count, func(index int) (T, error) {
		return iteratee(index), nil
	})
	return result
}

// TimesErr is similar to Times, with error handling for the iteratee function
func TimesErr[T any](count int, iteratee func(index int) (T, error)) ([]T, error) {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		if res, err := iteratee(i); err != nil {
			return nil, err
		} else {
			result[i] = res
		}
	}

	return result, nil
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
// Play: https://go.dev/play/p/DTzbeXZ6iEN
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
// Play: https://go.dev/play/p/g42Z3QSb53u
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	result, _ := UniqByErr(collection, func(item T) (U, error) {
		return iteratee(item), nil
	})
	return result
}

// UniqByErr is similar to UniqBy, with error handling for the iteratee function
func UniqByErr[T any, U comparable](collection []T, iteratee func(item T) (U, error)) ([]T, error) {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key, err := iteratee(item)
		if err != nil {
			return nil, err
		}

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result, nil
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// Play: https://go.dev/play/p/XnQBd_v6brd
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result, _ := GroupByErr(collection, func(item T) (U, error) {
		return iteratee(item), nil
	})
	return result
}

// GroupByErr is similar to GroupBy, with error handling for the iteratee function
func GroupByErr[T any, U comparable](collection []T, iteratee func(item T) (U, error)) (map[U][]T, error) {
	result := map[U][]T{}

	for _, item := range collection {
		key, err := iteratee(item)
		if err != nil {
			return nil, err
		}

		result[key] = append(result[key], item)
	}

	return result, nil
}

// Chunk returns an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
// Play: https://go.dev/play/p/EeKl0AuTehH
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}

	result := make([][]T, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last])
	}

	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// Play: https://go.dev/play/p/NfQ_nGjkgXW
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	result, _ := PartitionByErr(collection, func(item T) (K, error) {
		return iteratee(item), nil
	})
	return result
}

// PartitionByErr is similar to PartitionBy, with error handling for the iteratee function
func PartitionByErr[T any, K comparable](collection []T, iteratee func(item T) (K, error)) ([][]T, error) {
	result := [][]T{}
	seen := map[K]int{}

	for _, item := range collection {
		key, err := iteratee(item)
		if err != nil {
			return nil, err
		}

		resultIndex, ok := seen[key]
		if !ok {
			resultIndex = len(result)
			seen[key] = resultIndex
			result = append(result, []T{})
		}

		result[resultIndex] = append(result[resultIndex], item)
	}

	return result, nil

	// unordered:
	// groups := GroupBy[T, K](collection, iteratee)
	// return Values[K, []T](groups)
}

// Flatten returns an array a single level deep.
// Play: https://go.dev/play/p/rbp9ORaMpjw
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make([]T, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

// Interleave round-robin alternating input slices and sequentially appending value at index into result
// Play: https://go.dev/play/p/DDhlwrShbwe
func Interleave[T any](collections ...[]T) []T {
	if len(collections) == 0 {
		return []T{}
	}

	maxSize := 0
	totalSize := 0
	for _, c := range collections {
		size := len(c)
		totalSize += size
		if size > maxSize {
			maxSize = size
		}
	}

	if maxSize == 0 {
		return []T{}
	}

	result := make([]T, totalSize)

	resultIdx := 0
	for i := 0; i < maxSize; i++ {
		for j := range collections {
			if len(collections[j])-1 < i {
				continue
			}

			result[resultIdx] = collections[j][i]
			resultIdx++
		}
	}

	return result
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play: https://go.dev/play/p/Qp73bnTDnc7
func Shuffle[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/fhUMLvZ7vS6
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}

// Fill fills elements of array with `initial` value.
// Play: https://go.dev/play/p/VwR34GzqEub
func Fill[T Clonable[T]](collection []T, initial T) []T {
	result := make([]T, 0, len(collection))

	for range collection {
		result = append(result, initial.Clone())
	}

	return result
}

// Repeat builds a slice with N copies of initial value.
// Play: https://go.dev/play/p/g3uHXbmc3b6
func Repeat[T Clonable[T]](count int, initial T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}

	return result
}

// RepeatBy builds a slice with values returned by N calls of callback.
// Play: https://go.dev/play/p/ozZLCtX_hNU
func RepeatBy[T any](count int, predicate func(index int) T) []T {
	result, _ := RepeatByErr(count, func(index int) (T, error) {
		return predicate(index), nil
	})
	return result
}

// RepeatByErr is similar to RepeatBy, with error handling for the predicate function
func RepeatByErr[T any](count int, predicate func(index int) (T, error)) ([]T, error) {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		if res, err := predicate(i); err != nil {
			return nil, err
		} else {
			result = append(result, res)
		}
	}

	return result, nil
}

// KeyBy transforms a slice or an array of structs to a map based on a pivot callback.
// Play: https://go.dev/play/p/mdaClUAT-zZ
func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	result, _ := KeyByErr(collection, func(item V) (K, error) {
		return iteratee(item), nil
	})
	return result
}

// KeyByErr is similar to KeyBy, with error handling for the iteratee function
func KeyByErr[K comparable, V any](collection []V, iteratee func(item V) (K, error)) (map[K]V, error) {
	result := make(map[K]V, len(collection))

	for _, v := range collection {
		k, err := iteratee(v)
		if err != nil {
			return nil, err
		}
		result[k] = v
	}

	return result, nil
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Play: https://go.dev/play/p/WHa2CfMO3Lr
func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	result, _ := AssociateErr(collection, func(item T) (K, V, error) {
		k, v := transform(item)
		return k, v, nil
	})
	return result
}

// AssociateErr is similar to Associate, with error handling for the transform function
func AssociateErr[T any, K comparable, V any](collection []T, transform func(item T) (K, V, error)) (map[K]V, error) {
	result := make(map[K]V, len(collection))

	for _, t := range collection {
		k, v, err := transform(t)
		if err != nil {
			return nil, err
		}
		result[k] = v
	}

	return result, nil
}

// SliceToMap returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Alias of Associate().
// Play: https://go.dev/play/p/WHa2CfMO3Lr
func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	return Associate(collection, transform)
}

// SliceToMapErr is similar to SliceToMap, with error handling for the transform function
func SliceToMapErr[T any, K comparable, V any](collection []T, transform func(item T) (K, V, error)) (map[K]V, error) {
	return AssociateErr(collection, transform)
}

// Drop drops n elements from the beginning of a slice or array.
// Play: https://go.dev/play/p/JswS7vXRJP2
func Drop[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return make([]T, 0)
	}

	result := make([]T, 0, len(collection)-n)

	return append(result, collection[n:]...)
}

// DropRight drops n elements from the end of a slice or array.
// Play: https://go.dev/play/p/GG0nXkSJJa3
func DropRight[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return []T{}
	}

	result := make([]T, 0, len(collection)-n)
	return append(result, collection[:len(collection)-n]...)
}

// DropWhile drops elements from the beginning of a slice or array while the predicate returns true.
// Play: https://go.dev/play/p/7gBPYw2IK16
func DropWhile[T any](collection []T, predicate func(item T) bool) []T {
	result, _ := DropWhileErr(collection, func(item T) (bool, error) {
		return predicate(item), nil
	})
	return result
}

// DropWhileErr is similar to DropWhile, with error handling for the predicate function
func DropWhileErr[T any](collection []T, predicate func(item T) (bool, error)) ([]T, error) {
	i := 0
	for ; i < len(collection); i++ {
		res, err := predicate(collection[i])
		if err != nil {
			return nil, err
		}
		if !res {
			break
		}
	}

	result := make([]T, 0, len(collection)-i)
	return append(result, collection[i:]...), nil
}

// DropRightWhile drops elements from the end of a slice or array while the predicate returns true.
// Play: https://go.dev/play/p/3-n71oEC0Hz
func DropRightWhile[T any](collection []T, predicate func(item T) bool) []T {
	result, _ := DropRightWhileErr(collection, func(item T) (bool, error) {
		return predicate(item), nil
	})
	return result
}

// DropRightWhileErr is similar to DropRightWhile, with error handling for the predicate function
func DropRightWhileErr[T any](collection []T, predicate func(item T) (bool, error)) ([]T, error) {
	i := len(collection) - 1
	for ; i >= 0; i-- {
		res, err := predicate(collection[i])
		if err != nil {
			return nil, err
		}
		if !res {
			break
		}
	}

	result := make([]T, 0, i+1)
	return append(result, collection[:i+1]...), nil
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.
// Play: https://go.dev/play/p/YkLMODy1WEL
func Reject[V any](collection []V, predicate func(item V, index int) bool) []V {
	result, _ := RejectErr(collection, func(item V, index int) (bool, error) {
		return predicate(item, index), nil
	})
	return result
}

// RejectErr is similar to Reject, with error handling for the predicate function
func RejectErr[V any](collection []V, predicate func(item V, index int) (bool, error)) ([]V, error) {
	result := []V{}

	for i, item := range collection {
		res, err := predicate(item, i)
		if err != nil {
			return nil, err
		}
		if !res {
			result = append(result, item)
		}
	}

	return result, nil
}

// Count counts the number of elements in the collection that compare equal to value.
// Play: https://go.dev/play/p/Y3FlK54yveC
func Count[T comparable](collection []T, value T) (count int) {
	for _, item := range collection {
		if item == value {
			count++
		}
	}

	return count
}

// CountBy counts the number of elements in the collection for which predicate is true.
// Play: https://go.dev/play/p/ByQbNYQQi4X
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	result, _ := CountByErr(collection, func(item T) (bool, error) {
		return predicate(item), nil
	})
	return result
}

// CountByErr is similar to CountBy, with error handling for the predicate function
func CountByErr[T any](collection []T, predicate func(item T) (bool, error)) (count int, err error) {
	for _, item := range collection {
		if res, err := predicate(item); err != nil {
			return count, err
		} else if res {
			count++
		}
	}

	return count, nil
}

// CountValues counts the number of each element in the collection.
// Play: https://go.dev/play/p/-p-PyLT4dfy
func CountValues[T comparable](collection []T) map[T]int {
	result := make(map[T]int)

	for _, item := range collection {
		result[item]++
	}

	return result
}

// CountValuesBy counts the number of each element return from mapper function.
// Is equivalent to chaining lo.Map and lo.CountValues.
// Play: https://go.dev/play/p/2U0dG1SnOmS
func CountValuesBy[T any, U comparable](collection []T, mapper func(item T) U) map[U]int {
	result, _ := CountValuesByErr(collection, func(item T) (U, error) {
		return mapper(item), nil
	})
	return result
}

// CountValuesByErr is similar to CountValuesBy, with error handling for the mapper function
func CountValuesByErr[T any, U comparable](collection []T, mapper func(item T) (U, error)) (map[U]int, error) {
	result := make(map[U]int)

	for _, item := range collection {
		if res, err := mapper(item); err != nil {
			return nil, err
		} else {
			result[res]++
		}
	}

	return result, nil
}

// Subset returns a copy of a slice from `offset` up to `length` elements. Like `slice[start:start+length]`, but does not panic on overflow.
// Play: https://go.dev/play/p/tOQu1GhFcog
func Subset[T any](collection []T, offset int, length uint) []T {
	size := len(collection)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset > size {
		return []T{}
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	return collection[offset : offset+int(length)]
}

// Slice returns a copy of a slice from `start` up to, but not including `end`. Like `slice[start:end]`, but does not panic on overflow.
// Play: https://go.dev/play/p/8XWYhfMMA1h
func Slice[T any](collection []T, start int, end int) []T {
	size := len(collection)

	if start >= end {
		return []T{}
	}

	if start > size {
		start = size
	}
	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}
	if end < 0 {
		end = 0
	}

	return collection[start:end]
}

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/XfPzmf9gql6
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll returns a copy of the slice with all non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/a9xZFUHfYcV
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return Replace(collection, old, new, -1)
}

// Compact returns a slice of all non-zero elements.
// Play: https://go.dev/play/p/tXiy-iK6PAc
func Compact[T comparable](collection []T) []T {
	var zero T

	result := []T{}

	for _, item := range collection {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// IsSorted checks if a slice is sorted.
// Play: https://go.dev/play/p/mc3qR-t4mcx
func IsSorted[T constraints.Ordered](collection []T) bool {
	for i := 1; i < len(collection); i++ {
		if collection[i-1] > collection[i] {
			return false
		}
	}

	return true
}

// IsSortedByKey checks if a slice is sorted by iteratee.
// Play: https://go.dev/play/p/wiG6XyBBu49
func IsSortedByKey[T any, K constraints.Ordered](collection []T, iteratee func(item T) K) bool {
	result, _ := IsSortedByKeyErr(collection, func(item T) (K, error) {
		return iteratee(item), nil
	})
	return result
}

// IsSortedByKeyErr is similar to IsSortedByKey, with error handling for the iteratee function
func IsSortedByKeyErr[T any, K constraints.Ordered](collection []T, iteratee func(item T) (K, error)) (bool, error) {
	size := len(collection)

	for i := 0; i < size-1; i++ {
		resi, err := iteratee(collection[i])
		if err != nil {
			return false, err
		}
		resi1, err := iteratee(collection[i+1])
		if err != nil {
			return false, err
		}

		if resi > resi1 {
			return false, nil
		}
	}

	return true, nil
}
