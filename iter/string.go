//go:build go1.23

package iter

import "iter"

// ChunkString returns a sequence of strings split into groups of length size. If the string can't be split evenly,
// the final chunk will be the remaining characters.
func ChunkString[T ~string](str T, size int) iter.Seq[T] {
	if size <= 0 {
		panic("iter.ChunkString: size must be greater than 0")
	}

	return func(yield func(T) bool) {
		if len(str) == 0 {
			yield("")
			return
		}

		if size >= len(str) {
			yield(str)
			return
		}

		currentLen := 0
		currentStart := 0
		for i := range str {
			if currentLen == size {
				if !yield(str[currentStart:i]) {
					return
				}
				currentLen = 0
				currentStart = i
			}
			currentLen++
		}
		yield(str[currentStart:])
	}
}
