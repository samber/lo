package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialFunc2(t *testing.T) {
	add := MFunc2(func(a, b int) int {
		return a + b
	})
	assert.Equal(t, 3, add.Partial(1).Partial(2)())
	assert.Equal(t, 3, add.PartialR(1).PartialR(2)())
}

func TestPartialFunc10(t *testing.T) {
	type myInt1 int
	type myInt2 int
	type myInt3 int
	type myInt4 int
	type myInt5 int
	type myInt6 int
	type myInt7 int
	type myInt8 int
	type myInt9 int
	type myInt10 int

	add := MFunc10(func(a myInt1, b myInt2, c myInt3, d myInt4, e myInt5, f myInt6, g myInt7, h myInt8, i myInt9, j myInt10) int {
		return int(a) + int(b) + int(c) + int(d) + int(e) + int(f) + int(g) + int(h) + int(i) + int(j)
	})
	assert.Equal(t,
		55,
		add.
			Partial(1).
			Partial(2).
			Partial(3).
			Partial(4).
			Partial(5).
			Partial(6).
			Partial(7).
			Partial(8).
			Partial(9).
			Partial(10)())
	assert.Equal(t,
		55,
		add.
			PartialR(1).
			PartialR(2).
			PartialR(3).
			PartialR(4).
			PartialR(5).
			PartialR(6).
			PartialR(7).
			PartialR(8).
			PartialR(9).
			PartialR(10)())
}
