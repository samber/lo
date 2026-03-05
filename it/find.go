//go:build go1.23

package it

import (
	"iter"
	"slices"
	"time"

	"github.com/samber/lo"
	"github.com/samber/lo/internal/constraints"
	"github.com/samber/lo/internal/xrand"
)

// IndexOf returns the index at which the first occurrence of a value is found in a sequence or -1
// if the value cannot be found.
// Will iterate through the entire sequence if element is not found.
// Play: https://go.dev/play/p/1OZHU2yfb-m
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
// Play: https://go.dev/play/p/QPATR3VC5wT
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
// Play: https://go.dev/play/p/Fyj6uq-G5IH
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
// Play: https://go.dev/play/p/r6bF9Rmq5S0
func HasSuffix[T comparable](collection iter.Seq[T], suffix ...T) bool {
	if len(suffix) == 0 {
		return true
	}

	n := len(suffix)
	buf := make([]T, n)
	var i int

	for buf[i%n] = range collection {
		i++
	}

	if i < n {
		return false
	}

	for j := range suffix {
		if suffix[j] != buf[(i+j)%n] {
			return false
		}
	}

	return true
}

// Find searches for an element in a sequence based on a predicate. Returns element and true if element was found.
// Will iterate through the entire sequence if predicate never returns true.
// Play: https://go.dev/play/p/4w28pF_l58a
func Find[T any](collection iter.Seq[T], predicate func(item T) bool) (T, bool) {
	for item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	return lo.Empty[T](), false
}

// FindIndexOf searches for an element in a sequence based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Will iterate through the entire sequence if predicate never returns true.
// Play: https://go.dev/play/p/ihchBAEkhXO
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
// Play: https://go.dev/play/p/ezz6hXaC4Md
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
// Play: https://go.dev/play/p/1harvaiGMfI
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
// Play: https://go.dev/play/p/O8dwXEbT56F
func FindUniques[T comparable, I ~func(func(T) bool)](collection I) I {
	return FindUniquesBy(collection, func(item T) T { return item })
}

