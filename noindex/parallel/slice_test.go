package parallel

import (
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	is := assert.New(t)

	result1 := Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	})
	result2 := Map([]int64{1, 2, 3, 4}, func(x int64) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "3", "4"})
}

func TestForEach(t *testing.T) {
	is := assert.New(t)

	var counter uint64
	collection := []int{1, 2, 3, 4}
	ForEach(collection, func(x int) {
		atomic.AddUint64(&counter, 1)
	})

	is.Equal(uint64(4), atomic.LoadUint64(&counter))
}
