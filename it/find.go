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
// Will iterate through the entire sequence if element is not found.
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
// Will iterate through the entire sequence.
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

// HasPrefix returns true if the collection has the prefix.
// Will iterate at most the size of prefix.
func HasPrefix[T comparable](collection iter.Seq[T], prefix ...T) bool {
	if len(prefix) == 0 {
		return true
	}

	var i int

	for item := range collection {
		if item != prefix[i] {
			return false
		}
		i++
		if i == len(prefix) {
			return true
		}
	}

	return false
}

// HasSuffix returns true if the collection has the suffix.
// Will iterate through the entire sequence and allocate a slice the size of suffix.
func HasSuffix[T comparable](collection iter.Seq[T], suffix ...T) bool {
	if len(suffix) == 0 {
		return true
	}

	n := len(suffix)
	buf := make([]T, 0, n)
	var i int

	for item := range collection {
		if len(buf) < n {
			buf = append(buf, item)
		} else {
			buf[i] = item
		}
		i = (i + 1) % n
	}

	if len(buf) < n {
		return false
	}

	i += n
	for j := range suffix {
		if suffix[j] != buf[(i+j)%n] {
			return false
		}
	}

	return true
}

// Find searches for an element in a sequence based on a predicate. Returns element and true if element was found.
// Will iterate through the entire sequence if predicate never returns true.
func Find[T any](collection iter.Seq[T], predicate func(item T) bool) (T, bool) {
	return First(Filter(collection, predicate))
}

// FindIndexOf searches for an element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Will iterate through the entire sequence if predicate never returns true.
func FindIndexOf[T any](collection iter.Seq[T], predicate func(item T) bool) (T, int, bool) {
	var i int
	for item := range collection {
		if predicate(item) {
			return item, i, true
		}
		i++
	}

	return lo.Empty[T](), -1, false
}

// FindLastIndexOf searches for the last element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Will iterate through the entire sequence.
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
// Will iterate through the entire sequence if predicate never returns true.
func FindOrElse[T any](collection iter.Seq[T], fallback T, predicate func(item T) bool) T {
	if result, ok := Find(collection, predicate); ok {
		return result
	}

	return fallback
}

// FindUniques returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the collection.
// Will iterate through the entire sequence before yielding and allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func FindUniques[T comparable, I ~func(func(T) bool)](collection I) I {
	return FindUniquesBy(collection, func(item T) T { return item })
}

// FindUniquesBy returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will iterate through the entire sequence before yielding and allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func FindUniquesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	return func(yield func(T) bool) {
		isDupl := make(map[U]bool)

		for item := range collection {
			key := transform(item)

			if duplicated, ok := isDupl[key]; !ok {
				isDupl[key] = false
			} else if !duplicated {
				isDupl[key] = true
			}
		}

		for item := range collection {
			key := transform(item)

			if duplicated := isDupl[key]; !duplicated && !yield(item) {
				return
			}
		}
	}
}

// FindDuplicates returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the collection.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func FindDuplicates[T comparable, I ~func(func(T) bool)](collection I) I {
	return FindDuplicatesBy(collection, func(item T) T { return item })
}

// FindDuplicatesBy returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func FindDuplicatesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	return func(yield func(T) bool) {
		isDupl := make(map[U]lo.Tuple2[T, bool])

		for item := range collection {
			key := transform(item)

			if duplicated, ok := isDupl[key]; !ok {
				isDupl[key] = lo.Tuple2[T, bool]{A: item}
			} else if !duplicated.B {
				if !yield(duplicated.A) {
					return
				}
				isDupl[key] = lo.Tuple2[T, bool]{A: item, B: true}
			}
		}
	}
}

// Min search the minimum value of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
func Min[T constraints.Ordered](collection iter.Seq[T]) T {
	return MinBy(collection, func(a, b T) bool { return a < b })
}

// MinIndex search the minimum value of a collection and the index of the minimum value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
func MinIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	return MinIndexBy(collection, func(a, b T) bool { return a < b })
}

