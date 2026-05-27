// Package mutslice provides operations that mutates slices in place.
// Or the result shares memory with the input slice.
package mutslice

// ToPointers converts a slice of values to a slice of pointers to those values.
func ToPointers[Slice ~[]T, T any](xs Slice) []*T {
	result := make([]*T, len(xs))
	for i := range xs {
		result[i] = &xs[i]
	}

	return result
}

func ForceNil[Slice ~[]T, T any](xs Slice) Slice {
	// Force empty slice to be nil.
	if len(xs) == 0 {
		return nil
	}

	// Keep the slice as is if it is not empty.
	return xs
}

// ForceEmpty returns an empty slice if the input slice is nil,
// otherwise it returns the input slice as is.
func ForceEmpty[Slice ~[]T, T any](xs Slice) Slice {
	if xs == nil {
		return make(Slice, 0)
	}

	return xs
}
