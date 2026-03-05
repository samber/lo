package loslice

import "slices"

func IndexVal[Slice ~[]T, T comparable](xs Slice, val T) (index int, ok bool) {
	if i := slices.Index(xs, val); i >= 0 {
		return i, true
	}

	return
}

func Index[Slice ~[]T, T any](xs Slice, pred func(item T) bool) (index int, ok bool) {
	if i := slices.IndexFunc(xs, pred); i >= 0 {
		return i, true
	}

	return
}

func IIndex[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) (index int, ok bool) {
	for i, x := range xs {
		if ipred(i, x) {
			return i, true
		}
	}

	return
}

func RIndexVal[Slice ~[]T, T comparable](xs Slice, val T) (index int, ok bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if xs[i] == val {
			return i, true
		}
	}

	return
}

func RIndex[Slice ~[]T, T any](xs Slice, pred func(item T) bool) (index int, ok bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if pred(xs[i]) {
			return i, true
		}
	}

	return
}

func IRIndex[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) (index int, ok bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if ipred(i, xs[i]) {
			return i, true
		}
	}

	return
}

func IndicesVal[Slice ~[]T, T comparable](xs Slice, val T) (indices []int) {
	return indicesVal(xs, val, nil)
}

func IndicesValEx[Slice ~[]T, T comparable](mode AllocateMode, xs Slice, val T) (indices []int) {
	if xs == nil {
		return nil
	}

	indices = allocateCapacity[[]int](mode, len(xs), func() int { return CountVal(xs, val) })
	return indicesVal(xs, val, indices)
}

func indicesVal[Slice ~[]T, T comparable](xs Slice, val T, indices []int) []int {
	for i, x := range xs {
		if x == val {
			indices = append(indices, i)
		}
	}

	return indices
}

func RIndicesVal[Slice ~[]T, T comparable](xs Slice, val T) (indices []int) {
	indices = IndicesVal(xs, val)
	slices.Reverse(indices)

	return
}

func RIndicesValEx[Slice ~[]T, T comparable](mode AllocateMode, xs Slice, val T) (indices []int) {
	indices = IndicesValEx(mode, xs, val)
	slices.Reverse(indices)

	return
}

// Indices returns indices of elements matching the predicate.
func Indices[Slice ~[]T, T any](xs Slice, pred func(item T) bool) (indices []int) {
	return indicesImpl(xs, pred, nil)
}

// IndicesEx returns indices of elements matching the predicate, with allocation mode.
func IndicesEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, pred func(item T) bool) (indices []int) {
	if xs == nil {
		return nil
	}

	indices = allocateCapacity[[]int](mode, len(xs), func() int { return Count(xs, pred) })
	return indicesImpl(xs, pred, indices)
}

// Helper for Indices and IndicesEx.
func indicesImpl[Slice ~[]T, T any](xs Slice, pred func(item T) bool, indices []int) []int {
	for i, x := range xs {
		if pred(x) {
			indices = append(indices, i)
		}
	}
	return indices
}

func RIndices[Slice ~[]T, T any](xs Slice, pred func(item T) bool) (indices []int) {
	indices = Indices(xs, pred)
	slices.Reverse(indices)

	return
}

// RIndicesEx returns indices of elements matching the predicate, with allocation mode.
func RIndicesEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, pred func(item T) bool) (indices []int) {
	indices = IndicesEx(mode, xs, pred)
	slices.Reverse(indices)

	return
}

// IIndices returns indices of elements matching the indexed predicate.
func IIndices[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) (indices []int) {
	return iindices(xs, ipred, nil)
}

// IIndicesEx returns indices of elements matching the indexed predicate, with allocation mode.
func IIndicesEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, ipred func(int, T) bool) (indices []int) {
	if xs == nil {
		return nil
	}

	indices = allocateCapacity[[]int](mode, len(xs), func() int { return ICount(xs, ipred) })
	return iindices(xs, ipred, indices)
}

// Helper for IIndices and IIndicesEx.
func iindices[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool, indices []int) []int {
	for i, x := range xs {
		if ipred(i, x) {
			indices = append(indices, i)
		}
	}

	return indices
}

func IRIndices[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) (indices []int) {
	indices = IIndices(xs, ipred)
	slices.Reverse(indices)

	return
}

// IRIndicesEx returns indices of elements matching the indexed predicate, with allocation mode.
func IRIndicesEx[Slice ~[]T, T any](mode AllocateMode, xs Slice, ipred func(int, T) bool) (indices []int) {
	indices = IIndicesEx(mode, xs, ipred)
	slices.Reverse(indices)

	return
}

// IndicesNVal returns the first n indices of the value in the slice.
// Supports negative n to get the last n indices.
func IndicesNVal[Slice ~[]T, T comparable](xs Slice, n int, val T) (indices []int) {
	if xs == nil {
		return nil
	}

	if n == 0 {
		return IndicesVal(xs, val)
	}

	if n > 0 {
		if n >= len(xs) {
			return IndicesVal(xs, val)
		}

		indices = make([]int, 0, n)
		for i := 0; i < len(xs) && len(indices) < n; i++ {
			if xs[i] == val {
				indices = append(indices, i)
			}
		}
	} else {
		n = -n
		if n >= len(xs) {
			return RIndicesVal(xs, val)
		}

		indices = make([]int, 0, n)
		for i := len(xs) - 1; i >= 0 && len(indices) < n; i-- {
			if xs[i] == val {
				indices = append(indices, i)
			}
		}
	}

	return
}

// IndicesN returns the first n indices of the elements that satisfy the predicate in the slice.
func IndicesN[Slice ~[]T, T any](xs Slice, n int, pred func(item T) bool) (indices []int) {
	if xs == nil {
		return nil
	}

	if n == 0 {
		return Indices(xs, pred)
	}

	if n > 0 {
		if n >= len(xs) {
			return Indices(xs, pred)
		}

		indices = make([]int, 0, n)
		for i := 0; i < len(xs) && len(indices) < n; i++ {
			if pred(xs[i]) {
				indices = append(indices, i)
			}
		}
	} else {
		n = -n
		if n >= len(xs) {
			return RIndices(xs, pred)
		}

		indices = make([]int, 0, n)
		for i := len(xs) - 1; i >= 0 && len(indices) < n; i-- {
			if pred(xs[i]) {
				indices = append(indices, i)
			}
		}
	}

	return
}

// IIndicesN returns the first n indices of the elements that satisfy the predicate in the slice, with access to the index.
func IIndicesN[Slice ~[]T, T any](xs Slice, n int, ipred func(int, T) bool) (indices []int) {
	if xs == nil {
		return nil
	}

	if n == 0 {
		return IIndices(xs, ipred)
	}

	if n > 0 {
		if n >= len(xs) {
			return IIndices(xs, ipred)
		}

		indices = make([]int, 0, n)
		for i := 0; i < len(xs) && len(indices) < n; i++ {
			if ipred(i, xs[i]) {
				indices = append(indices, i)
			}
		}
	} else {
		n = -n
		if n >= len(xs) {
			return IRIndices(xs, ipred)
		}

		indices = make([]int, 0, n)
		for i := len(xs) - 1; i >= 0 && len(indices) < n; i-- {
			if ipred(i, xs[i]) {
				indices = append(indices, i)
			}
		}
	}

	return
}
