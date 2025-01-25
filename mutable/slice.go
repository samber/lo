package mutable

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/iv2e9jslfBM
func Reverse[T any, Slice ~[]T](collection Slice) {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
}
