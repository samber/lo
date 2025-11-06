package lo

import (
	"fmt"
	"time"

	"github.com/samber/lo/internal/constraints"
	"github.com/samber/lo/internal/xrand"
)

// IndexOf returns the index at which the first occurrence of a value is found in a slice or -1
// if the value cannot be found.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func IndexOf[T comparable](collection []T, element T) int {
	for i := range collection {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index at which the last occurrence of a value is found in a slice or -1
// if the value cannot be found.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func LastIndexOf[T comparable](collection []T, element T) int {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// HasPrefix returns true if the collection has the prefix.
// Play: https://go.dev/play/p/SrljzVDpMQM
func HasPrefix[T comparable](collection, prefix []T) bool {
	if len(collection) < len(prefix) {
		return false
	}

	for i := range prefix {
		if collection[i] != prefix[i] {
			return false
		}
	}

	return true
}

// HasSuffix returns true if the collection has the suffix.
// Play: https://go.dev/play/p/bJeLetQNAON
func HasSuffix[T comparable](collection, suffix []T) bool {
	if len(collection) < len(suffix) {
		return false
	}

	for i := range suffix {
		if collection[len(collection)-len(suffix)+i] != suffix[i] {
			return false
		}
	}

	return true
}

// Find searches for an element in a slice based on a predicate. Returns element and true if element was found.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	for i := range collection {
		if predicate(collection[i]) {
			return collection[i], true
		}
	}

	return Empty[T](), false
}

// FindIndexOf searches for an element in a slice based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Play: https://go.dev/play/p/XWSEM4Ic_t0
func FindIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	for i := range collection {
		if predicate(collection[i]) {
			return collection[i], i, true
		}
	}

	return Empty[T](), -1, false
}

// FindLastIndexOf searches for the last element in a slice based on a predicate and returns the index and true.
// Returns -1 and false if the element is not found.
// Play: https://go.dev/play/p/dPiMRtJ6cUx
func FindLastIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], i, true
		}
	}

	return Empty[T](), -1, false
}

// FindOrElse searches for an element in a slice based on a predicate. Returns the element if found or a given fallback value otherwise.
// Play: https://go.dev/play/p/Eo7W0lvKTky
func FindOrElse[T any](collection []T, fallback T, predicate func(item T) bool) T {
	result, ok := Find(collection, predicate)
	return Ternary(ok, result, fallback)
}

// FindKey returns the key of the first value matching.
// Play: https://go.dev/play/p/Bg0w1VDPYXx
func FindKey[K, V comparable](object map[K]V, value V) (K, bool) {
	return FindKeyBy(object, func(k K, v V) bool { return v == value })
}

// FindKeyBy returns the key of the first element predicate returns true for.
// Play: https://go.dev/play/p/9IbiPElcyo8
func FindKeyBy[K comparable, V any](object map[K]V, predicate func(key K, value V) bool) (K, bool) {
	for k, v := range object {
		if predicate(k, v) {
			return k, true
		}
	}

	return Empty[K](), false
}

// FindUniques returns a slice with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the collection.
func FindUniques[T comparable, Slice ~[]T](collection Slice) Slice {
	return FindUniquesBy(collection, func(item T) T { return item })
}

// FindUniquesBy returns a slice with all the elements that appear in the collection only once.
// The order of result values is determined by the order they occur in the slice. It accepts `iteratee` which is
// invoked for each element in the slice to generate the criterion by which uniqueness is computed.
func FindUniquesBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice {
	isDupl := make(map[U]bool, len(collection))

	for i := range collection {
		key := iteratee(collection[i])

		duplicated, ok := isDupl[key]
		if !ok {
			isDupl[key] = false
		} else if !duplicated {
			isDupl[key] = true
		}
	}

	result := make(Slice, 0, len(collection)-len(isDupl))

	for i := range collection {
		key := iteratee(collection[i])

		if duplicated := isDupl[key]; !duplicated {
			result = append(result, collection[i])
		}
	}

	return result
}

// FindDuplicates returns a slice with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order they occur in the collection.
func FindDuplicates[T comparable, Slice ~[]T](collection Slice) Slice {
	return FindDuplicatesBy(collection, func(item T) T { return item })
}

