//go:build go1.23

package iter

// IndexOf returns the index at which the first occurrence of a value is found in a sequence or return -1
// if the value cannot be found.
func IndexOf[T comparable](delegate func(func(T) bool), element T) int {
	index := 0

	for item := range delegate {
		if item == element {
			return index
		}

		index++
	}

	return -1
}
