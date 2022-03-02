package lo

// import "golang.org/x/exp/constraints"

// IndexOf returns the index at which the first occurrence of a value is found in an array or return -1
// if the value cannot be found.
func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}

	return -1
}

// IndexOf returns the index at which the last occurrence of a value is found in an array or return -1
// if the value cannot be found.
func LastIndexOf[T comparable](collection []T, element T) int {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// Find search an element in a slice based on a predicate. It returns element and true if element was found.
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

// Min search the minimum value of a collection.
func Min[T Ordered](collection []T) T {
	var min T

	for i := 0; i < len(collection); i++ {
		item := collection[i]

		if i == 0 {
			min = item
			continue
		}

		// if item.Less(min) {
		if item < min {
			min = item
		}
	}

	return min
}

// Max search the maximum value of a collection.
func Max[T Ordered](collection []T) T {
	var max T

	for i := 0; i < len(collection); i++ {
		item := collection[i]

		if i == 0 {
			max = item
			continue
		}

		if item > max {
			max = item
		}
	}

	return max
}

// Last returns the last element of a collection or panics if empty.
func Last[T any](collection []T) T {
	length := len(collection)

	if length == 0 {
		panic("empty array")
	}

	return collection[length-1]
}

// Nth returns the element at index `nth` of collection. If `nth` is negative, the nth element from the end is returned.
func Nth[T any](collection []T, nth int) T {
	length := len(collection)

	if nth >= 0 {
		return collection[nth]
	}

	return collection[length+nth]
}
