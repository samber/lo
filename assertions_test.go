package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.NotPanics(func() {
		Assert(true)
	})

	is.NotPanics(func() {
		Assert(true, "user defined message")
	})

	is.PanicsWithValue("assertion failed", func() {
		Assert(false)
	})

	is.PanicsWithValue("assertion failed: user defined message", func() {
		Assert(false, "user defined message")
	})

	//checks that the examples in `README.md` compile
	{
		age := 20
		is.NotPanics(func() {
			Assert(age >= 15)
		})
		is.NotPanics(func() {
			Assert(age >= 15, "user age must be >= 15")
		})
	}
}

func TestAssertf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.NotPanics(func() {
		Assertf(true, "user defined message")
	})

	is.NotPanics(func() {
		Assertf(true, "user defined message %d %d", 1, 2)
	})

	is.PanicsWithValue("assertion failed: user defined message", func() {
		Assertf(false, "user defined message")
	})

	is.PanicsWithValue("assertion failed: user defined message 1 2", func() {
		Assertf(false, "user defined message %d %d", 1, 2)
	})

	//checks that the example in `README.md` compiles
	{
		age := 7
		is.PanicsWithValue("assertion failed: user age must be >= 15, got 7", func() {
			Assertf(age >= 15, "user age must be >= 15, got %d", age)
		})
	}
}
