package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnzipBy789(t *testing.T) {
	t.Parallel()
	t.Run("UnzipBy7", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		in := []int{1, 2}
		a, b, c, d, e, f, g := UnzipBy7(in, func(v int) (int, int, int, int, int, int, int) {
			return v, v * 2, v * 3, v * 4, v * 5, v * 6, v * 7
		})
		is.Equal([]int{1, 2}, a)
		is.Equal([]int{2, 4}, b)
		is.Equal([]int{3, 6}, c)
		is.Equal([]int{4, 8}, d)
		is.Equal([]int{5, 10}, e)
		is.Equal([]int{6, 12}, f)
		is.Equal([]int{7, 14}, g)
	})

	t.Run("UnzipBy8", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		in := []int{1, 2}
		a, b, c, d, e, f, g, h := UnzipBy8(in, func(v int) (int, int, int, int, int, int, int, int) {
			return v, v * 2, v * 3, v * 4, v * 5, v * 6, v * 7, v * 8
		})
		is.Equal([]int{1, 2}, a)
		is.Equal([]int{2, 4}, b)
		is.Equal([]int{3, 6}, c)
		is.Equal([]int{4, 8}, d)
		is.Equal([]int{5, 10}, e)
		is.Equal([]int{6, 12}, f)
		is.Equal([]int{7, 14}, g)
		is.Equal([]int{8, 16}, h)
	})

	t.Run("UnzipBy9", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		in := []int{1, 2}
		a, b, c, d, e, f, g, h, i := UnzipBy9(in, func(v int) (int, int, int, int, int, int, int, int, int) {
			return v, v * 2, v * 3, v * 4, v * 5, v * 6, v * 7, v * 8, v * 9
		})
		is.Equal([]int{1, 2}, a)
		is.Equal([]int{2, 4}, b)
		is.Equal([]int{3, 6}, c)
		is.Equal([]int{4, 8}, d)
		is.Equal([]int{5, 10}, e)
		is.Equal([]int{6, 12}, f)
		is.Equal([]int{7, 14}, g)
		is.Equal([]int{8, 16}, h)
		is.Equal([]int{9, 18}, i)
	})
}
