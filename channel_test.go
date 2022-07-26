package lo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToChannel(t *testing.T) {
	is := assert.New(t)

	ch := ToChannel[int]([]int{1, 2, 3})

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

func TestGenerate(t *testing.T) {
	// is := assert.New(t)

	// next := func(i int64) int64 {
	// 	return i
	// }

	// for item := range Generator[int64](10, next) {
	// 	println(item)
	// }
}

func TestBatch(t *testing.T) {
	is := assert.New(t)

	ch := ToChannel[int]([]int{1, 2, 3})

	items1, length1 := Batch[int](ch, 2)
	items2, length2 := Batch[int](ch, 2)
	items3, length3 := Batch[int](ch, 2)

	is.Equal([]int{1, 2}, items1)
	is.Equal(2, length1)
	is.Equal([]int{3}, items2)
	is.Equal(1, length2)
	is.Equal([]int{}, items3)
	is.Equal(0, length3)
}

func TestBatchWithTimeout(t *testing.T) {
	is := assert.New(t)

	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
	}()

	items1, length1 := BatchWithTimeout[int](ch, 20, 15*time.Millisecond)
	is.Equal([]int{0, 1}, items1)
	is.Equal(2, length1)

	items2, length2 := BatchWithTimeout[int](ch, 20, 2*time.Millisecond)
	is.Equal([]int{}, items2)
	is.Equal(0, length2)

	items3, length3 := BatchWithTimeout[int](ch, 1, 30*time.Millisecond)
	is.Equal([]int{2}, items3)
	is.Equal(1, length3)

	items4, length4 := BatchWithTimeout[int](ch, 2, 25*time.Millisecond)
	is.Equal([]int{3, 4}, items4)
	is.Equal(2, length4)

	items5, length5 := BatchWithTimeout[int](ch, 3, 25*time.Millisecond)
	is.Equal([]int{}, items5)
	is.Equal(0, length5)
}
