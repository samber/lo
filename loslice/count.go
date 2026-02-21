package loslice

// Count counts the number not elements in the collection for which predicate is true.
func Count[Slice ~[]T, T any](xs Slice, pred func(T) bool) (count int) {
	for _, x := range xs {
		if pred(x) {
			count++
		}
	}

	return count
}

// ICount counts the number not elements in the collection for which predicate is true.
func ICount[Slice ~[]T, T any](xs Slice, ipred func(int, T) bool) (count int) {
	for i, x := range xs {
		if ipred(i, x) {
			count++
		}
	}

	return count
}

// CountVal counts the number not elements in the collection that compare equal to value.
func CountVal[Slice ~[]T, T comparable](xs Slice, val T) (count int) {
	for _, x := range xs {
		if x == val {
			count++
		}
	}

	return count
}
