//go:build go1.23

package it

import (
	"fmt"
	"iter"
	"slices"
	"time"

	"github.com/samber/lo"
	"github.com/samber/lo/internal/constraints"
	"github.com/samber/lo/internal/rand"
)

// IndexOf returns the index at which the first occurrence of a value is found in a sequence or -1
// if the value cannot be found.
func IndexOf[T comparable](collection iter.Seq[T], element T) int {
	var i int
	for item := range collection {
		if item == element {
			return i
		}
		i++
	}

	return -1
}

// LastIndexOf returns the index at which the last occurrence of a value is found in a sequence or -1
// if the value cannot be found.
func LastIndexOf[T comparable](collection iter.Seq[T], element T) int {
	index := -1
	var i int
	for item := range collection {
		if item == element {
			index = i
		}
		i++
	}

	return index
}

// Find searches for an element in a sequence based on a predicate. Returns element and true if element was found.
func Find[T any](collection iter.Seq[T], predicate func(item T) bool) (T, bool) {
	for item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

// FindIndexOf searches for an element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
func FindIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool) {
	var i int
	for item := range collection {
		if predicate(item) {
			return item, i, true
		}
		i++
	}

	var result T
	return result, -1, false
}

// FindLastIndexOf searches for the last element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
func FindLastIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool) {
	var result T
	index := -1
	var ok bool

	var i int
	for item := range collection {
		if predicate(item) {
			result = item
			index = i
			ok = true
		}
		i++
	}

	return result, index, ok
}

// FindOrElse searches for an element in a sequence based on a predicate. Returns the element if found or a given fallback value otherwise.
func FindOrElse[T any](collection iter.Seq[T], fallback T, predicate func(item T) bool) T {
	for item := range collection {
		if predicate(item) {
			return item
		}
	}

	return fallback
}

// FindUniques returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the collection.
func FindUniques[T comparable, I ~func(func(T) bool)](collection I) I {
	return func(yield func(T) bool) {
		isDupl := make(map[T]bool)

		for item := range collection {
			duplicated, ok := isDupl[item]
			if !ok {
				isDupl[item] = false
			} else if !duplicated {
				isDupl[item] = true
			}
		}

		for item := range collection {
			if duplicated := isDupl[item]; !duplicated && !yield(item) {
				return
			}
		}
	}
}

// FindUniquesBy returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the sequence. It accepts `iteratee` which is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
func FindUniquesBy[T any, U comparable, I ~func(func(T) bool)](collection I, iteratee func(item T) U) I {
	return func(yield func(T) bool) {
		isDupl := make(map[U]bool)

		for item := range collection {
			key := iteratee(item)

			duplicated, ok := isDupl[key]
			if !ok {
				isDupl[key] = false
			} else if !duplicated {
				isDupl[key] = true
			}
		}

		for item := range collection {
			key := iteratee(item)

			if duplicated := isDupl[key]; !duplicated && !yield(item) {
				return
			}
		}
	}
}

// FindDuplicates returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the collection.
func FindDuplicates[T comparable, I ~func(func(T) bool)](collection I) I {
	return func(yield func(T) bool) {
		isDupl := make(map[T]bool)

		for item := range collection {
			duplicated, ok := isDupl[item]
			if !ok {
				isDupl[item] = false
			} else if !duplicated {
				if !yield(item) {
					return
				}
				isDupl[item] = true
			}
		}
	}
}

// FindDuplicatesBy returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the sequence. It accepts `iteratee` which is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
func FindDuplicatesBy[T any, U comparable, I ~func(func(T) bool)](collection I, iteratee func(item T) U) I {
	return func(yield func(T) bool) {
		isDupl := make(map[U]bool)
		first := make(map[U]T)

		for item := range collection {
			key := iteratee(item)

			duplicated, ok := isDupl[key]
			if !ok {
				isDupl[key] = false
				first[key] = item
			} else if !duplicated {
				if !yield(first[key]) {
					return
				}
				isDupl[key] = true
			}
		}
	}
}

// Min search the minimum value of a collection.
// Returns zero value when the collection is empty.
func Min[T constraints.Ordered](collection iter.Seq[T]) T {
	first := true
	var mIn T

	for item := range collection {
		if first {
			mIn = item
			first = false
		} else if item < mIn {
			mIn = item
		}
	}

	return mIn
}

// MinIndex search the minimum value of a collection and the index of the minimum value.
// Returns (zero value, -1) when the collection is empty.
func MinIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	var mIn T
	index := -1

	var i int
	for item := range collection {
		if i == 0 || item < mIn {
			mIn = item
			index = i
		}
		i++
	}

	return mIn, index
}

// MinBy search the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value when the collection is empty.
func MinBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) T {
	first := true
	var mIn T

	for item := range collection {
		if first {
			mIn = item
			first = false
		} else if comparison(item, mIn) {
			mIn = item
		}
	}

	return mIn
}

// MinIndexBy search the minimum value of a collection using the given comparison function and the index of the minimum value.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
func MinIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int) {
	var mIn T
	index := -1

	var i int
	for item := range collection {
		if i == 0 || comparison(item, mIn) {
			mIn = item
			index = i
		}
		i++
	}

	return mIn, index
}

// Earliest search the minimum time.Time of a collection.
// Returns zero value when the collection is empty.
func Earliest(times iter.Seq[time.Time]) time.Time {
	return MinBy(times, func(a, b time.Time) bool { return a.Before(b) })
}

