package benchmark

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
)

func BenchmarkIndexOf(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[n-1] // worst case: last element
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.IndexOf(ints, target)
			}
		})
	}
}

func BenchmarkLastIndexOf(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[0]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.LastIndexOf(ints, target)
			}
		})
	}
}

func BenchmarkHasPrefix(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		prefix := ints[:n/10+1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.HasPrefix(ints, prefix)
			}
		})
	}
}

func BenchmarkHasSuffix(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		suffix := ints[n-n/10-1:]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.HasSuffix(ints, suffix)
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.Find(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkFindIndexOf(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[n-1]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = lo.FindIndexOf(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkFindLastIndexOf(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		target := ints[0]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = lo.FindLastIndexOf(ints, func(v int) bool { return v == target })
			}
		})
	}
}

func BenchmarkFindOrElse(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindOrElse(ints, -1, func(v int) bool { return v == -999 })
			}
		})
	}
}

func BenchmarkFindKey(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FindKey(m, n/2)
			}
		})
	}
}

func BenchmarkFindKeyBy(b *testing.B) {
	for _, n := range lengths {
		m := genMap(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FindKeyBy(m, func(_ string, v int) bool { return v == n/2 })
			}
		})
	}
}

func BenchmarkFindUniques(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindUniques(ints)
			}
		})
	}
}

func BenchmarkFindUniquesBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindUniquesBy(ints, func(v int) int { return v % 50 })
			}
		})
	}
}

func BenchmarkFindDuplicates(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindDuplicates(ints)
			}
		})
	}
}

func BenchmarkFindDuplicatesBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.FindDuplicatesBy(ints, func(v int) int { return v % 50 })
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Min(ints)
			}
		})
	}
}

func BenchmarkMinIndex(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MinIndex(ints)
			}
		})
	}
}

func BenchmarkMinBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MinBy(ints, func(a, b int) bool { return a < b })
			}
		})
	}
}

func BenchmarkMinIndexBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MinIndexBy(ints, func(a, b int) bool { return a < b })
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Max(ints)
			}
		})
	}
}

func BenchmarkMaxIndex(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MaxIndex(ints)
			}
		})
	}
}

func BenchmarkMaxBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MaxBy(ints, func(a, b int) bool { return a > b })
			}
		})
	}
}

func BenchmarkMaxIndexBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.MaxIndexBy(ints, func(a, b int) bool { return a > b })
			}
		})
	}
}

func BenchmarkFirst(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = lo.First(ints)
	}
}

func BenchmarkFirstOrEmpty(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_ = lo.FirstOrEmpty(ints)
	}
}

func BenchmarkLast(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = lo.Last(ints)
	}
}

func BenchmarkLastOrEmpty(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_ = lo.LastOrEmpty(ints)
	}
}

func BenchmarkNth(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = lo.Nth(ints, 50)
	}
}

func BenchmarkSample(b *testing.B) {
	ints := genSliceInt(100)
	for i := 0; i < b.N; i++ {
		_ = lo.Sample(ints)
	}
}

func BenchmarkSamples(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Samples(ints, n/4)
			}
		})
	}
}
