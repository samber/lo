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
