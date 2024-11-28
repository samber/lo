package la

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/thoas/go-funk"
	"iter"
	"math/rand"
	"slices"
	"strconv"
	"testing"
	"time"
)

var lengths = []int{10, 100, 1000}

func seqGenerator(size uint) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		for range size {
			if !yield(r.Int63()) {
				return
			}
		}
	}
}

func seq2Generator(size uint) iter.Seq2[int, int64] {
	return Enumerate(seqGenerator(size))
}

// BenchmarkMap benchmark the Map function from the main lo package and from the lazy
// la package to compare their performance.
//
// Since lazy package should be faster when we compose multiple operations — we
// try to simulate that case by layering the same work on top of each other.
func BenchmarkMap(b *testing.B) {
	arr := Collect(seqGenerator(100000), WithSliceCapacity(100000))
	seq := slices.Values(arr)

	for _, workLayers := range []int{0, 5, 25, 50, 100} {
		work := func(value string) string {
			return "mapped " + value
		}

		b.Run(fmt.Sprintf("Map_%d", workLayers), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				intermediate := lo.Map(arr, func(x int64, i int) string {
					return strconv.FormatInt(x, 10)
				})

				for range workLayers {
					intermediate = lo.Map(intermediate, func(value string, _ int) string {
						return "mapped " + value
					})
				}

				_ = intermediate
			}
		})

		b.Run(fmt.Sprintf("lazyMap_%d", workLayers), func(b *testing.B) {
			// This line made it not fair to [lo.Map] function but this line also main
			// difference between lazy and ordinal executions. With lo.Map we can't delay
			// execution until the result will be actually requested, so we should measure
			// the effect of such a delay.

			intermediate := Map(seq, func(x int64) string {
				return strconv.FormatInt(x, 10)
			})

			for range workLayers {
				intermediate = Map(intermediate, work)
			}

			for n := 0; n < b.N; n++ {
				_ = Collect(intermediate, WithSliceCapacity(len(arr)))
			}
		})

		b.Run(fmt.Sprintf("lazyReflect_%d", workLayers), func(b *testing.B) {
			// funk lib also supports lazy initialization of the pipeline and iterates it
			// later but still not supports iterator itself, so we pass arr here.
			intermediate := funk.LazyChain(arr)
			intermediate = intermediate.Map(func(x int64) string {
				return strconv.FormatInt(x, 10)
			})

			for range workLayers {
				intermediate = intermediate.Map(work)
			}

			for n := 0; n < b.N; n++ {
				_ = intermediate.Value().([]string)
			}
		})

		b.Run(fmt.Sprintf("for_%d", workLayers), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				results := make([]string, len(arr))

				for i, item := range arr {
					result := strconv.FormatInt(item, 10)

					// actually, this is different to anything above but do it just for fun.
					for range workLayers {
						result = "mapped " + result
					}

					results[i] = result
				}

				_ = results
			}
		})
	}
}

func BenchmarkChunk(b *testing.B) {
	for _, n := range lengths {
		strs := slices.Values(genSliceString(n))
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Collect(Chunk(strs, 5), WithSliceCapacity(n))
			}
		})
	}

	for _, n := range lengths {
		ints := slices.Values(genSliceInt(n))
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Collect(Chunk(ints, 5), WithSliceCapacity(n))
			}
		})
	}
}

func genSliceString(n int) []string {
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, strconv.Itoa(rand.Intn(100_000)))
	}
	return res
}

func genSliceInt(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(100_000))
	}
	return res
}

func BenchmarkFlatten(b *testing.B) {
	for _, n := range lengths {
		ints := make([][]int, 0, n)
		for i := 0; i < n; i++ {
			ints = append(ints, genSliceInt(n))
		}

		intsGen := slices.Values(ints)

		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = FlattenSlice(intsGen)
			}
		})
	}

	for _, n := range lengths {
		strs := make([][]string, 0, n)
		for i := 0; i < n; i++ {
			strs = append(strs, genSliceString(n))
		}

		strsGen := slices.Values(strs)

		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = FlattenSlice(strsGen)
			}
		})
	}
}

func BenchmarkReplace(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000}
	for _, n := range lengths {
		strs := genSliceString(n)
		strsGen := slices.Values(strs)

		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Replace(strsGen, strs[n/4], "123123", 10)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		intsGen := slices.Values(ints)

		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Replace(intsGen, ints[n/4], 123123, 10)
			}
		})
	}
}
