package benchmark

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
)

var lengths = []int{10, 100, 1000}

func BenchmarkChunk(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Chunk(strs, 5)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Chunk(ints, 5)
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
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Flatten(ints)
			}
		})
	}

	for _, n := range lengths {
		strs := make([][]string, 0, n)
		for i := 0; i < n; i++ {
			strs = append(strs, genSliceString(n))
		}
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Flatten(strs)
			}
		})
	}
}

func BenchmarkDrop(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Drop(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Drop(ints, n/4)
			}
		})
	}
}

func BenchmarkDropRight(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRight(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRight(ints, n/4)
			}
		})
	}
}

func BenchmarkDropWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkDropRightWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRightWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropRightWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkDropByIndex(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropByIndex(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.DropByIndex(ints, n/4)
			}
		})
	}
}

func BenchmarkReplace(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000}
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Replace(strs, strs[n/4], "123123", 10)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Replace(ints, ints[n/4], 123123, 10)
			}
		})
	}
}

func BenchmarkToSlicePtr(b *testing.B) {
	preallocated := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		_ = lo.ToSlicePtr(preallocated)
	}
}

func BenchmarkShuffle(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			b.ReportAllocs() // This reports memory allocations
			for i := 0; i < b.N; i++ {
				mutable.Shuffle(ints)
			}
		})
	}

	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				mutable.Shuffle(strs)
			}
		})
	}
}
