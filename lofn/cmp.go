package lofn

func CmpFromLess[T any](less func(a, b T) bool) func(a, b T) int {
	return func(a, b T) int {
		if less(a, b) {
			return -1
		} else if less(b, a) {
			return 1
		} else {
			return 0
		}
	}
}
