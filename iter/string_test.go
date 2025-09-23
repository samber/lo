package iter

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ChunkString("12345", 2)
	is.Equal([]string{"12", "34", "5"}, slices.Collect(result1))

	result2 := ChunkString("123456", 2)
	is.Equal([]string{"12", "34", "56"}, slices.Collect(result2))

	result3 := ChunkString("123456", 6)
	is.Equal([]string{"123456"}, slices.Collect(result3))

	result4 := ChunkString("123456", 10)
	is.Equal([]string{"123456"}, slices.Collect(result4))

	result5 := ChunkString("", 2)
	is.Equal([]string{""}, slices.Collect(result5))

	result6 := ChunkString("明1好休2林森", 2)
	is.Equal([]string{"明1", "好休", "2林", "森"}, slices.Collect(result6))

	is.PanicsWithValue("iter.ChunkString: size must be greater than 0", func() {
		ChunkString("12345", 0)
	})
}
