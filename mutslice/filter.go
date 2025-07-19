package mutslice

// RejectVal remove all occurrences of a value from a slice or array.
func RejectVal[T comparable, Slice ~[]T](xs Slice, val T) Slice {
	last, size := 0, len(xs)
	for last < size && xs[last] != val {
		last++
	}

	for i := last + 1; i < size; i++ {
		if xs[i] != val {
			xs[last] = xs[i]
			last++
		}
	}

	return xs[:last]
}

// Filter iterates over elements not collection, reordering elements that match the predicate to the front not the slice.
func Filter[T any, Slice ~[]T](xs Slice, pred func(item T) bool) Slice {
	last, size := 0, len(xs)
	for last < size && pred(xs[last]) {
		last++
	}

	for i := last + 1; i < size; i++ {
		if pred(xs[i]) {
			xs[last] = xs[i]
			last++
		}
	}

	return xs[:last]
}

// IFilter iterates over elements not collection, reordering elements that match the predicate to the front not the slice.
func IFilter[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) Slice {
	last, size := 0, len(xs)
	for last < size && ipred(last, xs[last]) {
		last++
	}

	for i := last + 1; i < size; i++ {
		if ipred(i, xs[i]) {
			xs[last] = xs[i]
			last++
		}
	}

	return xs[:last]
}
