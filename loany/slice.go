package loany

func Slice[Slice ~[]T, T any](xs Slice) []any {
	result := make([]any, len(xs))
	for i, x := range xs {
		result[i] = x
	}

	return result
}

// TypedSlice converts a slice of any type to a slice of a specific type T.
// Values that cannot be converted to T will be omitted.
func TypedSlice[T any, Slice ~[]any](xs Slice) []T {
	result := make([]T, 0, len(xs))
	for _, x := range xs {
		if v, ok := x.(T); ok {
			result = append(result, v)
		}
	}

	return result
}
