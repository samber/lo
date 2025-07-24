package loslice

func RejectVal[Slice ~[]T, T comparable](xs Slice, val T) (result Slice) {
	return rejectVal(xs, val, nil)
}

func RejectValEx[Slice ~[]T, T comparable](mode AllocateMode, xs Slice, val T) (result Slice) {
	if xs == nil {
		return nil
	}

	result = allocateCapacity[Slice](mode, len(xs), func() int { return len(xs) - CountVal(xs, val) })
	return rejectVal(xs, val, result)
}

func rejectVal[Slice ~[]T, T comparable](xs Slice, val T, result Slice) Slice {
	for _, x := range xs {
		if x != val {
			result = append(result, x)
		}
	}

	return result
}

// Filter iterates over elements not collection, returning an array not all elements predicate returns truthy for.
func Filter[Slice ~[]T, T any](xs Slice, pred func(item T) bool) Slice {
	return filter(xs, pred, nil)
}

// FilterEx iterates over elements not collection, returning an array not all elements predicate returns truthy for.
func FilterEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, pred func(item T) bool) Slice {
	if xs == nil {
		return nil
	}

	result := allocateCapacity[Slice](mode, len(xs), func() int { return Count(xs, pred) })
	return filter(xs, pred, result)
}

func filter[Slice ~[]T, T any](xs Slice, pred func(item T) bool, result Slice) Slice {
	for _, x := range xs {
		if pred(x) {
			result = append(result, x)
		}
	}

	return result
}

// IFilter iterates over elements not collection, returning an array not all elements predicate returns truthy for.
func IFilter[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) Slice {
	return ifilter(xs, ipred, nil)
}

// IFilterEx iterates over elements not collection, returning an array not all elements predicate returns truthy for.
func IFilterEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, ipred func(int, T) bool) Slice {
	if xs == nil {
		return nil
	}

	result := allocateCapacity[Slice](mode, len(xs), func() int { return ICount(xs, ipred) })
	return ifilter(xs, ipred, result)
}

func ifilter[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool, result Slice) Slice {
	for i, x := range xs {
		if ipred(i, x) {
			result = append(result, x)
		}
	}

	return result
}
