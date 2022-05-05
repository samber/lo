package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	is := assert.New(t)

	cb := func(x int) int { return x * x * x }
	tp := func(x int) int { return 3 * x }
	db := func(x int) int { return 2 * x }
	f := Pipeline(cb, tp, db)
	is.Equal(750, f(5))
}