// FindDuplicatesBy returns a slice with the first occurrence of each duplicated element in the collection.
// The order of result values is determined by the order they occur in the slice. It accepts `iteratee` which is
// invoked for each element in the slice to generate the criterion by which uniqueness is computed.
func FindDuplicatesBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice {
	isDupl := make(map[U]bool, len(collection))

	for i := range collection {
		key := iteratee(collection[i])

		duplicated, ok := isDupl[key]
		if !ok {
			isDupl[key] = false
		} else if !duplicated {
			isDupl[key] = true
		}
	}

	result := make(Slice, 0, len(collection)-len(isDupl))

	for i := range collection {
		key := iteratee(collection[i])

		if duplicated := isDupl[key]; duplicated {
			result = append(result, collection[i])
			isDupl[key] = false
		}
	}

	return result
}

// Min search the minimum value of a collection.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/r6e-Z8JozS8
func Min[T constraints.Ordered](collection []T) T {
	return MinBy(collection, func(a, b T) bool { return a < b })
}

// MinIndex search the minimum value of a collection and the index of the minimum value.
// Returns (zero value, -1) when the collection is empty.
func MinIndex[T constraints.Ordered](collection []T) (T, int) {
	return MinIndexBy(collection, func(a, b T) bool { return a < b })
}

// MinBy search the minimum value of a collection using the given comparison function.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns zero value when the collection is empty.
func MinBy[T any](collection []T, less func(a, b T) bool) T {
	var mIn T

	if len(collection) == 0 {
		return mIn
	}

	mIn = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if less(item, mIn) {
			mIn = item
		}
	}

	return mIn
}

// MinIndexBy search the minimum value of a collection using the given comparison function and the index of the minimum value.
// If several values of the collection are equal to the smallest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
func MinIndexBy[T any](collection []T, less func(a, b T) bool) (T, int) {
	var (
		mIn   T
		index int
	)

	if len(collection) == 0 {
		return mIn, -1
	}

	mIn = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if less(item, mIn) {
			mIn = item
			index = i
		}
	}

	return mIn, index
}

// Earliest search the minimum time.Time of a collection.
// Returns zero value when the collection is empty.
func Earliest(times ...time.Time) time.Time {
	return MinBy(times, func(a, b time.Time) bool { return a.Before(b) })
}

// EarliestBy search the minimum time.Time of a collection using the given iteratee function.
// Returns zero value when the collection is empty.
func EarliestBy[T any](collection []T, iteratee func(item T) time.Time) T {
	return MinBy(collection, func(a, b T) bool { return iteratee(a).Before(iteratee(b)) })
}

// Max searches the maximum value of a collection.
// Returns zero value when the collection is empty.
// Play: https://go.dev/play/p/r6e-Z8JozS8
func Max[T constraints.Ordered](collection []T) T {
	return MaxBy(collection, func(a, b T) bool { return a > b })
}

// MaxIndex searches the maximum value of a collection and the index of the maximum value.
// Returns (zero value, -1) when the collection is empty.
func MaxIndex[T constraints.Ordered](collection []T) (T, int) {
	return MaxIndexBy(collection, func(a, b T) bool { return a > b })
}

// MaxBy search the maximum value of a collection using the given comparison function.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns zero value when the collection is empty.
//
// Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.
// See https://github.com/samber/lo/issues/129
//
// Play: https://go.dev/play/p/JW1qu-ECwF7
func MaxBy[T any](collection []T, greater func(a, b T) bool) T {
	var mAx T

	if len(collection) == 0 {
		return mAx
	}

	mAx = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if greater(item, mAx) {
			mAx = item
		}
	}

	return mAx
}

// MaxIndexBy search the maximum value of a collection using the given comparison function and the index of the maximum value.
// If several values of the collection are equal to the greatest value, returns the first such value.
// Returns (zero value, -1) when the collection is empty.
//
// Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.
// See https://github.com/samber/lo/issues/129
//
// Play: https://go.dev/play/p/uaUszc-c9QK
func MaxIndexBy[T any](collection []T, greater func(a, b T) bool) (T, int) {
	var (
		mAx   T
		index int
	)

	if len(collection) == 0 {
		return mAx, -1
	}

	mAx = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if greater(item, mAx) {
			mAx = item
			index = i
		}
	}

	return mAx, index
}

// Latest search the maximum time.Time of a collection.
// Returns zero value when the collection is empty.
func Latest(times ...time.Time) time.Time {
	return MaxBy(times, func(a, b time.Time) bool { return a.After(b) })
}

