//go:build go1.23

package iter

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := ToChannel(2, values(1, 2, 3))

	r1, ok1 := <-ch
	r2, ok2 := <-ch
	r3, ok3 := <-ch
	is.True(ok1)
	is.Equal(1, r1)
	is.True(ok2)
	is.Equal(2, r2)
	is.True(ok3)
	is.Equal(3, r3)

	_, ok4 := <-ch
	is.False(ok4)
}

func TestFromChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := ToChannel(2, values(1, 2, 3))
	items := FromChannel(ch)

	is.Equal([]int{1, 2, 3}, slices.Collect(items))
}
