package mutable

import "github.com/samber/lo/internal/rand"

// Filter is a generic function that modifies the input slice in-place to contain only the elements
// that satisfy the provided predicate function. The predicate function takes an element of the slice and its index,
// and should return true for elements that should be kept and false for elements that should be removed.
// The function returns the modified slice, which may be shorter than the original if some elements were removed.
// Note that the order of elements in the original slice is preserved in the output.
func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice {
	j := 0
	for _, item := range collection {
		if predicate(item) {
			collection[j] = item
			j++
		}
	}
	return collection[:j]
}

// FilterI is a generic function that modifies the input slice in-place to contain only the elements
// that satisfy the provided predicate function. The predicate function takes an element of the slice and its index,
// and should return true for elements that should be kept and false for elements that should be removed.
// The function returns the modified slice, which may be shorter than the original if some elements were removed.
// Note that the order of elements in the original slice is preserved in the output.
func FilterI[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice {
	j := 0
	for i, item := range collection {
		if predicate(item, i) {
			collection[j] = item
			j++
		}
	}
	return collection[:j]
}

// Map is a generic function that modifies the input slice in-place to contain the result of applying the provided
// function to each element of the slice. The function returns the modified slice, which has the same length as the original.
func Map[T any, Slice ~[]T](collection Slice, fn func(item T) T) {
	for i := range collection {
		collection[i] = fn(collection[i])
	}
}

// MapI is a generic function that modifies the input slice in-place to contain the result of applying the provided
// function to each element of the slice. The function returns the modified slice, which has the same length as the original.
func MapI[T any, Slice ~[]T](collection Slice, fn func(item T, index int) T) {
	for i := range collection {
		collection[i] = fn(collection[i], i)
	}
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play: https://go.dev/play/p/2xb3WdLjeSJ
func Shuffle[T any, Slice ~[]T](collection Slice) {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/O-M5pmCRgzV
func Reverse[T any, Slice ~[]T](collection Slice) {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
}
