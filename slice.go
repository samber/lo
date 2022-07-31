package lo

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(V, int) bool) []V {
	result := []V{}

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection []T, callback func(T, int) (R, bool)) []R {
	result := []R{}

	for i, item := range collection {
		if r, ok := callback(item, i); ok {
			result = append(result, r)
		}
	}

	return result
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
func FlatMap[T any, R any](collection []T, iteratee func(T, int) []R) []R {
	result := []R{}

	for i, item := range collection {
		result = append(result, iteratee(item, i)...)
	}

	return result
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(R, T, int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}

	return initial
}

// ReduceRight helper is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight[T any, R any](collection []T, accumulator func(R, T, int) R, initial R) R {
	for i := len(collection) - 1; i >= 0; i-- {
		initial = accumulator(initial, collection[i], i)
	}

	return initial
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(T, int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
func Times[T any](count int, iteratee func(int) T) []T {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		result[i] = iteratee(i)
	}

	return result
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
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
func UniqBy[T any, U comparable](collection []T, iteratee func(T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}

	return result
}

// Chunk returns an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	result := make([][]T, 0, len(collection)/2+1)
	length := len(collection)

	for i := 0; i < length; i++ {
		chunk := i / size

		if i%size == 0 {
			result = append(result, make([]T, 0, size))
		}

		result[chunk] = append(result[chunk], collection[i])
	}

	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
func PartitionBy[T any, K comparable](collection []T, iteratee func(x T) K) [][]T {
	result := [][]T{}
	seen := map[K]int{}

	for _, item := range collection {
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

	// unordered:
	// groups := GroupBy[T, K](collection, iteratee)
	// return Values[K, []T](groups)
}

// Flatten returns an array a single level deep.
func Flatten[T any](collection [][]T) (result []T) {
	for _, item := range collection {
		result = append(result, item...)
	}
	return
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
func Shuffle[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
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
func Fill[T Clonable[T]](collection []T, initial T) []T {
	result := make([]T, 0, len(collection))

	for range collection {
		result = append(result, initial.Clone())
	}

	return result
}

// Repeat builds a slice with N copies of initial value.
func Repeat[T Clonable[T]](count int, initial T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}

	return result
}

// RepeatBy builds a slice with values returned by N calls of callback.
func RepeatBy[T any](count int, predicate func(int) T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, predicate(i))
	}

	return result
}

// KeyBy transforms a slice or an array of structs to a map based on a pivot callback.
func KeyBy[K comparable, V any](collection []V, iteratee func(V) K) map[K]V {
	result := make(map[K]V, len(collection))

	for _, v := range collection {
		k := iteratee(v)
		result[k] = v
	}

	return result
}

// SliceToMap2 returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap2[T any, K comparable, V any](collection []T, transform func(T) (K, V)) map[K]V {
	result := make(map[K]V)

	for _, t := range collection {
		k, v := transform(t)
		result[k] = v
	}

	return result
}

// SliceToMap3 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap3[T any, K1 comparable, K2 comparable, V any](collection []T, transform func(T) (K1, K2, V)) map[K1]map[K2]V {
	result := map[K1]map[K2]V{}

	for _, t := range collection {
		k1, k2, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]V{}
		}

		result[k1][k2] = v
	}

	return result
}

// SliceToMap4 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap4[T any, K1 comparable, K2 comparable, K3 comparable, V any](collection []T, transform func(T) (K1, K2, K3, V)) map[K1]map[K2]map[K3]V {
	result := map[K1]map[K2]map[K3]V{}

	for _, t := range collection {
		k1, k2, k3, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]map[K3]V{}
		}

		if _, ok := result[k1][k2]; !ok {
			result[k1][k2] = map[K3]V{}
		}

		result[k1][k2][k3] = v
	}

	return result
}

// SliceToMap5 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap5[T any, K1 comparable, K2 comparable, K3 comparable, K4 comparable, V any](collection []T, transform func(T) (K1, K2, K3, K4, V)) map[K1]map[K2]map[K3]map[K4]V {
	result := map[K1]map[K2]map[K3]map[K4]V{}

	for _, t := range collection {
		k1, k2, k3, k4, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]map[K3]map[K4]V{}
		}

		if _, ok := result[k1][k2]; !ok {
			result[k1][k2] = map[K3]map[K4]V{}
		}

		if _, ok := result[k1][k2][k3]; !ok {
			result[k1][k2][k3] = map[K4]V{}
		}

		result[k1][k2][k3][k4] = v
	}

	return result
}

