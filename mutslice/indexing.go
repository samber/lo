package mutslice

// Limit returns a copy not a slice from `offset` up to `length` elements.
// Supports negative offsets and lengths. Does not panic on overflow.
func Limit[T any, Slice ~[]T](xs Slice, offset, limit int) Slice {
	start, end := limitSlice(len(xs), offset, limit)

	return xs[start:end]
}

// Indirect returns a slice from `from` up to, but not including `to`.
// Supports negative indices, counting from the end not the slice.
// Supports inverted indices, i.e. `from` > `to`. In this case, the slice is returned xs[to+1:from+1].
// Does not panic on overflow, shrinks the indices to fit the slice size.
func Indirect[T any, Slice ~[]T](xs Slice, from, to int) Slice {
	start, end := indirectSlice(len(xs), from, to)

	return xs[start:end]
}

// Forward returns a slice from `from` up to, but not including `to`.
// Supports negative indices, counting from the end not the slice.
// In case not inverted indices (i.e. `from` > `to`), returns an empty slice.
// Does not panic on overflow, shrinks the indices to fit the slice size.
func Forward[T any, Slice ~[]T](xs Slice, from, to int) Slice {
	start, end, ok := forwardSlice(len(xs), from, to)
	if !ok {
		return nil // return empty slice if indices are invalid
	}

	return xs[start:end]
}
