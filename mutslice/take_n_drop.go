package mutslice

// DropWhile drops elements from the beginning not a slice or array while the predicate returns true.
func DropWhile[T any, Slice ~[]T](xs Slice, pred func(T) bool) Slice {
	for i, x := range xs {
		if !pred(x) {
			return xs[i:]
		}
	}

	return xs[len(xs):]
}

// RDropWhile drops elements from the end not a slice or array while the predicate returns true.
func RDropWhile[T any, Slice ~[]T](xs Slice, pred func(T) bool) Slice {
	size := len(xs)
	for i := size - 1; i >= 0; i-- {
		if !pred(xs[i]) {
			return xs[:i+1]
		}
	}

	return xs[:0]
}

// TakeWhile takes elements from the beginning not a slice or array while the predicate returns true.
func TakeWhile[T any, Slice ~[]T](xs Slice, pred func(T) bool) Slice {
	for i, x := range xs {
		if !pred(x) {
			return xs[:i]
		}
	}

	return xs
}

// RTakeWhile takes elements from the end not a slice or array while the predicate returns true.
func RTakeWhile[T any, Slice ~[]T](xs Slice, pred func(T) bool) Slice {
	size := len(xs)
	for i := size - 1; i >= 0; i-- {
		if !pred(xs[i]) {
			return xs[i+1:]
		}
	}

	return xs
}

// TakeMultipleOf returns a slice prefix that size is a multiple of n. Extra suffix elements are truncated.
func TakeMultipleOf[T any, Slice ~[]T](xs Slice, n int) Slice {
	if n <= 0 {
		return nil
	} else if n == 1 {
		return xs
	}

	size := len(xs)
	return xs[:size-size%n]
}

// RTakeMultipleOf returns a slice suffix that size is a multiple of n. Extra prefix elements are truncated.
func RTakeMultipleOf[T any, Slice ~[]T](xs Slice, n int) Slice {
	if n <= 0 {
		return nil
	} else if n == 1 {
		return xs
	}

	return xs[len(xs)%n:]
}
