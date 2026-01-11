//go:build go1.22

package xrand

import (
	"math/rand/v2"
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

	uSeed := uint64(seed)
	seededRand = rand.New(rand.NewPCG(uSeed, 0))
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
		return seededRand.IntN(n)
	}

	return rand.IntN(n)
}

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int64() int64 {
	if seededRand != nil {
		return seededRand.Int64()
	}

	return rand.Int64()
}
