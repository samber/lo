package benchmark

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
)

func BenchmarkContains(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Contains(ints, target)
			}
		})
	}
}

func BenchmarkContainsBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ContainsBy(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkEvery(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		subset := ints[:n/2]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Every(ints, subset)
			}
		})
	}
}

func BenchmarkEveryBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.EveryBy(ints, func(v int) bool { return v >= 0 })
			}
		})
	}
}

func BenchmarkSome(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		subset := []int{ints[n-1]}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Some(ints, subset)
			}
		})
	}
}

func BenchmarkSomeBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.SomeBy(ints, func(v int) bool { return v < 0 })
			}
		})
	}
}

func BenchmarkNone(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		subset := []int{-1, -2, -3}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.None(ints, subset)
			}
		})
	}
}

func BenchmarkNoneBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.NoneBy(ints, func(v int) bool { return v < 0 })
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Intersect(a, c)
			}
		})
	}
}

func BenchmarkIntersectBy(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.IntersectBy(func(v int) int { return v }, a, c)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		c := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Union(a, c)
			}
		})
	}
}

func BenchmarkWithout(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Without(ints, 1, 2, 3, 4, 5)
			}
		})
	}
}

func BenchmarkWithoutBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.WithoutBy(ints, func(v int) int { return v % 100 }, 1, 2, 3, 4, 5)
			}
		})
	}
}

func BenchmarkWithoutEmpty(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		// sprinkle some zeroes
		for j := 0; j < n/10; j++ {
			ints[j*10] = 0
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.WithoutEmpty(ints) //nolint:staticcheck
			}
		})
	}
}

func BenchmarkWithoutNth(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.WithoutNth(ints, 0, n/2, n-1)
			}
		})
	}
}

func BenchmarkElementsMatch(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		c := make([]int, n)
		copy(c, a)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ElementsMatch(a, c)
			}
		})
	}
}
