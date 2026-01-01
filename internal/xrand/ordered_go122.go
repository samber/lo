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

	// Convert int64 seed to [32]byte for ChaCha8
	// We use the seed to initialize both halves of the key
	var seedBytes [32]byte
	for i := 0; i < 8; i++ {
		seedBytes[i] = byte(seed >> (i * 8))
		seedBytes[i+8] = byte(seed >> (i * 8))
		seedBytes[i+16] = byte(seed >> (i * 8))
		seedBytes[i+24] = byte(seed >> (i * 8))
	}

	seededRand = rand.New(rand.NewChaCha8(seedBytes))
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
		return r.IntN(n)
	}

	return rand.IntN(n)
}

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func Int64() int64 {
	mu.Lock()
	r := seededRand
	mu.Unlock()

	if r != nil {
		return r.Int64()
	}

	return rand.Int64()
}