// EarliestBy search the minimum time.Time of a collection using the given iteratee function.
// Returns zero value when the collection is empty.
func EarliestBy[T any](collection iter.Seq[T], iteratee func(item T) time.Time) T {
	return MinBy(collection, func(a, b T) bool { return iteratee(a).Before(iteratee(b)) })
}

// Max searches the maximum value of a collection.
// Returns zero value when the collection is empty.
func Max[T constraints.Ordered](collection iter.Seq[T]) T {
	first := true
	var mAx T

	for item := range collection {
		if first {
			mAx = item
			first = false
		} else if item > mAx {
			mAx = item
		}
	}

	return mAx
}

// MaxIndex searches the maximum value of a collection and the index of the maximum value.
// Returns (zero value, -1) when the collection is empty.
func MaxIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	var mAx T
	index := -1

	var i int
	for item := range collection {
		if i == 0 || item > mAx {
			mAx = item
			index = i
		}
		i++
	}

	return mAx, index
}

// MaxBy search the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value when the collection is empty.
func MaxBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) T {
	first := true
	var mAx T

	for item := range collection {
		if first {
			mAx = item
			first = false
		} else if comparison(item, mAx) {
			mAx = item
		}
	}

	return mAx
}

// MaxIndexBy search the maximum value of a collection using the given comparison function and the index of the maximum value.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
func MaxIndexBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) (T, int) {
	var mAx T
	index := -1

	var i int
	for item := range collection {
		if i == 0 || comparison(item, mAx) {
			mAx = item
			index = i
		}
		i++
	}

	return mAx, index
}

// Latest search the maximum time.Time of a collection.
// Returns zero value when the collection is empty.
func Latest(times iter.Seq[time.Time]) time.Time {
	return MaxBy(times, func(a, b time.Time) bool { return a.After(b) })
}

// LatestBy search the maximum time.Time of a collection using the given iteratee function.
// Returns zero value when the collection is empty.
func LatestBy[T any](collection iter.Seq[T], iteratee func(item T) time.Time) T {
	return MaxBy(collection, func(a, b T) bool { return iteratee(a).After(iteratee(b)) })
}

// First returns the first element of a collection and check for availability of the first element.
func First[T any](collection iter.Seq[T]) (T, bool) {
	for item := range collection {
		return item, true
	}

	var t T
	return t, false
}

// FirstOrEmpty returns the first element of a collection or zero value if empty.
func FirstOrEmpty[T any](collection iter.Seq[T]) T {
	i, _ := First(collection)
	return i
}

// FirstOr returns the first element of a collection or the fallback value if empty.
func FirstOr[T any](collection iter.Seq[T], fallback T) T {
	i, ok := First(collection)
	if !ok {
		return fallback
	}

	return i
}

// Last returns the last element of a collection or error if empty.
func Last[T any](collection iter.Seq[T]) (T, bool) {
	var t T
	var ok bool
	for item := range collection {
		t = item
		ok = true
	}

	return t, ok
}

// LastOrEmpty returns the last element of a collection or zero value if empty.
func LastOrEmpty[T any](collection iter.Seq[T]) T {
	i, _ := Last(collection)
	return i
}

// LastOr returns the last element of a collection or the fallback value if empty.
func LastOr[T any](collection iter.Seq[T], fallback T) T {
	i, ok := Last(collection)
	if !ok {
		return fallback
	}

	return i
}

// Nth returns the element at index `nth` of collection. An error is returned when nth is out of bounds.
func Nth[T any, N constraints.Integer](collection iter.Seq[T], nth N) (T, error) {
	if nth >= 0 {
		var i N
		for item := range collection {
			if i == nth {
				return item, nil
			}
			i++
		}
	}

	var t T
	return t, fmt.Errorf("nth: %d out of bounds", nth)
}

// NthOr returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the fallback value instead of an error.
func NthOr[T any, N constraints.Integer](collection iter.Seq[T], nth N, fallback T) T {
	value, err := Nth(collection, nth)
	if err != nil {
		return fallback
	}
	return value
}

// NthOrEmpty returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the zero value (empty value) for that type.
func NthOrEmpty[T any, N constraints.Integer](collection iter.Seq[T], nth N) T {
	value, err := Nth(collection, nth)
	if err != nil {
		var zeroValue T
		return zeroValue
	}
	return value
}

// Sample returns a random item from collection.
func Sample[T any](collection iter.Seq[T]) T {
	return SampleBy(collection, rand.IntN)
}

// SampleBy returns a random item from collection, using randomIntGenerator as the random index generator.
func SampleBy[T any](collection iter.Seq[T], randomIntGenerator func(int) int) T {
	slice := slices.Collect(collection)
	return lo.SampleBy(slice, randomIntGenerator)
}

// Samples returns N random unique items from collection.
func Samples[T any, I ~func(func(T) bool)](collection I, count int) I {
	return SamplesBy(collection, count, rand.IntN)
}

// SamplesBy returns N random unique items from collection, using randomIntGenerator as the random index generator.
func SamplesBy[T any, I ~func(func(T) bool)](collection I, count int, randomIntGenerator func(int) int) I {
	slice := slices.Collect(iter.Seq[T](collection))
	seq := lo.SamplesBy(slice, count, randomIntGenerator)
	return I(slices.Values(seq))
}
