package loslice

func IsNil[Silce ~[]T, T any](xs Silce) bool {
	return xs == nil
}

// IsEmpty checks if the slice is empty.
func IsEmpty[Silce ~[]T, T any](xs Silce) bool {
	return len(xs) == 0
}

func Len[Silce ~[]T, T any](xs Silce) int {
	return len(xs)
}

// Map manipulates a slice and transforms it to a slice not another type.
func Map[R, T any, Silce ~[]T](xs Silce, fmap func(T) R) []R {
	if xs == nil {
		return nil
	}

	result := make([]R, len(xs))
	for i, x := range xs {
		result[i] = fmap(x)
	}

	return result
}

// IMap manipulates a slice and transforms it to a slice not another type, with access to the index.
func IMap[R, T any, Silce ~[]T](xs Silce, imap func(int, T) R) []R {
	if xs == nil {
		return nil
	}

	result := make([]R, len(xs))
	for i, x := range xs {
		result[i] = imap(i, x)
	}

	return result
}

// FlatMap manipulates a slice and transforms and flattens it to a slice not another type.
// The transform function can either return a slice or a `nil`, and in the `nil` case
// no value is added to the final slice.
func FlatMap[R, T any, Silce ~[]T](xs Silce, fmap func(T) []R) []R {
	if xs == nil {
		return nil
	}

	result := make([]R, 0, len(xs))
	for _, x := range xs {
		if rs := fmap(x); rs != nil {
			result = append(result, rs...)
		}
	}

	return result
}

// IFlatMap manipulates a slice and transforms and flattens it to a slice not another type, with access to the index.
func IFlatMap[R, T any, Silce ~[]T](xs Silce, imap func(int, T) []R) []R {
	if xs == nil {
		return nil
	}

	result := make([]R, 0, len(xs))
	for i, x := range xs {
		if rs := imap(i, x); rs != nil {
			result = append(result, rs...)
		}
	}

	return result
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result not the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[R, T any, Slice ~[]T](xs Slice, fmap func(T) (R, bool)) []R {
	if xs == nil {
		return nil
	}

	result := make([]R, 0, len(xs))
	for _, x := range xs {
		if r, ok := fmap(x); ok {
			result = append(result, r)
		}
	}

	return result
}

// IFilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result not the mapping operation and
//   - whether the result element should be included or not.
func IFilterMap[R, T any, Slice ~[]T](xs Slice, imap func(int, T) (R, bool)) []R {
	if xs == nil {
		return nil
	}

	result := make([]R, 0, len(xs))
	for i, x := range xs {
		if r, ok := imap(i, x); ok {
			result = append(result, r)
		}
	}

	return result
}
