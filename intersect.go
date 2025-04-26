package lo

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for i := range collection {
		if collection[i] == element {
			return true
		}
	}

	return false
}

// ContainsBy returns true if predicate function return true.
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if predicate(collection[i]) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained into a collection or if the subset is empty.
func Every[T comparable](collection []T, subset []T) bool {
	for i := range subset {
		if !Contains(collection, subset[i]) {
			return false
		}
	}

	return true
}

// EveryBy returns true if the predicate returns true for all elements in the collection or if the collection is empty.
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if !predicate(collection[i]) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection []T, subset []T) bool {
	for i := range subset {
		if Contains(collection, subset[i]) {
			return true
		}
	}

	return false
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if predicate(collection[i]) {
			return true
		}
	}

	return false
}

// None returns true if no element of a subset are contained into a collection or if the subset is empty.
func None[T comparable](collection []T, subset []T) bool {
	for i := range subset {
		if Contains(collection, subset[i]) {
			return false
		}
	}

	return true
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	for i := range collection {
		if predicate(collection[i]) {
			return false
		}
	}

	return true
}

// Intersect returns the intersection between two collections.
func Intersect[T comparable, Slice ~[]T](list1 Slice, list2 Slice) Slice {
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
// The first value is the collection of element absent of list2.
// The second value is the collection of element absent of list1.
func Difference[T comparable, Slice ~[]T](list1 Slice, list2 Slice) (Slice, Slice) {
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

// Without returns slice excluding all given values.
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
// It returns a new slice containing only the elements whose keys are not in the exclude list.
func WithoutBy[T any, K comparable](collection []T, iteratee func(item T) K, exclude ...K) []T {
	excludeMap := make(map[K]struct{}, len(exclude))
	for _, e := range exclude {
		excludeMap[e] = struct{}{}
	}

	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := excludeMap[iteratee(item)]; !ok {
			result = append(result, item)
		}
	}
	return result
}

// WithoutEmpty returns slice excluding zero values.
//
// Deprecated: Use lo.Compact instead.
func WithoutEmpty[T comparable, Slice ~[]T](collection Slice) Slice {
	return Compact(collection)
}

// WithoutNth returns slice excluding nth value.
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
// If there are duplicate elements, the number of appearances of each of them in both lists should match.
// The order of elements is not checked.
func ElementsMatch[T comparable, Slice ~[]T](list1 Slice, list2 Slice) bool {
	return ElementsMatchBy(list1, list2, func(item T) T { return item })
}

// ElementsMatchBy returns true if lists contain the same set of elements' keys (including empty set).
// If there are duplicate keys, the number of appearances of each of them in both lists should match.
// The order of elements is not checked.
func ElementsMatchBy[T any, K comparable](list1 []T, list2 []T, iteratee func(item T) K) bool {
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
