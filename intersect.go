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
func ContainsBy[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained into a collection.
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// Intersect returns the intersection between two collections.
// Elements returned will be in the order seen in list1
func Intersect[T comparable](list1 []T, list2 []T) []T {
	result := []T{}

	// store list2 values as map keys, so that lookup is O(1)
	// list2 deduplication is taken care of during this map creation
	// maintaining a boolean value allows us to ensure that when we encounter multiple intersections,
	// that the value is only appended to the result once
	list2values := make(map[T]bool)

	for _, elem := range list2 {
		list2values[elem] = false
	}

	for _, elem := range list1 {
		// this can be read as, if there's an intersection && unless we've already appended the value to the result
		if _, ok := list2values[elem]; ok && !list2values[elem] {
			// mark that we've appended this value to result
			list2values[elem] = true
			result = append(result, elem)
		}
	}

	return result
}

// Difference returns the difference between two collections.
// The first value is the collection of element absent of list2.
// The second value is the collection of element absent of list1.
func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
	left := []T{}
	right := []T{}

	seenLeft := map[T]struct{}{}
	seenRight := map[T]struct{}{}

	for _, elem := range list1 {
		seenLeft[elem] = struct{}{}
	}

	for _, elem := range list2 {
		seenRight[elem] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := seenRight[elem]; !ok {
			left = append(left, elem)
		}
	}

	for _, elem := range list2 {
		if _, ok := seenLeft[elem]; !ok {
			right = append(right, elem)
		}
	}

	return left, right
}

// Union returns all distinct elements from both collections.
// result returns will not change the order of elements relatively.
func Union[T comparable](list1 []T, list2 []T) []T {
	result := []T{}

	seen := map[T]struct{}{}
	hasAdd := map[T]struct{}{}

	for _, e := range list1 {
		seen[e] = struct{}{}
	}

	for _, e := range list2 {
		seen[e] = struct{}{}
	}

	for _, e := range list1 {
		if _, ok := seen[e]; ok {
			result = append(result, e)
			hasAdd[e] = struct{}{}
		}
	}

	for _, e := range list2 {
		if _, ok := hasAdd[e]; ok {
			continue
		}
		if _, ok := seen[e]; ok {
			result = append(result, e)
		}
	}

	return result
}
