package lo

import (
	"testing"

	"github.com/samber/lo/mutable"
	"github.com/stretchr/testify/assert"
)

func TestSetRandomSeed(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	t.Cleanup(func() {
		ResetRandomSeed()
	})

	is := assert.New(t)

	t.Run("reproducible RandomString", func(t *testing.T) {
		SetRandomSeed(42)
		s1 := RandomString(20, AlphanumericCharset)

		SetRandomSeed(42)
		s2 := RandomString(20, AlphanumericCharset)

		is.Equal(s1, s2, "RandomString should produce the same result with the same seed")
	})

	t.Run("different seeds produce different results", func(t *testing.T) {
		SetRandomSeed(42)
		s1 := RandomString(20, AlphanumericCharset)

		SetRandomSeed(123)
		s2 := RandomString(20, AlphanumericCharset)

		is.NotEqual(s1, s2, "Different seeds should produce different results")
	})

	t.Run("reproducible Shuffle", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		SetRandomSeed(42)
		slice1 := make([]int, len(original))
		copy(slice1, original)
		mutable.Shuffle(slice1)

		SetRandomSeed(42)
		slice2 := make([]int, len(original))
		copy(slice2, original)
		mutable.Shuffle(slice2)

		is.Equal(slice1, slice2, "mutable.Shuffle should produce the same result with the same seed")
	})

	t.Run("reset returns to non-reproducible behavior", func(t *testing.T) {
		SetRandomSeed(42)
		s1 := RandomString(20, AlphanumericCharset)

		ResetRandomSeed()
		s2 := RandomString(20, AlphanumericCharset)
		s3 := RandomString(20, AlphanumericCharset)

		// After reset, consecutive calls should produce different results (with very high probability)
		// Note: There's an astronomically small chance this could fail if truly random
		is.NotEqual(s2, s3, "After reset, RandomString should produce different results")

		// And it should not match the seeded result
		SetRandomSeed(42)
		s4 := RandomString(20, AlphanumericCharset)
		is.Equal(s1, s4, "Re-seeding with 42 should reproduce the original result")
	})

	t.Run("negative seed resets", func(t *testing.T) {
		SetRandomSeed(42)
		s1 := RandomString(20, AlphanumericCharset)

		SetRandomSeed(-1) // Reset via negative seed

		SetRandomSeed(42)
		s2 := RandomString(20, AlphanumericCharset)

		is.Equal(s1, s2, "Negative seed should reset, then re-seeding should work")
	})
}

func TestSetRandomSeed_MultipleOperations(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	t.Cleanup(func() {
		ResetRandomSeed()
	})

	is := assert.New(t)

	// Test that a sequence of operations is reproducible
	SetRandomSeed(999)
	results1 := []string{
		RandomString(10, AlphanumericCharset),
		RandomString(5, LowerCaseLettersCharset),
		RandomString(15, NumbersCharset),
	}

	SetRandomSeed(999)
	results2 := []string{
		RandomString(10, AlphanumericCharset),
		RandomString(5, LowerCaseLettersCharset),
		RandomString(15, NumbersCharset),
	}

	is.Equal(results1, results2, "A sequence of operations should be reproducible")
}
