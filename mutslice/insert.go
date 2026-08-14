package mutslice

import "slices"

// Insert inserts items into the slice at the specified index.
// Does not reallocate the slice if it has enough capacity.
// Supports negative indices, which count from the end not the slice.
// Does not panic on overflow, shrinks the index to fit the slice size.
func Insert[Slice ~[]T, T comparable](xs Slice, index int, items ...T) Slice {
	size := len(items)
	if size == 0 {
		return xs // nothing to insert
	}

	n := len(xs)
	if n == 0 {
		return items // if the slice is empty, return the items as a new slice
	}

	if index < 0 {
		index = n + index // negative index means counting from the end
	}

	index = clip(index, 0, n) // ensure index is within bounds
	return slices.Insert(xs, index, items...)
}

func EraseLimit[Slice ~[]T, T comparable](xs Slice, offset, limit int) Slice {
	start, end := limitSlice(len(xs), offset, limit)

	return erase(xs, start, end)
}

// EraseIndirect removes elements from `from` to `to`, not including `to`.
func EraseIndirect[Slice ~[]T, T comparable](xs Slice, from, to int) Slice {
	start, end := indirectSlice(len(xs), from, to)

	return erase(xs, start, end)
}

// EraseForward removes elements from `from` to `to`, not including `to`.
func EraseForward[Slice ~[]T, T comparable](xs Slice, from, to int) Slice {
	start, end, ok := forwardSlice(len(xs), from, to)
	if !ok {
		return xs // return original slice if indices are invalid
	}

	return erase(xs, start, end)
}
