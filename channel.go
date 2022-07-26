package lo

import (
	"time"
)

// ToChannel returns a read-only channels of collection elements.
func ToChannel[T any](collection []T) <-chan T {
	ch := make(chan T)

	go func() {
		for _, item := range collection {
			ch <- item
		}

		close(ch)
	}()

	return ch
}

// Generator implements the generator design pattern.
func Generator[T any](bufferSize int, generator func(int64) T) <-chan T {
	ch := make(chan T, bufferSize)

	go func() {
		var i int64 = 0

		// WARNING: infinite loop
		for {
			ch <- generator(i)
			i++
		}

		close(ch)
	}()

	return ch
}

// Batch creates a slice of n elements from a channel. Returns the slice and the slice length.
func Batch[T any](ch <-chan T, size int) (collection []T, length int) {
	buffer := make([]T, 0, size)
	index := 0

	for ; index < size; index++ {
		select {
		case item, ok := <-ch:
			if !ok {
				return buffer, index
			}

			buffer = append(buffer, item)
		}
	}

	return buffer, index
}

// BatchWithTimeout creates a slice of n elements from a channel, with timeout. Returns the slice and the slice length.
func BatchWithTimeout[T any](ch <-chan T, size int, timeout time.Duration) (collection []T, length int) {
	expire := time.After(timeout)

	buffer := make([]T, 0, size)
	index := 0

	for ; index < size; index++ {
		select {
		case item, ok := <-ch:
			if !ok {
				return buffer, index
			}

			buffer = append(buffer, item)

		case <-expire:
			return buffer, index
		}
	}

	return buffer, index
}
