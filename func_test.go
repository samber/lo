package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	is := assert.New(t)

	add := func(x, y int) int { return x + y }
	f := Bind(add, 5)
	is.Equal(15, f(10))
	is.Equal(0, f(-5))
}