// FindUniquesBy returns a sequence with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will iterate through the entire sequence before yielding and allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/TiwGIzeDuML
func FindUniquesBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I {
	return func(yield func(T) bool) {
		isDupl := make(map[U]bool)

		for item := range collection {
			key := transform(item)

			duplicated, seen := isDupl[key]
			if !duplicated {
				isDupl[key] = seen
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
// Play: https://go.dev/play/p/dw-VLQXKijT
func FindDuplicates[T comparable, I ~func(func(T) bool)](collection I) I {
	return FindDuplicatesBy(collection, func(item T) T { return item })
}

// FindDuplicatesBy returns a sequence with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order duplicates occur in the sequence. A transform function is
// invoked for each element in the sequence to generate the criterion by which uniqueness is computed.
// Will allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/tm1tZdC93OH
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
// Play: https://go.dev/play/p/0VihyYEaM-M
func Min[T constraints.Ordered](collection iter.Seq[T]) T {
	return MinBy(collection, func(a, b T) bool { return a < b })
}

// MinIndex search the minimum value of a collection and the index of the minimum value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/70ncPxECj6l
func MinIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	return MinIndexBy(collection, func(a, b T) bool { return a < b })
}

// MinBy search the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/J5koo8khN-g
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
// Play: https://go.dev/play/p/blldzWJpqVa
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
// Play: https://go.dev/play/p/fI6_S10H7Py
func Earliest(times iter.Seq[time.Time]) time.Time {
	return MinBy(times, func(a, b time.Time) bool { return a.Before(b) })
}

// EarliestBy search the minimum time.Time of a collection using the given transform function.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/y_Pf3Jmw-B4
func EarliestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T {
	return MinBy(collection, func(a, b T) bool { return transform(a).Before(transform(b)) })
}

// Max searches the maximum value of a collection.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/C2ZtW2bsBZ6
func Max[T constraints.Ordered](collection iter.Seq[T]) T {
	return MaxBy(collection, func(a, b T) bool { return a > b })
}

// MaxIndex searches the maximum value of a collection and the index of the maximum value.
// Returns (zero value, -1) when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/zeu2wUvhl5e
func MaxIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int) {
	return MaxIndexBy(collection, func(a, b T) bool { return a > b })
}

// MaxBy search the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/yBhXFJb5oxC
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
// Play: https://go.dev/play/p/MXyE6BTILjx
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
// Play: https://go.dev/play/p/r5Yq6ATSHoH
func Latest(times iter.Seq[time.Time]) time.Time {
	return MaxBy(times, func(a, b time.Time) bool { return a.After(b) })
}

// LatestBy search the maximum time.Time of a collection using the given transform function.
// Returns zero value when the collection is empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/o_daRzHrDUU
func LatestBy[T any](collection iter.Seq[T], transform func(item T) time.Time) T {
	return MaxBy(collection, func(a, b T) bool { return transform(a).After(transform(b)) })
}

// First returns the first element of a collection and check for availability of the first element.
// Will iterate at most once.
// Play: https://go.dev/play/p/EhNyrc8jPfY
func First[T any](collection iter.Seq[T]) (T, bool) {
	for item := range collection {
		return item, true
	}

	return lo.Empty[T](), false
}

// FirstOrEmpty returns the first element of a collection or zero value if empty.
// Will iterate at most once.
// Play: https://go.dev/play/p/NTUTgPCfevx
func FirstOrEmpty[T any](collection iter.Seq[T]) T {
	i, _ := First(collection)
	return i
}

// FirstOr returns the first element of a collection or the fallback value if empty.
// Will iterate at most once.
// Play: https://go.dev/play/p/wGFXI5NHkE2
func FirstOr[T any](collection iter.Seq[T], fallback T) T {
	if i, ok := First(collection); ok {
		return i
	}

	return fallback
}

// Last returns the last element of a collection or error if empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/eGZV-sSmn_Q
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
// Play: https://go.dev/play/p/teODFK4YqM4
func LastOrEmpty[T any](collection iter.Seq[T]) T {
	i, _ := Last(collection)
	return i
}

// LastOr returns the last element of a collection or the fallback value if empty.
// Will iterate through the entire sequence.
// Play: https://go.dev/play/p/HNubjW2Mrxs
func LastOr[T any](collection iter.Seq[T], fallback T) T {
	if i, ok := Last(collection); ok {
		return i
	}

	return fallback
}

// Nth returns the element at index `nth` of collection. An error is returned when nth is out of bounds.
// Will iterate n times through the sequence.
// Play: https://go.dev/play/p/FqgCobsKqva
func Nth[T any, N constraints.Integer](collection iter.Seq[T], nth N) (T, error) {
	value, ok := seqNth(collection, nth)

	return value, lo.Validate(ok, "nth: %d out of bounds", nth)
}

func seqNth[T any, N constraints.Integer](collection iter.Seq[T], nth N) (T, bool) {
	if nth >= 0 {
		var i N
		for item := range collection {
			if i == nth {
				return item, true
			}
			i++
		}
	}

	return lo.Empty[T](), false
}

// NthOr returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the fallback value instead of an error.
// Will iterate n times through the sequence.
// Play: https://go.dev/play/p/MNweuhpy4Ym
func NthOr[T any, N constraints.Integer](collection iter.Seq[T], nth N, fallback T) T {
	value, ok := seqNth(collection, nth)
	if !ok {
		return fallback
	}
	return value
}

// NthOrEmpty returns the element at index `nth` of collection.
// If `nth` is out of bounds, it returns the zero value (empty value) for that type.
// Will iterate n times through the sequence.
// Play: https://go.dev/play/p/pC0Zhu3EUhe
func NthOrEmpty[T any, N constraints.Integer](collection iter.Seq[T], nth N) T {
	value, _ := seqNth(collection, nth)
	return value
}

// Sample returns a random item from collection.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/YDJVX0UXYDi
func Sample[T any](collection iter.Seq[T]) T {
	return SampleBy(collection, xrand.IntN)
}

// SampleBy returns a random item from collection, using randomIntGenerator as the random index generator.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/QQooySxORib
func SampleBy[T any](collection iter.Seq[T], randomIntGenerator func(int) int) T {
	slice := slices.Collect(collection)
	return lo.SampleBy(slice, randomIntGenerator)
}

// Samples returns N random unique items from collection.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/GUTFx9LQ8pP
func Samples[T any, I ~func(func(T) bool)](collection I, count int) I {
	return SamplesBy(collection, count, xrand.IntN)
}

// SamplesBy returns N random unique items from collection, using randomIntGenerator as the random index generator.
// Will iterate through the entire sequence and allocate a slice large enough to hold all elements.
// Long input sequences can cause excessive memory usage.
// Play: https://go.dev/play/p/fX2FEtixrVG
func SamplesBy[T any, I ~func(func(T) bool)](collection I, count int, randomIntGenerator func(int) int) I {
	slice := slices.Collect(iter.Seq[T](collection))
	seq := lo.SamplesBy(slice, count, randomIntGenerator)
	return I(slices.Values(seq))
}
