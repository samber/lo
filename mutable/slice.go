package mutable

import "math/rand"

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
// The function returns the modified slice, which may be shorter than the original if some elements were removed.
// The order of elements in the original slice is preserved in the output.
// Play:
func Filter[T any](collection *[]T, predicate func(item T) bool) {
	FilterI(collection, func(item T, index int) bool {
		return predicate(item)
	})
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
// The function returns the modified slice, which may be shorter than the original if some elements were removed.
// The order of elements in the original slice is preserved in the output.
// Play:
func FilterI[T any](collection *[]T, predicate func(item T, index int) bool) {
	j := 0
	for i := range *collection {
		if predicate((*collection)[i], i) {
			(*collection)[j] = (*collection)[i]
			j++
		}
	}

	*collection = (*collection)[:j]
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
// Play:
func Uniq[T comparable](collection *[]T) {
	size := len(*collection)
	seen := make(map[T]struct{}, size)
	j := 0

	for i := 0; i < size; i++ {
		if _, ok := seen[(*collection)[i]]; ok {
			continue
		}

		seen[(*collection)[i]] = struct{}{}

		(*collection)[j] = (*collection)[i]
		j++
	}

	*collection = (*collection)[:j]
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
// Play:
func UniqBy[T any, U comparable](collection *[]T, iteratee func(item T) U) {
	size := len(*collection)
	seen := make(map[U]struct{}, size)
	j := 0

	for i := 0; i < size; i++ {
		key := iteratee((*collection)[i])
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}

		(*collection)[j] = (*collection)[i]
		j++
	}

	*collection = (*collection)[:j]
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play:
func Shuffle[T any](collection []T) {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play:
func Reverse[T any](collection []T) {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
}
