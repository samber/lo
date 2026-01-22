//go:build !go1.22

package xrand

import (
	"math/rand"
	"sync"
)

var (
	mu         sync.Mutex
	seededRand *rand.Rand
)

// SetSeed sets a custom seed for the random number generator.
// This allows for reproducible random sequences, which is useful for testing.
// Pass a negative value to reset to the default (non-reproducible) behavior.
func SetSeed(seed int64) {
	mu.Lock()
	defer mu.Unlock()

	if seed < 0 {
		seededRand = nil
		return
	}

	//nolint:staticcheck // rand.NewSource is fine here for backward compatibility
	seededRand = rand.New(rand.NewSource(seed))
}

// ResetSeed resets the random number generator to its default (non-reproducible) behavior.
func ResetSeed() {
	SetSeed(-1)
}

// Shuffle returns a slice of shuffled values. Uses the Fisher-Yates shuffle algorithm.
func Shuffle(n int, swap func(i, j int)) {
	if seededRand != nil {
		seededRand.Shuffle(n, swap)
		return
	}

	rand.Shuffle(n, swap)
}

// IntN returns, as an int, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func IntN(n int) int {
	if seededRand != nil {
		return seededRand.Intn(n)
	}

	// bearer:disable go_gosec_crypto_weak_random
	return rand.Intn(n)
}

// Int64 returns a pseudo-random 63-bit integer as an int64
// from the default Source.
// For Go < 1.22, this simulates the full int64 range by randomly
// negating the result of Int63().
func Int64() int64 {
	if seededRand != nil {
		n := seededRand.Int63()
		if seededRand.Intn(2) == 0 {
			return -n
		}
		return n
	}

	// bearer:disable go_gosec_crypto_weak_random
	n := rand.Int63()

	// bearer:disable go_gosec_crypto_weak_random
	if rand.Intn(2) == 0 {
		return -n
	}

	return n
}
