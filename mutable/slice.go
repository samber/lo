package mutable

import "github.com/samber/lo/internal/xrand"

// Filter overwrites collection's underlying array with the elements that satisfy
// predicate, in their original order, and returns a slice header whose length is the
// number kept. Use FilterI if you need the element's index.
//
// The caller's original slice variable still has its original length: only the returned
// slice has the new shorter length. Anything past that point in the original array is
// leftover from before the call (often duplicates of the last kept element). Either
// assign the result back, or re-slice with the returned length, if you want to discard
// the leftover.
// Play: https://go.dev/play/p/0jY3Z0B7O_5
func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice {
	j := 0
	for i := range collection {
		if predicate(collection[i]) {
			collection[j] = collection[i]
			j++
		}
	}
	return collection[:j]
}

// FilterI is like Filter but passes the element's index to predicate as well.
// See Filter for the in-place semantics (the caller's original slice keeps its length;
// only the returned slice has the new shorter length).
func FilterI[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice {
	j := 0
	for i := range collection {
		if predicate(collection[i], i) {
			collection[j] = collection[i]
			j++
		}
	}
	return collection[:j]
}

// Reject overwrites collection's underlying array with the elements that don't satisfy
// predicate, in their original order, and returns a slice header whose length is the
// number kept. Use RejectI if you need the element's index.
//
// The caller's original slice variable still has its original length: only the returned
// slice has the new shorter length. Anything past that point in the original array is
// leftover from before the call (often duplicates of the last kept element). Either
// assign the result back, or re-slice with the returned length, if you want to discard
// the leftover.
func Reject[T any, Slice ~[]T](collection Slice, predicate func(T) bool) Slice {
	j := 0
	for i := range collection {
		if !predicate(collection[i]) {
			collection[j] = collection[i]
			j++
		}
	}
	return collection[:j]
}

// RejectI is like Reject but passes the element's index to predicate as well.
// See Reject for the in-place semantics (the caller's original slice keeps its length;
// only the returned slice has the new shorter length).
func RejectI[T any, Slice ~[]T](collection Slice, predicate func(T, int) bool) Slice {
	j := 0
	for i := range collection {
		if !predicate(collection[i], i) {
			collection[j] = collection[i]
			j++
		}
	}
	return collection[:j]

}

// Map applies transform to each element of collection in place. Use MapI if you need
// the element's index.
// Play: https://go.dev/play/p/0jY3Z0B7O_5
func Map[T any, Slice ~[]T](collection Slice, transform func(item T) T) {
	for i := range collection {
		collection[i] = transform(collection[i])
	}
}

// MapI is like Map but passes the element's index to transform as well.
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
