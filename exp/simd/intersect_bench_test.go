//go:build goexperiment.simd

package simd

import (
	"testing"

	"github.com/samber/lo"
)

// benchContains sweeps a hit-position axis (miss, hit-first, hit-mid, hit-last) at each
// size: this is what tunes containsBlock, the number of vector compares accumulated
// between mask reductions. A larger containsBlock amortizes reduction overhead on
// miss/late-hit workloads but wastes more work on early hits.
func benchContains[T numeric](b *testing.B, wrapper, scalar func([]T, T) bool) {
	b.Helper()
	for _, bs := range benchmarkSizes {
		data := makeData[T](bs.size, 101)
		positions := []struct {
			name   string
			target T
		}{
			{"miss", pickAbsent(data)},
			{"hit-first", data[0]},
			{"hit-mid", data[len(data)/2]},
			{"hit-last", data[len(data)-1]},
		}
		b.Run(bs.name, func(b *testing.B) {
			for _, p := range positions {
				target := p.target
				b.Run(p.name+"/lo", func(b *testing.B) {
					b.ReportAllocs()
					for range b.N {
						_ = scalar(data, target)
					}
				})
				b.Run(p.name+"/simd", func(b *testing.B) {
					b.ReportAllocs()
					for range b.N {
						_ = wrapper(data, target)
					}
				})
			}
		})
	}
}

func BenchmarkContainsInt8(b *testing.B) {
	benchContains(b, ContainsInt8[int8], lo.Contains[int8])
}

func BenchmarkContainsInt32(b *testing.B) {
	benchContains(b, ContainsInt32[int32], lo.Contains[int32])
}

func BenchmarkContainsInt64(b *testing.B) {
	benchContains(b, ContainsInt64[int64], lo.Contains[int64])
}

func BenchmarkContainsFloat64(b *testing.B) {
	benchContains(b, ContainsFloat64[float64], lo.Contains[float64])
}
