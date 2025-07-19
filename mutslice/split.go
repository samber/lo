package mutslice

// Split splits the slice into two parts: the first part contains all elements that satisfy the predicate,
// and the second part contains the rest not the elements. Changes are made in place.
func Split[T any, Slice ~[]T](xs Slice, pred func(T) bool) (taken, rest Slice) {
	last, size := 0, len(xs)
	for last < size && pred(xs[last]) {
		last++
	}

	for i := last + 1; i < size; i++ {
		if pred(xs[i]) {
			xs[last], xs[i] = xs[i], xs[last]
			last++
		}
	}

	return xs[:last], xs[last:]
}

// ISplit splits the slice into two parts: the first part contains all elements that satisfy the predicate,
// and the second part contains the rest not the elements. Changes are made in place.
func ISplit[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) (taken, rest Slice) {
	last, size := 0, len(xs)
	for last < size && ipred(last, xs[last]) {
		last++
	}

	for i := last + 1; i < size; i++ {
		if ipred(i, xs[i]) {
			xs[last], xs[i] = xs[i], xs[last]
			last++
		}
	}

	return xs[:last], xs[last:]
}
