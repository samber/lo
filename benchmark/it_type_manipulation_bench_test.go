//go:build go1.23

package benchmark

import (
	"fmt"
	"testing"

	"github.com/samber/lo/it"
)

func BenchmarkItToSeqPtr(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ToSeqPtr(ints) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItFromSeqPtr(b *testing.B) {
	for _, n := range itLengths {
		ptrs := genIntPtrSeq(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.FromSeqPtr(ptrs) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItToAnySeq(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.ToAnySeq(ints) { //nolint:revive
				}
			}
		})
	}
}
