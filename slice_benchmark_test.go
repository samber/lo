package lo

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var lengths = []int{10, 100, 1000}

func BenchmarkChunk(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Chunk(strs, 5)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Chunk(ints, 5)
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
				_ = Flatten(ints)
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
				_ = Flatten(strs)
			}
		})
	}
}

func BenchmarkDrop(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Drop(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Drop(ints, n/4)
			}
		})
	}
}

func BenchmarkDropRight(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropRight(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropRight(ints, n/4)
			}
		})
	}
}

func BenchmarkDropWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkDropRightWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropRightWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropRightWhile(ints, func(v int) bool { return i < 10_000 })
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
				_ = Replace(strs, strs[n/4], "123123", 10)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Replace(ints, ints[n/4], 123123, 10)
			}
		})
	}
}
