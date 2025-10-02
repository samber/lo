//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo"
)

// Contains returns true if an element is present in a collection.
// Will iterate through the entire sequence if element is not found.
func Contains[T comparable](collection iter.Seq[T], element T) bool {
	return ContainsBy(collection, func(item T) bool { return item == element })
}

// ContainsBy returns true if predicate function return true.
// Will iterate through the entire sequence if predicate never returns true.
func ContainsBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	return IsNotEmpty(Filter(collection, predicate))
}

// Every returns true if all elements of a subset are contained in a collection or if the subset is empty.
// Will iterate through the entire sequence if subset elements always match.
func Every[T comparable](collection iter.Seq[T], subset ...T) bool {
	if len(subset) == 0 {
		return true
	}

	set := lo.Keyify(subset)
	for item := range collection {
		if _, ok := set[item]; ok {
			delete(set, item)
			if len(set) == 0 {
				return true
			}
		}
	}

	return false
}

// EveryBy returns true if the predicate returns true for all elements in the collection or if the collection is empty.
// Will iterate through the entire sequence if predicate never returns false.
func EveryBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	return IsEmpty(Reject(collection, predicate))
}

// Some returns true if at least 1 element of a subset is contained in a collection.
// If the subset is empty Some returns false.
// Will iterate through the entire sequence if subset elements never match.
func Some[T comparable](collection iter.Seq[T], subset ...T) bool {
	return SomeBy(collection, func(item T) bool { return lo.Contains(subset, item) })
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
// Will iterate through the entire sequence if predicate never returns true.
func SomeBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	return IsNotEmpty(Filter(collection, predicate))
}

// None returns true if no element of a subset is contained in a collection or if the subset is empty.
// Will iterate through the entire sequence if subset elements never match.
func None[T comparable](collection iter.Seq[T], subset ...T) bool {
	return NoneBy(collection, func(item T) bool { return lo.Contains(subset, item) })
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
// Will iterate through the entire sequence if predicate never returns true.
func NoneBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool {
	return IsEmpty(Filter(collection, predicate))
}

// Intersect returns the intersection between given collections.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func Intersect[T comparable, I ~func(func(T) bool)](lists ...I) I { //nolint:gocyclo
	if len(lists) == 0 {
		return I(Empty[T]())
	}

	if len(lists) == 1 {
		return lists[0]
	}

	return func(yield func(T) bool) {
		seen := make(map[T]bool)

		for i := len(lists) - 1; i >= 0; i-- {
			if i == len(lists)-1 {
				for item := range lists[i] {
					seen[item] = true
				}
				continue
			}

			if i == 0 {
				for item := range lists[0] {
					if _, ok := seen[item]; ok {
						if !yield(item) {
							return
						}
						delete(seen, item)
					}
				}
				continue
			}

			for k := range seen {
				seen[k] = false
			}

			for item := range lists[i] {
				if _, ok := seen[item]; ok {
					seen[item] = true
				}
			}

			for k, v := range seen {
				if !v {
					delete(seen, k)
				}
			}

			if len(seen) == 0 {
				return
			}
		}
	}
}

// Union returns all distinct elements from given collections.
// Will allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
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
// Will allocate a map large enough to hold all distinct excludes.
func Without[T comparable, I ~func(func(T) bool)](collection I, exclude ...T) I {
	return WithoutBy(collection, func(item T) T { return item }, exclude...)
}

// WithoutBy filters a sequence by excluding elements whose extracted keys match any in the exclude list.
// Returns a sequence containing only the elements whose keys are not in the exclude list.
// Will allocate a map large enough to hold all distinct excludes.
func WithoutBy[T any, K comparable, I ~func(func(T) bool)](collection I, transform func(item T) K, exclude ...K) I {
	set := lo.Keyify(exclude)
	return Reject(collection, func(item T) bool { return lo.HasKey(set, transform(item)) })
}

// WithoutNth returns a sequence excluding the nth value.
// Will allocate a map large enough to hold all distinct nths.
func WithoutNth[T comparable, I ~func(func(T) bool)](collection I, nths ...int) I {
	set := lo.Keyify(nths)
	return RejectI(collection, func(_ T, index int) bool { return lo.HasKey(set, index) })
}

// ElementsMatch returns true if lists contain the same set of elements (including empty set).
// If there are duplicate elements, the number of occurrences in each list should match.
// The order of elements is not checked.
// Will iterate through each sequence before returning and allocate a map large enough to hold all distinct elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func ElementsMatch[T comparable](list1, list2 iter.Seq[T]) bool {
	return ElementsMatchBy(list1, list2, func(item T) T { return item })
}

// ElementsMatchBy returns true if lists contain the same set of elements' keys (including empty set).
// If there are duplicate keys, the number of occurrences in each list should match.
// The order of elements is not checked.
// Will iterate through each sequence before returning and allocate a map large enough to hold all distinct transformed elements.
// Long heterogeneous input sequences can cause excessive memory usage.
func ElementsMatchBy[T any, K comparable](list1, list2 iter.Seq[T], transform func(item T) K) bool {
	counters := make(map[K]int)

	for item := range list1 {
		counters[transform(item)]++
	}

	for item := range list2 {
		counters[transform(item)]--
	}

	for _, count := range counters {
		if count != 0 {
			return false
		}
	}

	return true
}