// MinBy search the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
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
// Will iterate through the entire sequence.
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
// Will iterate through the entire sequence.
func Earliest(times iter.Seq[time.Time]) time.Time {
	return MinBy(times, func(a, b time.Time) bool { return a.Before(b) })
}

// EarliestBy search the minimum time.Time of a collection using the given transform function.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
func EarliestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T {
	return MinBy(collection, func(a, b T) bool { return transform(a).Before(transform(b)) })
}

// Max searches the maximum value of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
func Max[T constraints.Ordered](collection iter.Seq[T]) T {
	return MaxBy(collection, func(a, b T) bool { return a > b })
}

// MaxIndex searches the maximum value of a collection and the index of the maximum value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
func MaxIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	return MaxIndexBy(collection, func(a, b T) bool { return a > b })
}

// MaxBy search the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
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
// Will iterate through the entire sequence.
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
// Will iterate through the entire sequence.
func Latest(times iter.Seq[time.Time]) time.Time {
	return MaxBy(times, func(a, b time.Time) bool { return a.After(b) })
}

// LatestBy search the maximum time.Time of a collection using the given transform function.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
func LatestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T {
	return MaxBy(collection, func(a, b T) bool { return transform(a).After(transform(b)) })
}

// First returns the first element of a collection and check for availability of the first element.
// Will iterate at most once.
func First[T any](collection iter.Seq[T]) (T, bool) {
	for item := range collection {
		return item, true
	}

	return lo.Empty[T](), false
}

// FirstOrEmpty returns the first element of a collection or zero value if empty.
// Will iterate at most once.
func FirstOrEmpty[T any](collection iter.Seq[T]) T {
	i, _ := First(collection)
	return i
}

// FirstOr returns the first element of a collection or the fallback value if empty.
// Will iterate at most once.
func FirstOr[T any](collection iter.Seq[T], fallback T) T {
	if i, ok := First(collection); ok {
		return i
	}

	return fallback
}

// Last returns the last element of a collection or error if empty.
// Will iterate through the entire sequence.
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
// Will iterate through the entire sequence.
func LastOrEmpty[T any](collection iter.Seq[T]) T {
	i, _ := Last(collection)
	return i
}

// LastOr returns the last element of a collection or the fallback value if empty.
// Will iterate through the entire sequence.
func LastOr[T any](collection iter.Seq[T], fallback T) T {
	if i, ok := Last(collection); ok {
		return i
	}

	return fallback
}

// Nth returns the element at index `nth` of collection. An error is returned when nth is out of bounds.
// Will iterate n times through the sequence.
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

	return lo.Empty[T](), fmt.Errorf("nth: %d out of bounds", nth)
}

// NthOr returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the fallback value instead of an error.
// Will iterate n times through the sequence.
func NthOr[T any, N constraints.Integer](collection iter.Seq[T], nth N, fallback T) T {
	value, err := Nth(collection, nth)
	if err != nil {
		return fallback
	}
	return value
}

// NthOrEmpty returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the zero value (empty value) for that type.
// Will iterate n times through the sequence.
func NthOrEmpty[T any, N constraints.Integer](collection iter.Seq[T], nth N) T {
	value, _ := Nth(collection, nth)
	return value
}

// Sample returns a random item from collection.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
func Sample[T any](collection iter.Seq[T]) T {
	return SampleBy(collection, rand.IntN)
}

// SampleBy returns a random item from collection, using randomIntGenerator as the random index generator.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
func SampleBy[T any](collection iter.Seq[T], randomIntGenerator func(int) int) T {
	slice := slices.Collect(collection)
	return lo.SampleBy(slice, randomIntGenerator)
}

// Samples returns N random unique items from collection.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
func Samples[T any, I ~func(func(T) bool)](collection I, count int) I {
	return SamplesBy(collection, count, rand.IntN)
}

// SamplesBy returns N random unique items from collection, using randomIntGenerator as the random index generator.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
func SamplesBy[T any, I ~func(func(T) bool)](collection I, count int, randomIntGenerator func(int) int) I {
	slice := slices.Collect(iter.Seq[T](collection))
	seq := lo.SamplesBy(slice, count, randomIntGenerator)
	return I(slices.Values(seq))
}
