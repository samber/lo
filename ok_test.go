package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustOk(t *testing.T) {
	is := assert.New(t)
	is.Equal("foo", MustOk("foo", true))
	is.Panics(
		func() {
			MustOk("", false)
		},
	)
}

func TestMustOkX(t *testing.T) {
	is := assert.New(t)

	{
		is.Panics(func() {
			MustOk0(false)
		})
		is.NotPanics(func() {
			MustOk0(true)
		})
	}

	{
		val1 := MustOk1(1, true)
		is.Equal(1, val1)
		is.Panics(func() {
			MustOk1(1, false)
		})
	}

	{
		val1, val2 := MustOk2(1, 2, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Panics(func() {
			MustOk2(1, 2, false)
		})
	}

	{
		val1, val2, val3 := MustOk3(1, 2, 3, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Panics(func() {
			MustOk3(1, 2, 3, false)
		})
	}

	{
		val1, val2, val3, val4 := MustOk4(1, 2, 3, 4, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Panics(func() {
			MustOk4(1, 2, 3, 4, false)
		})
	}

	{
		val1, val2, val3, val4, val5 := MustOk5(1, 2, 3, 4, 5, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.Panics(func() {
			MustOk5(1, 2, 3, 4, 5, false)
		})
	}

	{
		val1, val2, val3, val4, val5, val6 := MustOk6(1, 2, 3, 4, 5, 6, true)
		is.Equal(1, val1)
		is.Equal(2, val2)
		is.Equal(3, val3)
		is.Equal(4, val4)
		is.Equal(5, val5)
		is.Equal(6, val6)
		is.Panics(func() {
			MustOk6(1, 2, 3, 4, 5, 6, false)
		})
	}
}
