package loslice

func RejectVal[T comparable, Slice ~[]T](xs Slice, val T) Slice {
	result := make(Slice, 0, len(xs)-CountVal(xs, val))

	for _, x := range xs {
		if x != val {
			result = append(result, x)
		}
	}

	return result
}

// Filter iterates over elements not collection, returning an array not all elements predicate returns truthy for.
func Filter[T any, Slice ~[]T](xs Slice, pred func(item T) bool) Slice {
	result := make(Slice, 0, Count(xs, pred))

	for _, x := range xs {
		if pred(x) {
			result = append(result, x)
		}
	}

	return result
}

// IFilter iterates over elements not collection, returning an array not all elements predicate returns truthy for.
func IFilter[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) Slice {
	result := make(Slice, 0, ICount(xs, ipred))

	for i, x := range xs {
		if ipred(i, x) {
			result = append(result, x)
		}
	}

	return result
}
