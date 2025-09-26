package lo

// Contains returns true if an element is present in a collection.
// Play: https://go.dev/play/p/W1EvyqY6t9j
func Contains[T comparable](collection []T, element T) bool {
	for i := range collection {
		if collection[i] == element {
			return true
		}
	}

	return false
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
	for i := range subset {
		if !Contains(collection, subset[i]) {
			return false
		}
	}

	return true
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
	for i := range subset {
		if Contains(collection, subset[i]) {
			return true
		}
	}

	return false
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
	for i := range subset {
		if Contains(collection, subset[i]) {
			return false
		}
	}

	return true
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

// Intersect returns the intersection between two collections.
// Play: https://go.dev/play/p/uuElL9X9e58
func Intersect[T comparable, Slice ~[]T](list1, list2 Slice) Slice {
	result := Slice{}
	seen := map[T]struct{}{}

	for i := range list1 {
		seen[list1[i]] = struct{}{}
	}

	for i := range list2 {
		if _, ok := seen[list2[i]]; ok {
			result = append(result, list2[i])
		}
	}

	return result
}

// Difference returns the difference between two collections.
// The first value is the collection of elements absent from list2.
// The second value is the collection of elements absent from list1.
// Play: https://go.dev/play/p/pKE-JgzqRpz
func Difference[T comparable, Slice ~[]T](list1, list2 Slice) (Slice, Slice) {
	left := Slice{}
	right := Slice{}

	seenLeft := map[T]struct{}{}
	seenRight := map[T]struct{}{}

	for i := range list1 {
		seenLeft[list1[i]] = struct{}{}
	}

	for i := range list2 {
		seenRight[list2[i]] = struct{}{}
	}

	for i := range list1 {
		if _, ok := seenRight[list1[i]]; !ok {
			left = append(left, list1[i])
		}
	}

	for i := range list2 {
		if _, ok := seenLeft[list2[i]]; !ok {
			right = append(right, list2[i])
		}
	}

	return left, right
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
			if _, ok := seen[lists[i][j]]; !ok {
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
	excludeMap := make(map[T]struct{}, len(exclude))
	for i := range exclude {
		excludeMap[exclude[i]] = struct{}{}
	}

	result := make(Slice, 0, len(collection))
	for i := range collection {
		if _, ok := excludeMap[collection[i]]; !ok {
			result = append(result, collection[i])
		}
	}
	return result
}

// WithoutBy filters a slice by excluding elements whose extracted keys match any in the exclude list.
// Returns a new slice containing only the elements whose keys are not in the exclude list.
// Play: https://go.dev/play/p/VgWJOF01NbJ
func WithoutBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K, exclude ...K) Slice {
	excludeMap := make(map[K]struct{}, len(exclude))
	for _, e := range exclude {
		excludeMap[e] = struct{}{}
	}

	result := make(Slice, 0, len(collection))
	for _, item := range collection {
		if _, ok := excludeMap[iteratee(item)]; !ok {
			result = append(result, item)
		}
	}
	return result
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
	length := len(collection)

	toRemove := make(map[int]struct{}, len(nths))
	for i := range nths {
		if nths[i] >= 0 && nths[i] <= length-1 {
			toRemove[nths[i]] = struct{}{}
		}
	}

	result := make(Slice, 0, len(collection))
	for i := range collection {
		if _, ok := toRemove[i]; !ok {
			result = append(result, collection[i])
		}
	}

	return result
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
