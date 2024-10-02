//go:build !go1.22

package rand

import "math/rand"

func Shuffle(n int, swap func(i, j int)) {
	rand.Shuffle(n, swap)
}

func IntN(n int) int {
	// bearer:disable go_gosec_crypto_weak_random
	return rand.Intn(n)
}

func Int64() int64 {
	// bearer:disable go_gosec_crypto_weak_random
	n := rand.Int63()
    
	// bearer:disable go_gosec_crypto_weak_random
	if rand.Intn(2) == 0 {
		return -n
	}

	return n
}
