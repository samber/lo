package lo

import "github.com/samber/lo/internal/xrand"

// SetRandomSeed sets a custom seed for the random number generator used by
// RandomString, Shuffle, and other functions that rely on random numbers.
// This allows for reproducible random sequences, which is useful for testing.
//
// Pass a negative value to reset to the default (non-reproducible) behavior.
//
// Example:
//
//	lo.SetRandomSeed(42)
//	s1 := lo.RandomString(10, lo.AlphanumericCharset) // Always same result with seed 42
//
//	lo.SetRandomSeed(42)
//	s2 := lo.RandomString(10, lo.AlphanumericCharset) // s1 == s2
//
//	lo.ResetRandomSeed() // Back to default random behavior
//
// Note: This function is NOT safe to call concurrently with functions that use
// the random number generator. It is intended to be called once at the start
// of a test or program.
func SetRandomSeed(seed int64) {
	xrand.SetSeed(seed)
}

// ResetRandomSeed resets the random number generator to its default
// (non-reproducible) behavior. This is equivalent to calling SetRandomSeed(-1).
//
// Example:
//
//	lo.SetRandomSeed(42)
//	// ... do some reproducible random operations ...
//	lo.ResetRandomSeed() // Back to default random behavior
func ResetRandomSeed() {
	xrand.ResetSeed()
}
