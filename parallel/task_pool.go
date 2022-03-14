package parallel

import (
	"math"
	"sync"
)

//TaskPool creates a pool of workers, working through n amount of jobs.
//The callback function receives the current job that is being worked on.
func TaskPool[T any](amount, poolSize int, fn func(int)) {
	cycles := int(math.Ceil(float64(amount) / float64(poolSize)))
	n := amount
	c := 0

	if poolSize > n {
		poolSize = n
	}

	wg := &sync.WaitGroup{}
	wg.Add(poolSize)

	jobChans := make([]chan int, poolSize)
	for i := 0; i < len(jobChans); i++ {
		jobChans[i] = make(chan int, cycles)
		go func(i int, in <-chan int) {
			for {
				v, ok := <-in
				if !ok {
					wg.Done()
					return
				}
				fn(v)
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
			jobChans[j] <- c
		}
	}

	for i := 0; i < len(jobChans); i++ {
		close(jobChans[i])
	}

	wg.Wait()
}
