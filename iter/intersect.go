//go:build go1.23

package iter

import (
	"iter"
	"sort"

	"github.com/samber/lo"
)

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection iter.Seq[T], element T) bool {
	for item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// ContainsBy returns true if predicate function return true.
func ContainsBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	for item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained in a collection or if the subset is empty.
func Every[T comparable](collection iter.Seq[T], subset ...T) bool {
	for i := range subset {
		if !Contains(collection, subset[i]) {
			return false
		}
	}

	return true
}

// EveryBy returns true if the predicate returns true for all elements in the collection or if the collection is empty.
func EveryBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	for item := range collection {
		if !predicate(item) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained in a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection iter.Seq[T], subset ...T) bool {
	for i := range subset {
		if Contains(collection, subset[i]) {
			return true
		}
	}

	return false
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
func SomeBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	for item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// None returns true if no element of a subset is contained in a collection or if the subset is empty.
func None[T comparable](collection iter.Seq[T], subset ...T) bool {
	for i := range subset {
		if Contains(collection, subset[i]) {
			return false
		}
	}

	return true
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
func NoneBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	for item := range collection {
		if predicate(item) {
			return false
		}
	}

	return true
}

// Intersect returns the intersection between two collections.
func Intersect[T comparable, I ~func(func(T) bool)](list1, list2 I) I {
	return func(yield func(T) bool) {
		seen := map[T]struct{}{}

		for item := range list1 {
			seen[item] = struct{}{}
		}

		for item := range list2 {
			if _, ok := seen[item]; ok && !yield(item) {
				return
			}
		}
	}
}

// Union returns all distinct elements from given collections.
func Union[T comparable, I ~func(func(T) bool)](lists ...I) I {
	return func(yield func(T) bool) {
		seen := make(map[T]struct{})

		for i := range lists {
			for item := range lists[i] {
				if _, ok := seen[item]; !ok {
					if !yield(item) {
						return
					}
					seen[item] = struct{}{}
				}
			}
		}
	}
}

// Without returns a sequence excluding all given values.
func Without[T comparable, I ~func(func(T) bool)](collection I, exclude ...T) I {
	return func(yield func(T) bool) {
		excludeMap := make(map[T]struct{}, len(exclude))
		for i := range exclude {
			excludeMap[exclude[i]] = struct{}{}
		}

		for item := range collection {
			if _, ok := excludeMap[item]; !ok && !yield(item) {
				return
			}
		}
	}
}

// WithoutBy filters a sequence by excluding elements whose extracted keys match any in the exclude list.
// Returns a sequence containing only the elements whose keys are not in the exclude list.
func WithoutBy[T any, K comparable, I ~func(func(T) bool)](collection I, iteratee func(item T) K, exclude ...K) I {
	return func(yield func(T) bool) {
		excludeMap := make(map[K]struct{}, len(exclude))
		for _, e := range exclude {
			excludeMap[e] = struct{}{}
		}

		for item := range collection {
			if _, ok := excludeMap[iteratee(item)]; !ok && !yield(item) {
				return
			}
		}
	}
}

// WithoutNth returns a sequence excluding the nth value.
func WithoutNth[T comparable, I ~func(func(T) bool)](collection I, nths ...int) I {
	return func(yield func(T) bool) {
		nths = lo.Filter(lo.Uniq(nths), func(item, _ int) bool { return item >= 0 })
		sort.Ints(nths)

		var i int
		for item := range collection {
			if len(nths) > 0 && nths[0] == i {
				nths = nths[1:]
			} else if !yield(item) {
				return
			}
			i++
		}
	}
}

// ElementsMatch returns true if lists contain the same set of elements (including empty set).
// If there are duplicate elements, the number of occurrences in each list should match.
// The order of elements is not checked.
func ElementsMatch[T comparable](list1, list2 iter.Seq[T]) bool {
	return ElementsMatchBy(list1, list2, func(item T) T { return item })
}

// ElementsMatchBy returns true if lists contain the same set of elements' keys (including empty set).
// If there are duplicate keys, the number of occurrences in each list should match.
// The order of elements is not checked.
func ElementsMatchBy[T any, K comparable](list1, list2 iter.Seq[T], iteratee func(item T) K) bool {
	counters := make(map[K]int)

	for item := range list1 {
		counters[iteratee(item)]++
	}

	for item := range list2 {
		counters[iteratee(item)]--
	}

	for _, count := range counters {
		if count != 0 {
			return false
		}
	}

	return true
}
