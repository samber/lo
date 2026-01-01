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
	mu.Lock()
	r := seededRand
	mu.Unlock()

	if r != nil {
		r.Shuffle(n, swap)
		return
	}

	rand.Shuffle(n, swap)
}

// IntN returns, as an int, a pseudo-random number in the half-open interval [0,n)
// from the default Source.
// It panics if n <= 0.
func IntN(n int) int {
	mu.Lock()
	r := seededRand
	mu.Unlock()

	if r != nil {
		return r.Intn(n)
	}

	// bearer:disable go_gosec_crypto_weak_random
	return rand.Intn(n)
}

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int64() int64 {
	mu.Lock()
	r := seededRand
	mu.Unlock()

	if r != nil {
		n := r.Int63()
		if r.Intn(2) == 0 {
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