// SliceToMap6 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap6[T any, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, V any](collection []T, transform func(T) (K1, K2, K3, K4, K5, V)) map[K1]map[K2]map[K3]map[K4]map[K5]V {
	result := map[K1]map[K2]map[K3]map[K4]map[K5]V{}

	for _, t := range collection {
		k1, k2, k3, k4, k5, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]map[K3]map[K4]map[K5]V{}
		}

		if _, ok := result[k1][k2]; !ok {
			result[k1][k2] = map[K3]map[K4]map[K5]V{}
		}

		if _, ok := result[k1][k2][k3]; !ok {
			result[k1][k2][k3] = map[K4]map[K5]V{}
		}

		if _, ok := result[k1][k2][k3][k4]; !ok {
			result[k1][k2][k3][k4] = map[K5]V{}
		}

		result[k1][k2][k3][k4][k5] = v
	}

	return result
}

// SliceToMap7 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap7[T any, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, V any](collection []T, transform func(T) (K1, K2, K3, K4, K5, K6, V)) map[K1]map[K2]map[K3]map[K4]map[K5]map[K6]V {
	result := map[K1]map[K2]map[K3]map[K4]map[K5]map[K6]V{}

	for _, t := range collection {
		k1, k2, k3, k4, k5, k6, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]map[K3]map[K4]map[K5]map[K6]V{}
		}

		if _, ok := result[k1][k2]; !ok {
			result[k1][k2] = map[K3]map[K4]map[K5]map[K6]V{}
		}

		if _, ok := result[k1][k2][k3]; !ok {
			result[k1][k2][k3] = map[K4]map[K5]map[K6]V{}
		}

		if _, ok := result[k1][k2][k3][k4]; !ok {
			result[k1][k2][k3][k4] = map[K5]map[K6]V{}
		}

		if _, ok := result[k1][k2][k3][k4][k5]; !ok {
			result[k1][k2][k3][k4][k5] = map[K6]V{}
		}

		result[k1][k2][k3][k4][k5][k6] = v
	}

	return result
}

// SliceToMap8 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap8[T any, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, V any](collection []T, transform func(T) (K1, K2, K3, K4, K5, K6, K7, V)) map[K1]map[K2]map[K3]map[K4]map[K5]map[K6]map[K7]V {
	result := map[K1]map[K2]map[K3]map[K4]map[K5]map[K6]map[K7]V{}

	for _, t := range collection {
		k1, k2, k3, k4, k5, k6, k7, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]map[K3]map[K4]map[K5]map[K6]map[K7]V{}
		}

		if _, ok := result[k1][k2]; !ok {
			result[k1][k2] = map[K3]map[K4]map[K5]map[K6]map[K7]V{}
		}

		if _, ok := result[k1][k2][k3]; !ok {
			result[k1][k2][k3] = map[K4]map[K5]map[K6]map[K7]V{}
		}

		if _, ok := result[k1][k2][k3][k4]; !ok {
			result[k1][k2][k3][k4] = map[K5]map[K6]map[K7]V{}
		}

		if _, ok := result[k1][k2][k3][k4][k5]; !ok {
			result[k1][k2][k3][k4][k5] = map[K6]map[K7]V{}
		}

		if _, ok := result[k1][k2][k3][k4][k5][k6]; !ok {
			result[k1][k2][k3][k4][k5][k6] = map[K7]V{}
		}

		result[k1][k2][k3][k4][k5][k6][k7] = v
	}

	return result
}

