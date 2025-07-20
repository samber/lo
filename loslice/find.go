package loslice

import "slices"

func IndexVal[T comparable, Slice ~[]T](xs Slice, val T) (index int, ok bool) {
	if i := slices.Index(xs, val); i >= 0 {
		return i, true
	}

	return
}

func Index[T any, Slice ~[]T](xs Slice, pred func(item T) bool) (index int, ok bool) {
	if i := slices.IndexFunc(xs, pred); i >= 0 {
		return i, true
	}

	return
}

func IIndex[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) (index int, ok bool) {
	for i, x := range xs {
		if ipred(i, x) {
			return i, true
		}
	}

	return
}

func RIndexVal[T comparable, Slice ~[]T](xs Slice, val T) (index int, ok bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if xs[i] == val {
			return i, true
		}
	}

	return
}

func RIndex[T any, Slice ~[]T](xs Slice, pred func(item T) bool) (index int, ok bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if pred(xs[i]) {
			return i, true
		}
	}

	return
}

func IRIndex[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) (index int, ok bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if ipred(i, xs[i]) {
			return i, true
		}
	}

	return
}

func IndicesVal[T comparable, Slice ~[]T](xs Slice, val T) (indices []int) {
	indices = make([]int, 0, CountVal(xs, val))
	for i, x := range xs {
		if x == val {
			indices = append(indices, i)
		}
	}

	return
}

func RIndicesVal[T comparable, Slice ~[]T](xs Slice, val T) (indices []int) {
	indices = IndicesVal(xs, val)
	slices.Reverse(indices)

	return
}

func Indices[T any, Slice ~[]T](xs Slice, pred func(item T) bool) (indices []int) {
	indices = make([]int, 0, Count(xs, pred))
	for i, x := range xs {
		if pred(x) {
			indices = append(indices, i)
		}
	}

	return
}

func RIndices[T any, Slice ~[]T](xs Slice, pred func(item T) bool) (indices []int) {
	indices = Indices(xs, pred)
	slices.Reverse(indices)

	return
}

func IIndices[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) (indices []int) {
	indices = make([]int, 0, ICount(xs, ipred))
	for i, x := range xs {
		if ipred(i, x) {
			indices = append(indices, i)
		}
	}

	return
}

func IRIndices[T any, Slice ~[]T](xs Slice, ipred func(int, T) bool) (indices []int) {
	indices = IIndices(xs, ipred)
	slices.Reverse(indices)

	return
}

// IndicesNVal returns the first n indices of the value in the slice.
// Supports negative n to get the last n indices.
func IndicesNVal[T comparable, Slice ~[]T](xs Slice, n int, val T) (indices []int) {
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
func IndicesN[T any, Slice ~[]T](xs Slice, n int, pred func(item T) bool) (indices []int) {
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
func IIndicesN[T any, Slice ~[]T](xs Slice, n int, ipred func(int, T) bool) (indices []int) {
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
