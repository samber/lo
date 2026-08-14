package loslice_test

import (
	"github.com/samber/lo/loslice"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const seed = 42

func TestUniqInts(t *testing.T) {
	t.Parallel()

	rnd := rand.New(rand.NewSource(seed))

	xs := make([]int, 100)
	for i := range xs {
		xs[i] = rnd.Intn(4) // Random integers between 0 and 9
	}

	uniq := loslice.Uniq(xs)
	assert.Len(t, uniq, 4)

	for _, x := range []int{0, 1, 2, 3} {
		assert.Contains(t, uniq, x)
	}
}
