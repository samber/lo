package lo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartial(t *testing.T) {
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}
