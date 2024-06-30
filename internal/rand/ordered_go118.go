//go:build !go1.22

package rand

import "math/rand"

func Shuffle(n int, swap func(i, j int)) {
	rand.Shuffle(n, swap)
}

func IntN(n int) int {
	return rand.Intn(n)
}
