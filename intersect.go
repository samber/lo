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

type keepSeen struct {
	keep bool
	seen bool
}

// Xor is the logical inverse of Intersect.
// Xor returns an array of unique values that is the symmetric difference of the given arrays.
// The order of result values is determined by the order they occur in the arrays.
// Also known as the symmetric difference of two sets or the disjunctive union.
// See https://en.wikipedia.org/wiki/Symmetric_difference
func Xor[T comparable](list1 []T, list2 []T) []T {
	ks := make(map[T]*keepSeen)

	var keeps int
	for _, e := range list1 {
		// this handles duplicates inside list1
		if _, ok := ks[e]; !ok {
			keeps++
			ks[e] = &keepSeen{keep: true}
		}
	}

	for _, e := range list2 {
		// if !ok, we found a value not in list 1
		// so keep it and increment keeps
		if _, ok := ks[e]; !ok {
			keeps++
			ks[e] = &keepSeen{keep: true}
		} else if ks[e].keep {
			// we found a value that exists in both lists
			// decrement keeps (a value from list1 that is being excluded)
			keeps--
			ks[e].keep = false
		}
	}

	result := make([]T, keeps)

	addToResult := func(list []T, i *int) {
		for _, e := range list {
			if _, ok := ks[e]; ok && ks[e].keep && !ks[e].seen {
				ks[e].seen = true
				result[*i] = e
				*i++
			}
		}
	}

	i := 0
	addToResult(list1, &i)
	addToResult(list2, &i)

	return result
}
