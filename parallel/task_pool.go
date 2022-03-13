package parallel

import (
	"math"
	"sync"
)

type job[T any] struct {
	val T
	i   int
}

func NewTaskPool[T any](source []T, poolSize int, fn func(T, int)) {
	cycles := int(math.Ceil(float64(len(source)) / float64(poolSize)))
	n := len(source)
	c := 0

	wg := &sync.WaitGroup{}
	wg.Add(poolSize)

	jobChans := make([]chan job[T], poolSize)
	for i := 0; i < len(jobChans); i++ {
		jobChans[i] = make(chan job[T], cycles)
		go func(i int, in <-chan job[T]) {
			for {
				v, ok := <-in
				if !ok {
					wg.Done()
					return
				}
				fn(v.val, v.i)
			}
		}(i, jobChans[i])
	}

	for i := 0; i < cycles; i++ {
		n -= poolSize
		batchSize := 0
		if n < 0 {
			batchSize = n + poolSize
		} else {
			batchSize = poolSize
		}

		for j := 0; j < batchSize; c, j = c+1, j+1 {
			jobChans[j] <- job[T]{
				val: source[c],
				i:   c,
			}
		}
	}

	for i := 0; i < len(jobChans); i++ {
		close(jobChans[i])
	}

	wg.Wait()
}
