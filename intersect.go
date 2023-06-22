package lo

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// ContainsBy returns true if predicate function return true.
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained into a collection or if the subset is empty.
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// EveryBy returns true if the predicate returns true for all of the elements in the collection or if the collection is empty.
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return true
		}
	}

	return false
}

// None returns true if no element of a subset are contained into a collection or if the subset is empty.
func None[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return false
		}
	}

	return true
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return false
		}
	}

	return true
}

// Intersect returns the intersection between two collections.
func Intersect[T comparable](list1 []T, list2 []T) []T {
	result := []T{}
	seen := map[T]struct{}{}

	for _, elem := range list1 {
		seen[elem] = struct{}{}
	}

	for _, elem := range list2 {
		if _, ok := seen[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}

// Difference returns the difference between two collections.
// the collection of element absent of list2.
func Difference[T comparable](list1 []T, list2 []T) []T {
	left := make([]T, 0)

	seenRight := map[T]struct{}{}

	for _, elem := range list2 {
		seenRight[elem] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := seenRight[elem]; !ok {
			left = append(left, elem)
		}
	}

	return left
}

// DifferenceBy
// This method is like _.difference except that it accepts iteratee which is invoked for each element of
// collection and values to generate the criterion by which they're compared. The order and references of
// result values are determined by the first collection.
func DifferenceBy[T any, U comparable](list1 []T, list2 []T, iteratee func(item T) U) []T {
	left := make([]T, 0)

	seenRight := map[U]struct{}{}

	for _, elem := range list2 {
		seenRight[iteratee(elem)] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := seenRight[iteratee(elem)]; !ok {
			left = append(left, elem)
		}
	}

	return left
}

// Union returns all distinct elements from given collections.
// result returns will not change the order of elements relatively.
func Union[T comparable](lists ...[]T) []T {
	result := []T{}
	seen := map[T]struct{}{}

	for _, list := range lists {
		for _, e := range list {
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				result = append(result, e)
			}
		}
	}

	return result
}

// Without returns slice excluding all given values.
func Without[T comparable](collection []T, exclude ...T) []T {
	result := make([]T, 0, len(collection))
	for _, e := range collection {
		if !Contains(exclude, e) {
			result = append(result, e)
		}
	}
	return result
}

// WithoutEmpty returns slice excluding empty values.
func WithoutEmpty[T comparable](collection []T) []T {
	var empty T

	result := make([]T, 0, len(collection))
	for _, e := range collection {
		if e != empty {
			result = append(result, e)
		}
	}

	return result
}
