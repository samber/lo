//go:build go1.23

package it

import (
	"maps"
	"slices"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSeqToChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := SeqToChannel(2, values(1, 2, 3))

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

func TestSeqToChannel2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := SeqToChannel2(2, maps.All(map[string]int{"a": 1, "b": 2, "c": 3}))

	r1, ok1 := <-ch
	r2, ok2 := <-ch
	r3, ok3 := <-ch
	is.True(ok1)
	is.True(ok2)
	is.True(ok3)
	is.ElementsMatch([]lo.Tuple2[string, int]{{A: "a", B: 1}, {A: "b", B: 2}, {A: "c", B: 3}}, []lo.Tuple2[string, int]{r1, r2, r3})

	_, ok4 := <-ch
	is.False(ok4)
}

func TestChannelToSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := SeqToChannel(2, values(1, 2, 3))
	items := ChannelToSeq(ch)

	is.Equal([]int{1, 2, 3}, slices.Collect(items))
}