// SliceToMap9 returns a map containing key-value tuples provided by transform function applied to elements of the given slice.
// If any of two tuples would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
func SliceToMap9[T any, K1 comparable, K2 comparable, K3 comparable, K4 comparable, K5 comparable, K6 comparable, K7 comparable, K8 comparable, V any](collection []T, transform func(T) (K1, K2, K3, K4, K5, K6, K7, K8, V)) map[K1]map[K2]map[K3]map[K4]map[K5]map[K6]map[K7]map[K8]V {
	result := map[K1]map[K2]map[K3]map[K4]map[K5]map[K6]map[K7]map[K8]V{}

	for _, t := range collection {
		k1, k2, k3, k4, k5, k6, k7, k8, v := transform(t)

		if _, ok := result[k1]; !ok {
			result[k1] = map[K2]map[K3]map[K4]map[K5]map[K6]map[K7]map[K8]V{}
		}

		if _, ok := result[k1][k2]; !ok {
			result[k1][k2] = map[K3]map[K4]map[K5]map[K6]map[K7]map[K8]V{}
		}

		if _, ok := result[k1][k2][k3]; !ok {
			result[k1][k2][k3] = map[K4]map[K5]map[K6]map[K7]map[K8]V{}
		}

		if _, ok := result[k1][k2][k3][k4]; !ok {
			result[k1][k2][k3][k4] = map[K5]map[K6]map[K7]map[K8]V{}
		}

		if _, ok := result[k1][k2][k3][k4][k5]; !ok {
			result[k1][k2][k3][k4][k5] = map[K6]map[K7]map[K8]V{}
		}

		if _, ok := result[k1][k2][k3][k4][k5][k6]; !ok {
			result[k1][k2][k3][k4][k5][k6] = map[K7]map[K8]V{}
		}

		if _, ok := result[k1][k2][k3][k4][k5][k6][k7]; !ok {
			result[k1][k2][k3][k4][k5][k6][k7] = map[K8]V{}
		}

		result[k1][k2][k3][k4][k5][k6][k7][k8] = v
	}

	return result
}

// SliceToMap returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Alias of SliceToMap2().
func SliceToMap[T any, K comparable, V any](collection []T, transform func(T) (K, V)) map[K]V {
	return SliceToMap2(collection, transform)
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Alias of SliceToMap().
func Associate[T any, K comparable, V any](collection []T, transform func(T) (K, V)) map[K]V {
	return SliceToMap(collection, transform)
}

// Drop drops n elements from the beginning of a slice or array.
func Drop[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return make([]T, 0)
	}

	result := make([]T, len(collection)-n)
	for i := n; i < len(collection); i++ {
		result[i-n] = collection[i]
	}

	return result
}

// DropWhile drops elements from the beginning of a slice or array while the predicate returns true.
func DropWhile[T any](collection []T, predicate func(T) bool) []T {
	i := 0
	for ; i < len(collection); i++ {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, len(collection)-i)

	for j := 0; i < len(collection); i, j = i+1, j+1 {
		result[j] = collection[i]
	}

	return result
}

// DropRight drops n elements from the end of a slice or array.
func DropRight[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return make([]T, 0)
	}

	result := make([]T, len(collection)-n)
	for i := len(collection) - 1 - n; i >= 0; i-- {
		result[i] = collection[i]
	}

	return result
}

// DropRightWhile drops elements from the end of a slice or array while the predicate returns true.
func DropRightWhile[T any](collection []T, predicate func(T) bool) []T {
	i := len(collection) - 1
	for ; i >= 0; i-- {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, i+1)

	for ; i >= 0; i-- {
		result[i] = collection[i]
	}

	return result
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.
func Reject[V any](collection []V, predicate func(V, int) bool) []V {
	result := []V{}

	for i, item := range collection {
		if !predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Count counts the number of elements in the collection that compare equal to value.
func Count[T comparable](collection []T, value T) (count int) {
	for _, item := range collection {
		if item == value {
			count++
		}
	}

	return count
}

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection []T, predicate func(T) bool) (count int) {
	for _, item := range collection {
		if predicate(item) {
			count++
		}
	}

	return count
}

// Subset returns a copy of a slice from `offset` up to `length` elements. Like `slice[start:start+length]`, but does not panic on overflow.
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
func Slice[T comparable](collection []T, start int, end int) []T {
	size := len(collection)

	if start >= end {
		return []T{}
	}

	if start > size {
		start = size
	}

	if end > size {
		end = size
	}

	return collection[start:end]
}

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	size := len(collection)
	result := make([]T, 0, size)

	for _, item := range collection {
		if item == old && n != 0 {
			result = append(result, new)
			n--
		} else {
			result = append(result, item)
		}
	}

	return result
}

// ReplaceAll returns a copy of the slice with all non-overlapping instances of old replaced by new.
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return Replace(collection, old, new, -1)
}

// Compact returns a slice of all non-zero elements.
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
func IsSorted[T constraints.Ordered](collection []T) bool {
	for i := 1; i < len(collection); i++ {
		if collection[i-1] > collection[i] {
			return false
		}
	}

	return true
}

// IsSortedByKey checks if a slice is sorted by iteratee.
func IsSortedByKey[T any, K constraints.Ordered](collection []T, iteratee func(T) K) bool {
	size := len(collection)

	for i := 0; i < size-1; i++ {
		if iteratee(collection[i]) > iteratee(collection[i+1]) {
			return false
		}
	}

	return true
}
