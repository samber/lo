//go:build go1.23

package it

import "iter"

// ToChannel returns a read-only channel of collection elements.
func ToChannel[T any](bufferSize int, collection iter.Seq[T]) <-chan T {
	ch := make(chan T, bufferSize)

	go func() {
		for item := range collection {
			ch <- item
		}

		close(ch)
	}()

	return ch
}

// FromChannel returns a sequence built from channels items. Blocks until channel closes.
func FromChannel[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range ch {
			if !yield(item) {
				return
			}
		}
	}
}
