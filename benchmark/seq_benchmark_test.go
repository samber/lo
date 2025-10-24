//go:build go1.23

package benchmark

import (
	"fmt"
	"iter"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/samber/lo/it"
)

var itLengths = []int{10, 100, 1000}

func BenchmarkItChunk(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Chunk(strs, 5)
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Chunk(ints, 5)
			}
		})
	}
}

func genStrings(n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for range n {
			if !yield(strconv.Itoa(rand.IntN(100_000))) {
				break
			}
		}
	}
}

func genInts(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for range n {
			if !yield(rand.IntN(100_000)) {
				break
			}
		}
	}
}

func BenchmarkItFlatten(b *testing.B) {
	for _, n := range itLengths {
		ints := make([]iter.Seq[int], 0, n)
		for i := 0; i < n; i++ {
			ints = append(ints, genInts(n))
		}
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Flatten(ints)
			}
		})
	}

	for _, n := range itLengths {
		strs := make([]iter.Seq[string], 0, n)
		for i := 0; i < n; i++ {
			strs = append(strs, genStrings(n))
		}
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Flatten(strs)
			}
		})
	}
}

func BenchmarkItDrop(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Drop(strs, n/4)
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Drop(ints, n/4)
			}
		})
	}
}

func BenchmarkItDropWhile(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.DropWhile(strs, func(v string) bool { return len(v) < 4 })
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.DropWhile(ints, func(v int) bool { return i < 10_000 })
			}
		})
	}
}

func BenchmarkItDropByIndex(b *testing.B) {
	for _, n := range itLengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.DropByIndex(strs, n/4)
			}
		})
	}

	for _, n := range itLengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.DropByIndex(ints, n/4)
			}
		})
	}
}

func BenchmarkItReplace(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000}

	for _, n := range lengths {
		strs := genStrings(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Replace(strs, "321321", "123123", 10)
			}
		})
	}

	for _, n := range lengths {
		ints := genInts(n)
		b.Run(fmt.Sprintf("ints%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = it.Replace(ints, 321321, 123123, 10)
			}
		})
	}
}
