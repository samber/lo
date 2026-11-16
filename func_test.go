package lo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartial(t *testing.T) {
	t.Parallel()
	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial(add, 5)

	tests := []struct {
		name     string
		y        int
		expected string
	}{
		{name: "positive", y: 10, expected: "15"},
		{name: "negative", y: -5, expected: "0"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, f(tt.y))
		})
	}
}

func TestPartial1(t *testing.T) {
	t.Parallel()
	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial1(add, 5)

	tests := []struct {
		name     string
		y        int
		expected string
	}{
		{name: "positive", y: 10, expected: "15"},
		{name: "negative", y: -5, expected: "0"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, f(tt.y))
		})
	}
}

func TestPartial2(t *testing.T) {
	t.Parallel()
	add := func(x float64, y, z int) string {
		return strconv.Itoa(int(x) + y + z)
	}
	f := Partial2(add, 5)

	tests := []struct {
		name     string
		y        int
		z        int
		expected string
	}{
		{name: "positive", y: 10, z: 9, expected: "24"},
		{name: "negative", y: -5, z: 8, expected: "8"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, f(tt.y, tt.z))
		})
	}
}

func TestPartial3(t *testing.T) {
	t.Parallel()
	add := func(x float64, y, z int, a float32) string {
		return strconv.Itoa(int(x) + y + z + int(a))
	}
	f := Partial3(add, 5)

	tests := []struct {
		name     string
		y        int
		z        int
		a        float32
		expected string
	}{
		{name: "positive", y: 10, z: 9, a: -3, expected: "21"},
		{name: "negative", y: -5, z: 8, a: 7, expected: "15"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, f(tt.y, tt.z, tt.a))
		})
	}
}

func TestPartial4(t *testing.T) {
	t.Parallel()
	add := func(x float64, y, z int, a float32, b int32) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b))
	}
	f := Partial4(add, 5)

	tests := []struct {
		name     string
		y        int
		z        int
		a        float32
		b        int32
		expected string
	}{
		{name: "positive", y: 10, z: 9, a: -3, b: 0, expected: "21"},
		{name: "negative", y: -5, z: 8, a: 7, b: -1, expected: "14"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, f(tt.y, tt.z, tt.a, tt.b))
		})
	}
}

func TestPartial5(t *testing.T) {
	t.Parallel()
	add := func(x float64, y, z int, a float32, b int32, c int) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b) + c)
	}
	f := Partial5(add, 5)

	tests := []struct {
		name     string
		y        int
		z        int
		a        float32
		b        int32
		c        int
		expected string
	}{
		{name: "positive", y: 10, z: 9, a: -3, b: 0, c: 5, expected: "26"},
		{name: "negative", y: -5, z: 8, a: 7, b: -1, c: 7, expected: "21"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, f(tt.y, tt.z, tt.a, tt.b, tt.c))
		})
	}
}
