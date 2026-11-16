//go:build go1.23

package benchmark

import (
	"fmt"
	"testing"

	"github.com/samber/lo/it"
)

func BenchmarkItSum(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Sum(ints)
			}
		})
	}
}

func BenchmarkItSumBy(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.SumBy(ints, func(x int) int { return x })
			}
		})
	}
}

func BenchmarkItProduct(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Product(ints)
			}
		})
	}
}

func BenchmarkItMean(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Mean(ints)
			}
		})
	}
}

func BenchmarkItMode(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Mode(ints)
			}
		})
	}
}

func BenchmarkItRange(b *testing.B) {
	for _, n := range itLengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				for range it.Range(n) { //nolint:revive
				}
			}
		})
	}
}

func BenchmarkItLength(b *testing.B) {
	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for range b.N {
				_ = it.Length(ints)
			}
		})
	}
}
