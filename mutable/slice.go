package mutable

import "github.com/samber/lo/internal/xrand"

// Filter is like lo.Filter but reuses the input slice's buffer instead of
// allocating a new one. The returned slice may be shorter than the input,
// so callers must use the return value. Removed slots in the backing array
// are zeroed out so the surviving elements don't keep stale references
// alive (and so callers who accidentally inspect the original slice past
// the new length see zero values, not phantom duplicates).
// Play: https://go.dev/play/p/0jY3Z0B7O_5
func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice {
	j := 0
	for i := range collection {
		if predicate(collection[i]) {
			collection[j] = collection[i]
			j++
		}
	}
	var zero T
	for k := j; k < len(collection); k++ {
		collection[k] = zero
	}
	return collection[:j]
}

// FilterI is like Filter but the predicate also receives the element's
// index in the input slice. The same caveats apply: callers must use the
// return value, and removed slots are zeroed out.
func FilterI[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice {
	j := 0
	for i := range collection {
		if predicate(collection[i], i) {
			collection[j] = collection[i]
			j++
		}
	}
	var zero T
	for k := j; k < len(collection); k++ {
		collection[k] = zero
	}
	return collection[:j]
}

// Map applies transform to each element of collection in-place. The length
// of the slice is unchanged.
// Play: https://go.dev/play/p/0jY3Z0B7O_5
func Map[T any, Slice ~[]T](collection Slice, transform func(item T) T) {
	for i := range collection {
		collection[i] = transform(collection[i])
	}
}

// MapI is like Map but the transform also receives the element's index.
func MapI[T any, Slice ~[]T](collection Slice, transform func(item T, index int) T) {
	for i := range collection {
		collection[i] = transform(collection[i], i)
	}
}

// Shuffle returns a slice of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play: https://go.dev/play/p/2xb3WdLjeSJ
func Shuffle[T any, Slice ~[]T](collection Slice) {
	xrand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
}

// Reverse reverses a slice so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/O-M5pmCRgzV
func Reverse[T any, Slice ~[]T](collection Slice) {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i++ {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
}
