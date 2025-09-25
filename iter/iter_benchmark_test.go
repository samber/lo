//go:build go1.23

package iter

import (
	"fmt"
	"iter"
	"math/rand"
	"strconv"
	"testing"
)

var lengths = []int{10, 100, 1000}

func BenchmarkChunk(b *testing.B) {
	for _, n := range lengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Chunk(strs, 5)
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Chunk(ints, 5)
			}
		})
	}
}

func genStrings(n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for range n {
			if !yield(strconv.Itoa(rand.Intn(100_000))) {
				break
			}
		}
	}
}

func genInts(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for range n {
			if !yield(rand.Intn(100_000)) {
				break
			}
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	for _, n := range lengths {
		ints := make([]iter.Seq[int], 0, n)
		for i := 0; i < n; i++ {
			ints = append(ints, genInts(n))
		}
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Flatten(ints)
			}
		})
	}

	for _, n := range lengths {
		strs := make([]iter.Seq[string], 0, n)
		for i := 0; i < n; i++ {
			strs = append(strs, genStrings(n))
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
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Drop(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Drop(ints, n/4)
			}
		})
	}
}

func BenchmarkDropWhile(b *testing.B) {
	for _, n := range lengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkDropByIndex(b *testing.B) {
	for _, n := range lengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropByIndex(strs, n/4)
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = DropByIndex(ints, n/4)
			}
		})
	}
}

func BenchmarkReplace(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000}
	for _, n := range lengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Replace(strs, "321321", "123123", 10)
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Replace(ints, 321321, 123123, 10)
			}
		})
	}
}
