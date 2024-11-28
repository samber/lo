package la

import (
	"github.com/stretchr/testify/assert"
	"maps"
	"math"
	"slices"
	"strconv"
	"testing"
)

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	}))

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestTimes2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := maps.Collect(Times2(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	}))

	is.Equal(len(result1), 3)
	is.Equal(result1, map[int]string{0: "0", 1: "1", 2: "2"})
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Repeat(2, foo{"a"}))
	result2 := slices.Collect(Repeat(0, foo{"a"}))

	is.Equal(result1, []foo{{"a"}, {"a"}})
	is.Equal(result2, ([]foo)(nil))
}

func TestRepeatBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cb := func(i int) int {
		return int(math.Pow(float64(i), 2))
	}

	result1 := slices.Collect(RepeatBy(0, cb))
	result2 := slices.Collect(RepeatBy(2, cb))
	result3 := slices.Collect(RepeatBy(5, cb))

	is.Equal(([]int)(nil), result1)
	is.Equal([]int{0, 1}, result2)
	is.Equal([]int{0, 1, 4, 9, 16}, result3)
}
