package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexesMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var items = []string{"a", "b", "c"}

	indexMap := IndexMap(items)
	is.Len(indexMap, len(items))
	is.Equal(0, indexMap["a"])
	is.Equal(1, indexMap["b"])
	is.Equal(2, indexMap["c"])
}

func TestSortAs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	var items = []Tuple2[uint, string]{
		{A: 1, B: "hello"},
		{A: 2, B: "world"},
		{A: 3, B: "happy"},
		{A: 4, B: "new year"},
	}
	var keys = []uint{3, 4, 1, 2}

	SortAs(items, keys, func(t Tuple2[uint, string]) uint {
		return t.A
	})

	for i := range keys {
		is.Equal(items[i].A, keys[i])
	}
	is.Equal(items[0].B, "happy")
	is.Equal(items[1].B, "new year")
	is.Equal(items[2].B, "hello")
	is.Equal(items[3].B, "world")
}
