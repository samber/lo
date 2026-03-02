//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo"
)

// SeqToChannel returns a read-only channel of collection elements.
// Play: https://go.dev/play/p/id3jqJPffT6
func SeqToChannel[T any](bufferSize int, collection iter.Seq[T]) <-chan T {
	ch := make(chan T, bufferSize)

	go func() {
		for item := range collection {
			ch <- item
		}

		close(ch)
	}()

	return ch
}

// SeqToChannel2 returns a read-only channel of collection elements.
// Play: https://go.dev/play/p/rpJdVnXUaG-
func SeqToChannel2[K, V any](bufferSize int, collection iter.Seq2[K, V]) <-chan lo.Tuple2[K, V] {
	ch := make(chan lo.Tuple2[K, V], bufferSize)

	go func() {
		for k, v := range collection {
			ch <- lo.Tuple2[K, V]{A: k, B: v}
		}

		close(ch)
	}()

	return ch
}

// ChannelToSeq returns a sequence built from channels items. Blocks until channel closes.
// Play: https://go.dev/play/p/IXqSs2Ooqpm
func ChannelToSeq[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range ch {
			if !yield(item) {
				return
			}
		}
	}
}

// Buffer returns a sequence of slices, each containing up to size items read from the channel.
// The last slice may be smaller if the channel closes before filling the buffer.
func Buffer[T any](ch <-chan T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for {
			buffer := make([]T, 0, size)

			for range size {
				item, ok := <-ch
				if !ok {
					if len(buffer) > 0 {
						yield(buffer)
					}
					return
				}
				buffer = append(buffer, item)
			}

			if !yield(buffer) {
				return
			}
		}
	}
}