// LatestBy search the maximum time.Time of a collection using the given iteratee function.
// Returns zero value when the collection is empty.
func LatestBy[T any](collection []T, iteratee func(item T) time.Time) T {
	return MaxBy(collection, func(a, b T) bool { return iteratee(a).After(iteratee(b)) })
}

// First returns the first element of a collection and check for availability of the first element.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func First[T any](collection []T) (T, bool) {
	length := len(collection)

	if length == 0 {
		return Empty[T](), false
	}

	return collection[0], true
}

// FirstOrEmpty returns the first element of a collection or zero value if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func FirstOrEmpty[T any](collection []T) T {
	i, _ := First(collection)
	return i
}

// FirstOr returns the first element of a collection or the fallback value if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func FirstOr[T any](collection []T, fallback T) T {
	i, ok := First(collection)
	if !ok {
		return fallback
	}

	return i
}

// Last returns the last element of a collection or error if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func Last[T any](collection []T) (T, bool) {
	length := len(collection)

	if length == 0 {
		return Empty[T](), false
	}

	return collection[length-1], true
}

// LastOrEmpty returns the last element of a collection or zero value if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func LastOrEmpty[T any](collection []T) T {
	i, _ := Last(collection)
	return i
}

// LastOr returns the last element of a collection or the fallback value if empty.
// Play: https://go.dev/play/p/ul45Z0y2EFO
func LastOr[T any](collection []T, fallback T) T {
	i, ok := Last(collection)
	if !ok {
		return fallback
	}

	return i
}

// Nth returns the element at index `nth` of collection. If `nth` is negative, the nth element
// from the end is returned. An error is returned when nth is out of slice bounds.
// Play: https://go.dev/play/p/sHoh88KWt6B
func Nth[T any, N constraints.Integer](collection []T, nth N) (T, error) {
	n := int(nth)
	l := len(collection)
	if n >= l || -n > l {
		return Empty[T](), fmt.Errorf("nth: %d out of slice bounds", n)
	}

	if n >= 0 {
		return collection[n], nil
	}
	return collection[l+n], nil
}

// NthOr returns the element at index `nth` of collection.
// If `nth` is negative, it returns the nth element from the end.
// If `nth` is out of slice bounds, it returns the fallback value instead of an error.
// Play: https://go.dev/play/p/sHoh88KWt6B
func NthOr[T any, N constraints.Integer](collection []T, nth N, fallback T) T {
	value, err := Nth(collection, nth)
	if err != nil {
		return fallback
	}
	return value
}

// NthOrEmpty returns the element at index `nth` of collection.
// If `nth` is negative, it returns the nth element from the end.
// If `nth` is out of slice bounds, it returns the zero value (empty value) for that type.
// Play: https://go.dev/play/p/sHoh88KWt6B
func NthOrEmpty[T any, N constraints.Integer](collection []T, nth N) T {
	value, _ := Nth(collection, nth)
	return value
}

// randomIntGenerator is a function that should return a random integer in the range [0, n)
// where n is the argument passed to the randomIntGenerator.
type randomIntGenerator func(n int) int

// Sample returns a random item from collection.
// Play: https://go.dev/play/p/vCcSJbh5s6l
func Sample[T any](collection []T) T {
	return SampleBy(collection, xrand.IntN)
}

// SampleBy returns a random item from collection, using randomIntGenerator as the random index generator.
// Play: https://go.dev/play/p/HDmKmMgq0XN
func SampleBy[T any](collection []T, randomIntGenerator randomIntGenerator) T {
	size := len(collection)
	if size == 0 {
		return Empty[T]()
	}
	return collection[randomIntGenerator(size)]
}

// Samples returns N random unique items from collection.
// Play: https://go.dev/play/p/vCcSJbh5s6l
func Samples[T any, Slice ~[]T](collection Slice, count int) Slice {
	return SamplesBy(collection, count, xrand.IntN)
}

// SamplesBy returns N random unique items from collection, using randomIntGenerator as the random index generator.
// Play: https://go.dev/play/p/HDmKmMgq0XN
func SamplesBy[T any, Slice ~[]T](collection Slice, count int, randomIntGenerator randomIntGenerator) Slice {
	size := len(collection)

	cOpy := append(Slice{}, collection...)

	results := Slice{}

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i

		index := randomIntGenerator(size - i)
		results = append(results, cOpy[index])

		// Removes element.
		// It is faster to swap with last element and remove it.
		cOpy[index] = cOpy[copyLength-1]
		cOpy = cOpy[:copyLength-1]
	}

	return results
}
