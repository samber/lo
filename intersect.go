package lo

// Contains returns true if an element is present in a collection.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func Contains[T comparable](collection []T, element T) bool {
	return ContainsBy(collection, func(item T) bool { return item == element })
}

// ContainsBy returns true if predicate function return true.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if predicate(collection[i]) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained in a collection or if the subset is empty.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func Every[T comparable](collection, subset []T) bool {
	return len(subset) == 0 || EveryBy(subset, Partial(HasKey, Keyify(collection)))
}

// EveryBy returns true if the predicate returns true for all elements in the collection or if the collection is empty.
// Play: https://go.dev/play/p/dn1-vhHsq9x
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if !predicate(collection[i]) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained in a collection.
// If the subset is empty Some returns false.
// Play: https://go.dev/play/p/Lj4ceFkeT9V
func Some[T comparable](collection, subset []T) bool {
	return len(subset) > 0 && SomeBy(collection, Partial(HasKey, Keyify(subset)))
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
// Play: https://go.dev/play/p/DXF-TORBudx
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if predicate(collection[i]) {
			return true
		}
	}

	return false
}

// None returns true if no element of a subset is contained in a collection or if the subset is empty.
// Play: https://go.dev/play/p/fye7JsmxzPV
func None[T comparable](collection, subset []T) bool {
	return len(subset) == 0 || NoneBy(collection, Partial(HasKey, Keyify(subset)))
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
// Play: https://go.dev/play/p/O64WZ32H58S
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if predicate(collection[i]) {
			return false
		}
	}

	return true
}

// Intersect returns the intersection between collections.
// Play: https://go.dev/play/p/uuElL9X9e58
func Intersect[T comparable, Slice ~[]T](lists ...Slice) Slice {
	return IntersectBy(func(item T) T { return item }, lists...)
}

// IntersectBy returns the intersection between two collections using a custom key selector function.
func IntersectBy[T any, K comparable, Slice ~[]T](transform func(T) K, lists ...Slice) Slice {
	if len(lists) == 0 {
		return Slice{}
	}

	if len(lists) == 1 {
		return lists[0]
	}

	seen := make(map[K]bool)

	for i := len(lists) - 1; i >= 0; i-- {
		if i == len(lists)-1 {
			for _, item := range lists[i] {
				k := transform(item)
				seen[k] = true
			}
			continue
		}

		if i == 0 {
			result := make(Slice, 0, len(seen))
			for _, item := range lists[0] {
				k := transform(item)
				if HasKey(seen, k) {
					result = append(result, item)
					delete(seen, k)
				}
			}
			return result
		}

		for k := range seen {
			seen[k] = false
		}

		for _, item := range lists[i] {
			k := transform(item)
			if HasKey(seen, k) {
				seen[k] = true
			}
		}

		for k, v := range seen {
			if !v {
				delete(seen, k)
			}
		}

		if len(seen) == 0 {
			break
		}
	}

	return Slice{}
}

// Difference returns the difference between two collections.
// The first value is the collection of elements absent from list2.
// The second value is the collection of elements absent from list1.
// Play: https://go.dev/play/p/pKE-JgzqRpz
func Difference[T comparable, Slice ~[]T](list1, list2 Slice) (Slice, Slice) {
	result1 := _reject(list1, Partial(HasKey, Keyify(list2)))
	result2 := _reject(list2, Partial(HasKey, Keyify(list1)))
	return result1, result2
}

// Union returns all distinct elements from given collections.
// result returns will not change the order of elements relatively.
// Play: https://go.dev/play/p/DI9RVEB_qMK
func Union[T comparable, Slice ~[]T](lists ...Slice) Slice {
	var capLen int

	for _, list := range lists {
		capLen += len(list)
	}

	result := make(Slice, 0, capLen)
	seen := make(map[T]struct{}, capLen)

	for i := range lists {
		for j := range lists[i] {
			if !HasKey(seen, lists[i][j]) {
				seen[lists[i][j]] = struct{}{}
				result = append(result, lists[i][j])
			}
		}
	}

	return result
}

// Without returns a slice excluding all given values.
// Play: https://go.dev/play/p/5j30Ux8TaD0
func Without[T comparable, Slice ~[]T](collection Slice, exclude ...T) Slice {
	return WithoutBy(collection, func(item T) T { return item }, exclude...)
}

// WithoutBy filters a slice by excluding elements whose extracted keys match any in the exclude list.
// Returns a new slice containing only the elements whose keys are not in the exclude list.
// Play: https://go.dev/play/p/VgWJOF01NbJ
func WithoutBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K, exclude ...K) Slice {
	excludeMap := Keyify(exclude)
	return _reject(collection, func(item T) bool { return HasKey(excludeMap, iteratee(item)) })
}

// WithoutEmpty returns a slice excluding zero values.
//
// Deprecated: Use lo.Compact instead.
func WithoutEmpty[T comparable, Slice ~[]T](collection Slice) Slice {
	return Compact(collection)
}

// WithoutNth returns a slice excluding the nth value.
// Play: https://go.dev/play/p/5g3F9R2H1xL
func WithoutNth[T comparable, Slice ~[]T](collection Slice, nths ...int) Slice {
	toRemove := Keyify(nths)
	return Reject(collection, func(_ T, index int) bool { return HasKey(toRemove, index) })
}

// ElementsMatch returns true if lists contain the same set of elements (including empty set).
// If there are duplicate elements, the number of occurrences in each list should match.
// The order of elements is not checked.
// Play: https://go.dev/play/p/XWSEM4Ic_t0
func ElementsMatch[T comparable, Slice ~[]T](list1, list2 Slice) bool {
	return ElementsMatchBy(list1, list2, func(item T) T { return item })
}

// ElementsMatchBy returns true if lists contain the same set of elements' keys (including empty set).
// If there are duplicate keys, the number of occurrences in each list should match.
// The order of elements is not checked.
// Play: https://go.dev/play/p/XWSEM4Ic_t0
func ElementsMatchBy[T any, K comparable](list1, list2 []T, iteratee func(item T) K) bool {
	if len(list1) != len(list2) {
		return false
	}

	if len(list1) == 0 {
		return true
	}

	counters := make(map[K]int, len(list1))

	for _, el := range list1 {
		counters[iteratee(el)]++
	}

	for _, el := range list2 {
		counters[iteratee(el)]--
	}

	for _, count := range counters {
		if count != 0 {
			return false
		}
	}

	return true
}
